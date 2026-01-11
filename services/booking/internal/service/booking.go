package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"booking/internal/repository/models"

	"github.com/google/uuid"
)

type BookingTransactionRepository interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
}

type BookingRepository interface {
	CreateBooking(ctx context.Context, tx pgx.Tx, b models.CreateBooking) (models.Booking, error)
	GetBookingsByHotelInfo(
		ctx context.Context,
		tx pgx.Tx,
		bookingRef models.BookingRef,
		limit uint64,
		offset uint64,
	) (models.BookingList, error)
	GetBookingByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) (models.Booking, error)
	UpdateBookingGuestInfoByID(ctx context.Context, tx pgx.Tx, id uuid.UUID, b models.UpdateBooking) error
	UpdateBookingStatusByID(ctx context.Context, tx pgx.Tx, id uuid.UUID, b models.BookingStatusInfo) error
	DeleteBookingByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
}

type BookingRoomRepository interface {
	CreateBookingRoom(ctx context.Context, tx pgx.Tx, bRoom models.CreateBookingRoom) (models.BookingRoom, error)
	GetBookingRoomsByBookingID(ctx context.Context, tx pgx.Tx, bookingID uuid.UUID) ([]models.BookingRoom, error)
	GetBookingRoomByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) (models.BookingRoom, error)
	UpdateBookingRoomGuestCountsByID(
		ctx context.Context,
		tx pgx.Tx,
		id uuid.UUID,
		bRoom models.BookingRoomGuestCounts,
	) error
	DeleteBookingRoomByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
}

type RoomLockRepository interface {
	CreateRoomLock(ctx context.Context, tx pgx.Tx, roomLock models.CreateRoomLock) (models.RoomLock, error)
	GetRoomsLockByBookingID(ctx context.Context, tx pgx.Tx, bookingID uuid.UUID) ([]models.RoomLock, error)
	GetRoomLockByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) (models.RoomLock, error)
	UpdateRoomLockActivityByID(
		ctx context.Context,
		tx pgx.Tx,
		id uuid.UUID,
		roomLock models.UpdateRoomLockActivity,
	) error
	DeleteRoomLockByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
}

func (s *Service) BookingCreate(
	ctx context.Context,
	b models.CreateBooking,
	rooms []models.CreateBookingRoom,
) (models.Booking, error) {

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return models.Booking{}, err
	}
	defer tx.Rollback(ctx)

	newBooking, err := s.repo.CreateBooking(ctx, tx, b)
	if err != nil {
		return models.Booking{}, err
	}

	var newRoom models.BookingRoom
	for _, room := range rooms {
		room.BookingID = newBooking.ID

		newRoom, err = s.repo.CreateBookingRoom(ctx, tx, room)
		if err != nil {
			return models.Booking{}, err
		}

		roomLock := models.CreateRoomLock{
			RoomID:    newRoom.RoomID,
			BookingID: newBooking.ID,
			StayRange: models.DateRange{
				Start: b.CheckIn,
				End:   b.CheckOut,
			},
			ExpiresAt: time.Now().Add(15 * time.Minute),
		}

		if _, err = s.repo.CreateRoomLock(ctx, tx, roomLock); err != nil {
			return models.Booking{}, err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return models.Booking{}, err
	}

	return newBooking, nil
}

//func (s *Service) BookingGetAll(
//	ctx context.Context,
//	bookingRef models.BookingRef,
//	page uint64,
//	limit uint64,
//) (models.BookingList, error) {
//	offset := (page - 1) * limit
//
//	bookingList, err := s.repo.BookingGetAll(ctx, bookingRef, limit, offset)
//	if err != nil {
//		return models.BookingList{}, err
//	}
//
//	return bookingList, nil
//}
//
//func (s *Service) BookingGetByID(ctx context.Context, id uuid.UUID) (models.Booking, error) {
//	b, err := s.repo.BookingGetByID(ctx, id)
//	if err != nil {
//		return models.Booking{}, err
//	}
//
//	return b, nil
//}
//
//func (s *Service) BookingUpdateByID(ctx context.Context, id uuid.UUID, b models.UpdateBooking) error {
//	if err := s.repo.BookingUpdateByID(ctx, id, b); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *Service) BookingStatusUpdateByID(ctx context.Context, id uuid.UUID, b models.BookingStatusInfo) error {
//	if err := s.repo.BookingStatusUpdateByID(ctx, id, b); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *Service) BookingDeleteByID(ctx context.Context, id uuid.UUID) error {
//	if err := s.repo.BookingDeleteByID(ctx, id); err != nil {
//		return err
//	}
//
//	return nil
//}
