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

	moderatedRooms := r.Group("/:countryCode/:citySlug/hotels/:hotelId/rooms")
	moderatedRooms.Use(
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("USER_ROLE_MODERATOR", "USER_ROLE_ADMIN"),
	)
	moderatedRooms.POST("", rr.h.CreateRoom)
	moderatedRooms.PUT("/:roomId", rr.h.UpdateRoom)
	moderatedRooms.PATCH("/:roomId/status", rr.h.UpdateRoomStatus)
	moderatedRooms.DELETE("/:roomId", rr.h.DeleteRoom)
}
