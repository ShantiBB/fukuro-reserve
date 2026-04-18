package handler

import hotelservice "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/service/hotel"

type HotelHandler struct {
	service *hotelservice.Service
}

func NewHotelHandler(service *hotelservice.Service) *HotelHandler {
	return &HotelHandler{service: service}
}
