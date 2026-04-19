package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/repository/models"
	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/service/utils/helper"
	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/utils/consts"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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

func (s *Service) GetUnavailableRoomIDs(
	ctx context.Context,
	bookingRef models.BookingRef,
	checkIn time.Time,
	checkOut time.Time,
) ([]uuid.UUID, error) {
	roomIDs, err := s.repo.GetUnavailableRoomIDs(ctx, nil, bookingRef, checkIn, checkOut)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get unavailable room ids", "err", err)
		return nil, err
	}

	return roomIDs, nil
}

func (s *Service) QuoteBooking(
	ctx context.Context,
	checkIn time.Time,
	checkOut time.Time,
	currency string,
	rooms []*models.CreateBookingRoom,
) (*models.BookingQuote, error) {
	total, err := helper.CalculateTotalAmount(checkIn, checkOut, rooms, decimal.Zero)
	if err != nil {
		slog.ErrorContext(ctx, "failed to calculate booking quote", "err", err)
		return nil, err
	}

	nights, err := helper.Nights(checkIn, checkOut)
	if err != nil {
		slog.ErrorContext(ctx, "failed to calculate booking nights", "err", err)
		return nil, err
	}
	nightsDec := decimal.NewFromInt(int64(nights))

	quoteRooms := make([]*models.BookingQuoteRoom, len(rooms))
	for i, room := range rooms {
		quoteRooms[i] = &models.BookingQuoteRoom{
			RoomID:        room.RoomID,
			Adults:        room.Adults,
			Children:      room.Children,
			PricePerNight: room.PricePerNight,
			TotalAmount:   room.PricePerNight.Mul(nightsDec),
		}
	}

	return &models.BookingQuote{
		CheckIn:     checkIn,
		CheckOut:    checkOut,
		Currency:    currency,
		TotalAmount: total,
		Rooms:       quoteRooms,
		Nights:      uint32(nights),
	}, nil
}

func (s *Service) GetBookingById(ctx context.Context, bookingRef models.BookingRef, bookingID uuid.UUID) (*models.Booking, error) {
	booking, err := s.repo.GetBookingByID(ctx, nil, bookingRef, bookingID)
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
	bookingRef models.BookingRef,
	bookingID uuid.UUID,
	status models.BookingStatus,
) error {
	checkOut, err := s.repo.UpdateBookingStatusByID(ctx, nil, bookingRef, bookingID, status)
	if err != nil {
		slog.ErrorContext(ctx, "failed to update booking status", "err", err)
		return err
	}

	var roomLockStatus = &models.RoomLockActivity{}

	switch status {
	case models.BookingStatusConfirmed:
		roomLockStatus.IsActive = false
		roomLockStatus.ExpiresAt = &checkOut
	case models.BookingStatusCancelled:
		now := time.Now()
		roomLockStatus.IsActive = true
		roomLockStatus.ExpiresAt = &now
	default:
		return consts.ErrInvalidBookingStatus
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

func (s *Service) DeleteBookingByID(ctx context.Context, bookingRef models.BookingRef, id uuid.UUID) error {
	if err := s.repo.DeleteBookingByID(ctx, nil, bookingRef, id); err != nil {
		slog.ErrorContext(ctx, "failed to delete booking by id", "err", err)
		return err
	}

	return nil
}
