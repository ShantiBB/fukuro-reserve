package service

import (
	"context"
	"log/slog"
	"time"

	"booking/internal/repository/models"
	"booking/internal/service/utils/helper"
	"booking/pkg/lib/utils/consts"

	"github.com/google/uuid"
)

func (s *Service) BookingCreate(
	ctx context.Context,
	b *models.CreateBooking,
	rooms []*models.CreateBookingRoom,
) (*models.Booking, error) {
	if b == nil {
		return nil, consts.ErrNilObject
	}

	var err error
	b.FinalTotalAmount, err = helper.CalculateTotalAmount(b.CheckIn, b.CheckOut, rooms, b.ExpectedTotalAmount)
	if err != nil {
		slog.ErrorContext(ctx, "failed calculate total amount", "err", err)
		return nil, err
	}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to begin transaction", "err", err)
		return nil, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	newBooking, err := s.repo.CreateBooking(ctx, tx, b)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create booking", "err", err)
		return nil, err
	}

	for _, room := range rooms {
		room.BookingID = newBooking.ID
	}

	newRooms, err := s.repo.CreateBookingRooms(ctx, tx, newBooking.ID, rooms)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create booking rooms", "err", err)
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
		slog.ErrorContext(ctx, "failed to create room locks", "err", err)
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
		slog.ErrorContext(ctx, "failed to commit transaction", "err", err)
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
		slog.ErrorContext(ctx, "failed to get bookings", "err", err)
		return nil, err
	}

	bookingIDs := make([]uuid.UUID, len(bookingList.Bookings))
	for i, booking := range bookingList.Bookings {
		bookingIDs[i] = booking.ID
	}

	allRooms, err := s.repo.GetBookingRoomsByBookingIDs(ctx, nil, bookingIDs)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get booking rooms", "err", err)
		return nil, err
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
		slog.ErrorContext(ctx, "failed to get booking by id", "err", err)
		return nil, err
	}

	allRooms, err := s.repo.GetBookingRoomsWithLockByBookingIDs(ctx, nil, []uuid.UUID{booking.ID})
	if err != nil {
		slog.ErrorContext(ctx, "failed to get booking rooms by booking id", "err", err)
		return nil, err
	}

	booking.BookingRooms = allRooms
	return booking, nil
}

func (s *Service) UpdateBookingStatus(
	ctx context.Context,
	bookingID uuid.UUID,
	status models.BookingStatus,
) error {
	if err := s.repo.UpdateBookingStatusByID(ctx, nil, bookingID, status); err != nil {
		slog.ErrorContext(ctx, "failed to update booking status", "err", err)
		return err
	}

	var roomLockStatus = &models.RoomLockActivity{}
	t := time.Now()
	if status == models.BookingStatusCancelled {
		roomLockStatus.IsActive = true
		roomLockStatus.ExpiresAt = &t
	}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to begin transaction", "err", err)
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if err = s.repo.UpdateRoomLocksActivityByID(ctx, nil, bookingID, roomLockStatus); err != nil {
		slog.ErrorContext(ctx, "failed to update room locks activity", "err", err)
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to commit transaction", "err", err)
		return err
	}

	return nil
}

func (s *Service) DeleteBookingByID(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.DeleteBookingByID(ctx, nil, id); err != nil {
		slog.ErrorContext(ctx, "failed to delete booking by id", "err", err)
		return err
	}

	return nil
}
