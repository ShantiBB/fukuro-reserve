package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpswagger "github.com/swaggo/http-swagger"

	_ "github.com/ShantiBB/fukuro-reserve/services/gateway/docs"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/handler"
)

func New(r chi.Router, authHandler *handler.AuthHandler, hotelHandler *handler.HotelHandler, bookingHandler *handler.BookingHandler) {
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Timeout(60 * time.Second))

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

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/docs/swagger/*", httpswagger.WrapHandler)

		authRouter("/auth", r, authHandler)
		hotelRouter("/hotels", r, hotelHandler)
		roomRouter("/rooms", r, hotelHandler)
		bookingRouter("/bookings", r, bookingHandler)
	})
}
