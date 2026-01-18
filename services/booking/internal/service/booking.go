package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"booking/internal/repository/models"
	"booking/internal/service/utils/helper"
	"booking/pkg/utils/consts"

	"github.com/google/uuid"
)

func (s *Service) BookingCreate(
	ctx context.Context,
	b *models.CreateBooking,
	rooms []*models.CreateBookingRoom,
) (*models.Booking, error) {
	if b == nil {
		return nil, errors.New("booking cannot be nil")
	}

	var err error
	b.FinalTotalAmount, err = helper.CalculateFinalTotalAmount(b.CheckIn, b.CheckOut, rooms, b.ExpectedTotalAmount)
	if err != nil {
		return nil, err
	}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	newBooking, err := s.repo.CreateBooking(ctx, tx, b)
	if err != nil {
		return nil, err
	}

	for _, room := range rooms {
		room.BookingID = newBooking.ID
	}

	newRooms, err := s.repo.CreateBookingRooms(ctx, tx, newBooking.ID, rooms)
	if err != nil {
		return nil, err
	}

	locks := make([]*models.CreateRoomLock, len(newRooms))
	for i, nr := range newRooms {
		locks[i] = &models.CreateRoomLock{
			RoomID:    nr.RoomID,
			BookingID: newBooking.ID,
			StayRange: models.DateRange{
				Start: b.CheckIn,
				End:   b.CheckOut,
			},
			ExpiresAt: time.Now().Add(consts.ExpireRoomLockMinutes * time.Minute),
		}
	}

	newLocks, err := s.repo.CreateRoomLocks(ctx, tx, locks)
	if err != nil {
		return nil, err
	}

	locksByRoomID := make(map[uuid.UUID]models.RoomLockShort)
	for _, lock := range newLocks {
		locksByRoomID[lock.RoomID] = models.RoomLockShort{
			ID:        lock.ID,
			ISActive:  lock.ISActive,
			ExpiresAt: lock.ExpiresAt,
			CreatedAt: lock.CreatedAt,
		}
	}

	for i := range newRooms {
		if lock, exists := locksByRoomID[newRooms[i].RoomID]; exists {
			newRooms[i].RoomLock = lock
		}
	}

	newBooking.BookingRooms = newRooms

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return newBooking, nil
}

func (s *Service) GetBookings(
	ctx context.Context,
	bookingRef models.BookingRef,
	page uint64,
	limit uint64,
) (*models.BookingList, error) {
	offset := (page - 1) * limit
	bookingList, err := s.repo.GetBookingsByHotelInfo(ctx, nil, bookingRef, limit, offset)
	if err != nil {
		return nil, err
	}

	bookingIDs := make([]uuid.UUID, len(bookingList.Bookings))
	for i, booking := range bookingList.Bookings {
		bookingIDs[i] = booking.ID
	}

	allRooms, err := s.repo.GetBookingRoomsByBookingIDs(ctx, nil, bookingIDs)
	if err != nil {
		return nil, fmt.Errorf("get booking rooms: %w", err)
	}

	roomsByBookingID := make(map[uuid.UUID][]*models.BookingRoom)
	for _, room := range allRooms {
		roomsByBookingID[room.BookingID] = append(roomsByBookingID[room.BookingID], room)
	}

	for i := range bookingList.Bookings {
		bookingList.Bookings[i].BookingRooms = roomsByBookingID[bookingList.Bookings[i].ID]
	}

	return bookingList, nil
}

func (s *Service) GetBookingById(ctx context.Context, bookingID uuid.UUID) (*models.Booking, error) {
	booking, err := s.repo.GetBookingByID(ctx, nil, bookingID)
	if err != nil {
		return nil, err
	}

	allRooms, err := s.repo.GetBookingRoomsWithLockByBookingIDs(ctx, nil, []uuid.UUID{booking.ID})
	if err != nil {
		return nil, fmt.Errorf("get booking rooms: %w", err)
	}

	booking.BookingRooms = allRooms
	return booking, nil
}

//func (s *Service) BookingGetByID(ctx context.Context, id uuid.UUID) (models.Booking, error) {
//	b, err := s.repo.BookingGetByID(ctx, id)
//	if err != nil {
//		return models.Booking{}, err
//	}
//
//	return b, nil
//}

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
