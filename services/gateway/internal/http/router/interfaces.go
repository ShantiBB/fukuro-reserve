package router

import "github.com/gin-gonic/gin"

type RouteRegistrar interface {
	Register(r *gin.RouterGroup)
}

type AuthHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
	RefreshToken(*gin.Context)
	GetUsers(*gin.Context)
	CreateUser(*gin.Context)
	GetUser(*gin.Context)
	UpdateUser(*gin.Context)
	UpdateUserActivity(*gin.Context)
	UpdateUserRole(*gin.Context)
	DeleteUser(*gin.Context)
}

type HotelHandler interface {
	CreateHotel(*gin.Context)
	GetHotels(*gin.Context)
	GetHotel(*gin.Context)
	UpdateHotel(*gin.Context)
	UpdateHotelTitle(*gin.Context)
	DeleteHotel(*gin.Context)
	CreateRoom(*gin.Context)
	GetRooms(*gin.Context)
	GetRoom(*gin.Context)
	UpdateRoom(*gin.Context)
	UpdateRoomStatus(*gin.Context)
	DeleteRoom(*gin.Context)
}

type BookingHandler interface {
	CreateBooking(*gin.Context)
	GetBookings(*gin.Context)
	GetBooking(*gin.Context)
	ConfirmBooking(*gin.Context)
	CancelBooking(*gin.Context)
	DeleteBooking(*gin.Context)
}
