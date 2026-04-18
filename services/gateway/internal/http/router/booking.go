package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/handler"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

func bookingRouter(pattern string, r chi.Router, h *handler.BookingHandler) {
	r.Route(pattern, func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/", h.CreateBooking)
		r.Get("/", h.GetBookings)
		r.Get("/{bookingId}", h.GetBooking)
		r.Patch("/{bookingId}/confirm", h.ConfirmBooking)
		r.Patch("/{bookingId}/cancel", h.CancelBooking)
		r.Delete("/{bookingId}", h.DeleteBooking)
	})
}
