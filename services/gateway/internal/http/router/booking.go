package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type bookingHandler interface {
	CreateBooking(*gin.Context)
	GetBookings(*gin.Context)
	GetBooking(*gin.Context)
	ConfirmBooking(*gin.Context)
	CancelBooking(*gin.Context)
	DeleteBooking(*gin.Context)
}

type bookingRoutes struct {
	h       bookingHandler
	pattern string
}

func NewBookingRoutes(pattern string, h bookingHandler) RouteRegistrar {
	return bookingRoutes{pattern: pattern, h: h}
}

func (br bookingRoutes) Register(r *gin.RouterGroup) {
	bookings := r.Group(br.pattern)
	bookings.Use(middleware.AuthMiddleware())

	bookings.POST("", br.h.CreateBooking)
	bookings.GET("", br.h.GetBookings)
	bookings.GET("/:bookingId", br.h.GetBooking)
	bookings.PATCH("/:bookingId/confirm", br.h.ConfirmBooking)
	bookings.PATCH("/:bookingId/cancel", br.h.CancelBooking)
	bookings.DELETE("/:bookingId", br.h.DeleteBooking)
}
