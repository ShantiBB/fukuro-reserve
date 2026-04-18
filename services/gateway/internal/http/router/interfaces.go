package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RouteRegistrar interface {
	Register(r chi.Router)
}

type AuthHandler interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	RefreshToken(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	CreateUser(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	UpdateUserActivity(http.ResponseWriter, *http.Request)
	UpdateUserRole(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

type HotelHandler interface {
	CreateHotel(http.ResponseWriter, *http.Request)
	GetHotels(http.ResponseWriter, *http.Request)
	GetHotel(http.ResponseWriter, *http.Request)
	UpdateHotel(http.ResponseWriter, *http.Request)
	UpdateHotelTitle(http.ResponseWriter, *http.Request)
	DeleteHotel(http.ResponseWriter, *http.Request)
	CreateRoom(http.ResponseWriter, *http.Request)
	GetRooms(http.ResponseWriter, *http.Request)
	GetRoom(http.ResponseWriter, *http.Request)
	UpdateRoom(http.ResponseWriter, *http.Request)
	UpdateRoomStatus(http.ResponseWriter, *http.Request)
	DeleteRoom(http.ResponseWriter, *http.Request)
}

type BookingHandler interface {
	CreateBooking(http.ResponseWriter, *http.Request)
	GetBookings(http.ResponseWriter, *http.Request)
	GetBooking(http.ResponseWriter, *http.Request)
	ConfirmBooking(http.ResponseWriter, *http.Request)
	CancelBooking(http.ResponseWriter, *http.Request)
	DeleteBooking(http.ResponseWriter, *http.Request)
}
