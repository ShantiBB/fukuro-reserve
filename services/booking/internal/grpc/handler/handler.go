package handler

import (
	"context"
	"sync"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"

	"buf.build/go/protovalidate"
)

type Handler struct {
	bookingv1.UnimplementedBookingServiceServer
	svc       BookingService
	validator protovalidate.Validator
}

func New(svc BookingService, validator protovalidate.Validator) *Handler {
	return &Handler{svc: svc, validator: validator}
}

type BookingService interface {
	BookingCreate(ctx context.Context, b models.CreateBooking, rooms []models.CreateBookingRoom) (models.Booking, error)
}

var (
	validatorOnce sync.Once
	validator     protovalidate.Validator
	validatorErr  error
)

func getValidator() (protovalidate.Validator, error) {
	validatorOnce.Do(
		func() {
			validator, validatorErr = protovalidate.New()
		},
	)
	return validator, validatorErr
}
