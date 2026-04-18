package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type authRoutes struct {
	h       AuthHandler
	pattern string
}

func NewAuthRoutes(pattern string, h AuthHandler) RouteRegistrar {
	return authRoutes{pattern: pattern, h: h}
}

func (ar authRoutes) Register(r *gin.RouterGroup) {
	auth := r.Group(ar.pattern)
	auth.POST("/register", ar.h.Register)
	auth.POST("/login", ar.h.Login)
	auth.POST("/refresh", ar.h.RefreshToken)

	users := auth.Group("")
	users.Use(middleware.AuthMiddleware())
	users.GET("/users", ar.h.GetUsers)
	users.POST("/users", ar.h.CreateUser)
	users.GET("/users/:id", ar.h.GetUser)
	users.PUT("/users/:id", ar.h.UpdateUser)
	users.PATCH("/users/:id/activity", ar.h.UpdateUserActivity)
	users.PATCH("/users/:id/role", ar.h.UpdateUserRole)
	users.DELETE("/users/:id", ar.h.DeleteUser)
}
