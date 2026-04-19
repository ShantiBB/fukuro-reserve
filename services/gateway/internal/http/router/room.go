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
	slugRoomsPublic := r.Group("/:countryCode/:citySlug/hotels/slug/:hotelSlug/rooms")
	slugRoomsPublic.GET("", rr.h.GetRoomsByHotelSlug)
	slugRoomsPublic.GET("/:roomId", rr.h.GetRoomByHotelSlug)

	idRoomsPublic := r.Group("/:countryCode/:citySlug/hotels/:hotelId/rooms")
	idRoomsPublic.GET("", rr.h.GetRoomsByHotelID)
	idRoomsPublic.GET("/:roomId", rr.h.GetRoomByHotelID)

	idRoomsProtected := r.Group("/:countryCode/:citySlug/hotels/:hotelId/rooms")
	idRoomsProtected.Use(middleware.AuthMiddleware())
	idRoomsProtected.POST("", rr.h.CreateRoom)
	idRoomsProtected.PUT("/:roomId", rr.h.UpdateRoom)
	idRoomsProtected.PATCH("/:roomId/status", rr.h.UpdateRoomStatus)
	idRoomsProtected.DELETE("/:roomId", rr.h.DeleteRoom)
}
