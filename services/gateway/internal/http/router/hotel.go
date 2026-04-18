package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type hotelRoutes struct {
	h       HotelHandler
	pattern string
}

func NewHotelRoutes(pattern string, h HotelHandler) RouteRegistrar {
	return hotelRoutes{pattern: pattern, h: h}
}

func (hr hotelRoutes) Register(r chi.Router) {
	r.Route(hr.pattern, func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/", hr.h.CreateHotel)
		r.Get("/", hr.h.GetHotels)
		r.Get("/{countryCode}/{citySlug}/{hotelSlug}", hr.h.GetHotel)
		r.Put("/{countryCode}/{citySlug}/{hotelSlug}", hr.h.UpdateHotel)
		r.Patch("/{countryCode}/{citySlug}/{hotelSlug}/title", hr.h.UpdateHotelTitle)
		r.Delete("/{countryCode}/{citySlug}/{hotelSlug}", hr.h.DeleteHotel)
	})
}
