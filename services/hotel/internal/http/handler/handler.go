package handler

import (
	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
	"hotel/internal/repository/postgres/models"
)

type Service interface {
	HotelService
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) HotelCreateRequestToEntity(req request.HotelCreate) models.HotelCreate {
	location := models.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return models.HotelCreate{
		Name:        req.Name,
		OwnerID:     req.OwnerID,
		Description: req.Description,
		Address:     req.Address,
		Location:    location,
	}
}

func (h *Handler) HotelRequestToEntity(req models.Hotel) response.Hotel {
	location := response.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return response.Hotel{
		Name:        req.Name,
		OwnerID:     req.OwnerID,
		Description: req.Description,
		Address:     req.Address,
		Location:    location,
	}
}

func (h *Handler) HotelResponseToEntity(req models.Hotel) response.Hotel {
	location := response.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return response.Hotel{
		ID:          req.ID,
		Name:        req.Name,
		OwnerID:     req.OwnerID,
		Description: req.Description,
		Address:     req.Address,
		Rating:      req.Rating,
		Location:    location,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}
}
