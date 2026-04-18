package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type bookingRoutes struct {
	h       BookingHandler
	pattern string
}

func NewBookingRoutes(pattern string, h BookingHandler) RouteRegistrar {
	return bookingRoutes{pattern: pattern, h: h}
}

func (br bookingRoutes) Register(r chi.Router) {
	r.Route(br.pattern, func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/", br.h.CreateBooking)
		r.Get("/", br.h.GetBookings)
		r.Get("/{bookingId}", br.h.GetBooking)
		r.Patch("/{bookingId}/confirm", br.h.ConfirmBooking)
		r.Patch("/{bookingId}/cancel", br.h.CancelBooking)
		r.Delete("/{bookingId}", br.h.DeleteBooking)
	})
}
