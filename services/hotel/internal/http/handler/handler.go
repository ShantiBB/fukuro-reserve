package handler

import (
	"fukuro-reserve/pkg/utils/consts"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	HotelService
	RoomService
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) customValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return consts.FieldRequired.Error()
	default:
		return consts.InternalServer.Error()
	}
}
