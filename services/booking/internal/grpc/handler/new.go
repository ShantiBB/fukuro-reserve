package handler

import (
	"context"
	"time"

	"buf.build/go/protovalidate"
	"github.com/google/uuid"

	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/repository/models"
)

type BookingService interface {
	BookingCreate(
		ctx context.Context, b *models.CreateBooking, rooms []*models.CreateBookingRoom,
	) (*models.Booking, error)
	GetBookings(
		ctx context.Context, bookingRef models.BookingRef, page uint64, limit uint64,
	) (*models.BookingList, error)
	GetUnavailableRoomIDs(
		ctx context.Context, bookingRef models.BookingRef, checkIn time.Time, checkOut time.Time,
	) ([]uuid.UUID, error)
	QuoteBooking(
		ctx context.Context,
		checkIn time.Time,
		checkOut time.Time,
		currency string,
		rooms []*models.CreateBookingRoom,
	) (*models.BookingQuote, error)
	GetBookingById(ctx context.Context, bookingRef models.BookingRef, bookingID uuid.UUID) (*models.Booking, error)
	UpdateBookingStatus(ctx context.Context, bookingRef models.BookingRef, bookingID uuid.UUID, status models.BookingStatus) error
	DeleteBookingByID(ctx context.Context, bookingRef models.BookingRef, id uuid.UUID) error
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
