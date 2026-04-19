package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type hotelHandler interface {
	CreateHotel(*gin.Context)
	GetHotels(*gin.Context)
	GetHotelBySlug(*gin.Context)
	GetHotelByID(*gin.Context)
	UpdateHotel(*gin.Context)
	UpdateHotelTitle(*gin.Context)
	DeleteHotel(*gin.Context)
}

type hotelRoutes struct {
	h       hotelHandler
	pattern string
}

func NewHotelRoutes(pattern string, h hotelHandler) RouteRegistrar {
	return hotelRoutes{pattern: pattern, h: h}
}

func (hr hotelRoutes) Register(r *gin.RouterGroup) {
	publicHotels := r.Group("/:countryCode/:citySlug/hotels")
	publicHotels.GET("", hr.h.GetHotels)
	publicHotels.GET("/slug/:hotelSlug", hr.h.GetHotelBySlug)
	publicHotels.GET("/:hotelId", hr.h.GetHotelByID)

	protectedHotels := r.Group("/:countryCode/:citySlug/hotels")
	protectedHotels.Use(middleware.AuthMiddleware())
	protectedHotels.POST("", hr.h.CreateHotel)
	protectedHotels.PUT("/:hotelId", hr.h.UpdateHotel)
	protectedHotels.PATCH("/:hotelId/title", hr.h.UpdateHotelTitle)
	protectedHotels.DELETE("/:hotelId", hr.h.DeleteHotel)
}
