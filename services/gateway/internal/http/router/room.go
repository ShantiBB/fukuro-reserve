package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type roomHandler interface {
	CreateRoom(*gin.Context)
	GetRoomsByHotelSlug(*gin.Context)
	GetRoomByHotelSlug(*gin.Context)
	GetRoomsByHotelID(*gin.Context)
	GetRoomByHotelID(*gin.Context)
	UpdateRoom(*gin.Context)
	UpdateRoomStatus(*gin.Context)
	DeleteRoom(*gin.Context)
}

type roomRoutes struct {
	h       roomHandler
	pattern string
}

func NewRoomRoutes(pattern string, h roomHandler) RouteRegistrar {
	return roomRoutes{pattern: pattern, h: h}
}

func (rr roomRoutes) Register(r *gin.RouterGroup) {
	slugRooms := r.Group("/:countryCode/:citySlug/hotels/slug/:hotelSlug/rooms")
	slugRooms.Use(middleware.AuthMiddleware())

	idRooms := r.Group("/:countryCode/:citySlug/hotels/:hotelId/rooms")
	idRooms.Use(middleware.AuthMiddleware())

	slugRooms.GET("", rr.h.GetRoomsByHotelSlug)
	slugRooms.GET("/:roomId", rr.h.GetRoomByHotelSlug)

	idRooms.POST("", rr.h.CreateRoom)
	idRooms.GET("", rr.h.GetRoomsByHotelID)
	idRooms.GET("/:roomId", rr.h.GetRoomByHotelID)
	idRooms.PUT("/:roomId", rr.h.UpdateRoom)
	idRooms.PATCH("/:roomId/status", rr.h.UpdateRoomStatus)
	idRooms.DELETE("/:roomId", rr.h.DeleteRoom)
}
