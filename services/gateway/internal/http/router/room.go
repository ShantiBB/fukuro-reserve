package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/handler"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

func roomRouter(pattern string, r chi.Router, h *handler.HotelHandler) {
	r.Route(pattern, func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/", h.CreateRoom)
		r.Get("/", h.GetRooms)
		r.Get("/{roomId}", h.GetRoom)
		r.Put("/{roomId}", h.UpdateRoom)
		r.Patch("/{roomId}/status", h.UpdateRoomStatus)
		r.Delete("/{roomId}", h.DeleteRoom)
	})
}
