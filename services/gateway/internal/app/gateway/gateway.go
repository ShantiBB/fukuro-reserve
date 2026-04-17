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
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/handler"
	httpMiddleware "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
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

	// Middleware
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Timeout(60 * time.Second))

	// CORS
	r.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: true,
				MaxAge:           300,
			},
		),
	)

	// Create handlers
	authHandler := handler.NewAuthHandler(app.Clients)
	hotelHandler := handler.NewHotelHandler(app.Clients)
	bookingHandler := handler.NewBookingHandler(app.Clients)

	// Routes
	r.Route(
		"/api/v1", func(r chi.Router) {
			// Auth routes
			r.Mount("/auth", authHandler.Routes())

			// Hotel routes
			r.Mount("/hotels", hotelHandler.HotelRoutes())
			r.Mount("/rooms", hotelHandler.RoomRoutes())

			// Booking routes
			r.Mount("/bookings", bookingHandler.Routes())
		},
	)

	// Health check
	r.Get(
		"/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
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
