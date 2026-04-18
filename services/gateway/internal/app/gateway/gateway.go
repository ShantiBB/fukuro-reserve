package gateway

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/handler"
	httpMiddleware "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/router"
	"github.com/go-chi/chi/v5"
)

type App struct {
	Config  *config.Config
	Logger  *slog.Logger
	Clients *clients.Clients
}

func New(cfg *config.Config, logger *slog.Logger) (*App, error) {
	grpcClients, err := clients.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC clients: %w", err)
	}

	return &App{
		Config:  cfg,
		Logger:  logger,
		Clients: grpcClients,
	}, nil
}

func (app *App) MustRun() {
	slog.SetDefault(app.Logger)

	// Set JWT secret for auth middleware
	httpMiddleware.SetJWTSecret(app.Config.JWT.AccessSecret)

	// Create router
	r := chi.NewRouter()

	// Create handlers
	authHandler := handler.NewAuthHandler(app.Clients)
	hotelHandler := handler.NewHotelHandler(app.Clients)
	bookingHandler := handler.NewBookingHandler(app.Clients)

	router.New(r, authHandler, hotelHandler, bookingHandler)

	// Health check
	r.Get(
		"/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte("OK")); err != nil {
				slog.Error("Failed to write health response", "error", err)
			}
		},
	)

	// Start server
	addr := fmt.Sprintf("%s:%d", app.Config.HTTP.Host, app.Config.HTTP.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		slog.Info("Starting HTTP server", "address", addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Failed to start HTTP server", "error", err)
		}
	}()

	app.gracefulShutdown(server)
}

func (app *App) gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	slog.Info("Shutting down HTTP server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error shutting down HTTP server", "error", err)
	}

	if err := app.Clients.Close(); err != nil {
		slog.Error("Error closing gRPC connections", "error", err)
	}

	slog.Info("HTTP server stopped")
}

func (app *App) Close() error {
	return app.Clients.Close()
}
