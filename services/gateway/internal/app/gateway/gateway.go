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

	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/handler"
	httpMiddleware "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/router"
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

	// Create handlers
	authHandler := handler.NewAuthHandler(app.Clients, app.Config.Pagination)
	hotelHandler := handler.NewHotelHandler(app.Clients, app.Config.Pagination)
	bookingHandler := handler.NewBookingHandler(app.Clients, app.Config.Pagination)

	// Create router
	r := chi.NewRouter()
	router.New(
		r,
		app.Config.HTTP,
		app.Config.CORS,
		router.NewAuthRoutes("/auth", authHandler),
		router.NewHotelRoutes("/hotels", hotelHandler),
		router.NewRoomRoutes("/rooms", hotelHandler),
		router.NewBookingRoutes("/bookings", bookingHandler),
	)

	// Start server
	addr := fmt.Sprintf("%s:%d", app.Config.HTTP.Host, app.Config.HTTP.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Duration(app.Config.HTTP.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(app.Config.HTTP.WriteTimeoutSec) * time.Second,
		IdleTimeout:  time.Duration(app.Config.HTTP.IdleTimeoutSec) * time.Second,
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

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(app.Config.HTTP.ShutdownTimeoutSec)*time.Second,
	)
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
