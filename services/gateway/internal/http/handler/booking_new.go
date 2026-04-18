package handler

import bookingservice "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/service/booking"

type BookingHandler struct {
	service *bookingservice.Service
}

func NewBookingHandler(service *bookingservice.Service) *BookingHandler {
	return &BookingHandler{service: service}
}
