package service

import (
	"booking/internal/repository/models"
	"context"

	"github.com/google/uuid"
)

type BookingRepository interface {
	BookingCreate(ctx context.Context, b models.CreateBooking) (models.Booking, error)
	BookingGetAll(
		ctx context.Context,
		bookingRef models.BookingRef,
		limit uint64,
		offset uint64,
	) (models.BookingList, error)
	BookingGetByID(ctx context.Context, id uuid.UUID) (models.Booking, error)
	BookingUpdateByID(ctx context.Context, id uuid.UUID, b models.UpdateBooking) error
	BookingStatusUpdateByID(ctx context.Context, id uuid.UUID, b models.BookingStatusInfo) error
	BookingDeleteByID(ctx context.Context, id uuid.UUID) error
}

func (s *Service) BookingCreate(ctx context.Context, b models.CreateBooking) (models.Booking, error) {

	newBooking, err := s.repo.BookingCreate(ctx, b)
	if err != nil {
		return models.Booking{}, err
	}

	return newBooking, nil
}

func (s *Service) BookingGetAll(
	ctx context.Context,
	bookingRef models.BookingRef,
	page uint64,
	limit uint64,
) (models.BookingList, error) {
	offset := (page - 1) * limit

	bookingList, err := s.repo.BookingGetAll(ctx, bookingRef, limit, offset)
	if err != nil {
		return models.BookingList{}, err
	}

	return bookingList, nil
}

func (s *Service) BookingGetByID(ctx context.Context, id uuid.UUID) (models.Booking, error) {
	b, err := s.repo.BookingGetByID(ctx, id)
	if err != nil {
		return models.Booking{}, err
	}

	return b, nil
}

func (s *Service) BookingUpdateByID(ctx context.Context, id uuid.UUID, b models.UpdateBooking) error {
	if err := s.repo.BookingUpdateByID(ctx, id, b); err != nil {
		return err
	}

	return nil
}

func (s *Service) BookingStatusUpdateByID(ctx context.Context, id uuid.UUID, b models.BookingStatusInfo) error {
	if err := s.repo.BookingStatusUpdateByID(ctx, id, b); err != nil {
		return err
	}

	return nil
}

func (s *Service) BookingDeleteByID(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.BookingDeleteByID(ctx, id); err != nil {
		return err
	}

	return nil
}
