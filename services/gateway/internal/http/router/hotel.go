package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type hotelRoutes struct {
	h       HotelHandler
	pattern string
}

func NewHotelRoutes(pattern string, h HotelHandler) RouteRegistrar {
	return hotelRoutes{pattern: pattern, h: h}
}

func (hr hotelRoutes) Register(r *gin.RouterGroup) {
	hotels := r.Group(hr.pattern)
	hotels.Use(middleware.AuthMiddleware())

	hotels.POST("", hr.h.CreateHotel)
	hotels.GET("", hr.h.GetHotels)
	hotels.GET("/:countryCode/:citySlug/:hotelSlug", hr.h.GetHotel)
	hotels.PUT("/:countryCode/:citySlug/:hotelSlug", hr.h.UpdateHotel)
	hotels.PATCH("/:countryCode/:citySlug/:hotelSlug/title", hr.h.UpdateHotelTitle)
	hotels.DELETE("/:countryCode/:citySlug/:hotelSlug", hr.h.DeleteHotel)
}
