package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type roomHandler interface {
	CreateRoom(*gin.Context)
	GetRoom(*gin.Context)
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
	rooms := r.Group(rr.pattern)
	rooms.Use(middleware.AuthMiddleware())

	rooms.POST("", rr.h.CreateRoom)
	rooms.GET("/:roomId", rr.h.GetRoom)
	rooms.PUT("/:roomId", rr.h.UpdateRoom)
	rooms.PATCH("/:roomId/status", rr.h.UpdateRoomStatus)
	rooms.DELETE("/:roomId", rr.h.DeleteRoom)
}
