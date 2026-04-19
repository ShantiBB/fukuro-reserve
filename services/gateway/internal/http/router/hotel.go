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
	locationHotels := r.Group("/:countryCode/:citySlug/hotels")
	locationHotels.Use(middleware.AuthMiddleware())

	locationHotels.POST("", hr.h.CreateHotel)
	locationHotels.GET("", hr.h.GetHotels)
	locationHotels.GET("/slug/:hotelSlug", hr.h.GetHotelBySlug)
	locationHotels.GET("/:hotelId", hr.h.GetHotelByID)
	locationHotels.PUT("/:hotelId", hr.h.UpdateHotel)
	locationHotels.PATCH("/:hotelId/title", hr.h.UpdateHotelTitle)
	locationHotels.DELETE("/:hotelId", hr.h.DeleteHotel)
}
