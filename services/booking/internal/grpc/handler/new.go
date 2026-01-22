package handler

import (
	"context"

	"buf.build/go/protovalidate"
	"github.com/google/uuid"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
)

type BookingService interface {
	BookingCreate(
		ctx context.Context, b *models.CreateBooking, rooms []*models.CreateBookingRoom,
	) (*models.Booking, error)
	GetBookings(
		ctx context.Context, bookingRef models.BookingRef, page uint64, limit uint64,
	) (*models.BookingList, error)
	GetBookingById(ctx context.Context, bookingID uuid.UUID) (*models.Booking, error)
	UpdateBookingStatus(ctx context.Context, bookingID uuid.UUID, status models.BookingStatus) error
	DeleteBookingByID(ctx context.Context, id uuid.UUID) error
}

type Service interface {
	BookingService
}

type Handler struct {
	bookingv1.UnimplementedBookingServiceServer
	svc       Service
	validator protovalidate.Validator
}

func New(svc Service, validator protovalidate.Validator) *Handler {
	return &Handler{svc: svc, validator: validator}
}
