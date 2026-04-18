package handler

import "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/service"

type AuthHandler struct {
	service *service.Auth
}

type HotelHandler struct {
	service *service.Hotel
}

type BookingHandler struct {
	service *service.Booking
}

func NewAuthHandler(svc *service.Auth) *AuthHandler {
	return &AuthHandler{service: svc}
}

func NewHotelHandler(svc *service.Hotel) *HotelHandler {
	return &HotelHandler{service: svc}
}

func NewBookingHandler(svc *service.Booking) *BookingHandler {
	return &BookingHandler{service: svc}
}
