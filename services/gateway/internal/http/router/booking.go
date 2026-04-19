package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type bookingHandler interface {
	CreateBooking(*gin.Context)
	QuoteBooking(*gin.Context)
	GetBookings(*gin.Context)
	GetMyBookings(*gin.Context)
	GetRoomBookings(*gin.Context)
	GetAvailability(*gin.Context)
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
	r.GET("/:countryCode/:citySlug/hotels/:hotelId/rooms/availability", br.h.GetAvailability)
	r.POST("/:countryCode/:citySlug/hotels/:hotelId/bookings/quote", br.h.QuoteBooking)

	myBookings := r.Group("/:countryCode/:citySlug/users/me/bookings")
	myBookings.Use(middleware.AuthMiddleware())
	myBookings.GET("", br.h.GetMyBookings)

	bookings := r.Group("/:countryCode/:citySlug/hotels/:hotelId/bookings")
	bookings.Use(middleware.AuthMiddleware())

	bookings.POST("", br.h.CreateBooking)
	bookings.GET("", br.h.GetBookings)
	bookings.GET("/:bookingId", br.h.GetBooking)
	bookings.PATCH("/:bookingId/confirm", br.h.ConfirmBooking)
	bookings.PATCH("/:bookingId/cancel", br.h.CancelBooking)
	bookings.DELETE("/:bookingId", br.h.DeleteBooking)

	roomBookings := r.Group("/:countryCode/:citySlug/hotels/:hotelId/rooms/:roomId/bookings")
	roomBookings.Use(middleware.AuthMiddleware())
	roomBookings.GET("", br.h.GetRoomBookings)
}
