package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type roomRoutes struct {
	h       HotelHandler
	pattern string
}

func NewRoomRoutes(pattern string, h HotelHandler) RouteRegistrar {
	return roomRoutes{pattern: pattern, h: h}
}

func (rr roomRoutes) Register(r chi.Router) {
	r.Route(rr.pattern, func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/", rr.h.CreateRoom)
		r.Get("/", rr.h.GetRooms)
		r.Get("/{roomId}", rr.h.GetRoom)
		r.Put("/{roomId}", rr.h.UpdateRoom)
		r.Patch("/{roomId}/status", rr.h.UpdateRoomStatus)
		r.Delete("/{roomId}", rr.h.DeleteRoom)
	})
}
