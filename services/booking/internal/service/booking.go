package service

import (
	"context"
	"fmt"
	"time"

	"booking/internal/repository/models"
	"booking/internal/service/utils/helper"
	"booking/pkg/utils/consts"

	"github.com/google/uuid"
)

func (s *Service) BookingCreate(
	ctx context.Context,
	b models.CreateBooking,
	rooms []models.CreateBookingRoom,
) (models.Booking, error) {
	var err error
	b.FinalTotalAmount, err = helper.CalculateFinalTotalAmount(b.CheckIn, b.CheckOut, rooms, b.ExpectedTotalAmount)
	if err != nil {
		return models.Booking{}, err
	}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return models.Booking{}, err
	}
	defer tx.Rollback(ctx)

	newBooking, err := s.repo.CreateBooking(ctx, tx, b)
	if err != nil {
		return models.Booking{}, err
	}

	for i := range rooms {
		rooms[i].BookingID = newBooking.ID
	}

	newRooms, err := s.repo.CreateBookingRooms(ctx, tx, rooms)
	if err != nil {
		return models.Booking{}, err
	}

	newBooking.BookingRooms = newRooms

	locks := make([]models.CreateRoomLock, 0, len(newRooms))
	for _, nr := range newRooms {
		locks = append(
			locks, models.CreateRoomLock{
				RoomID:    nr.RoomID,
				BookingID: newBooking.ID,
				StayRange: models.DateRange{
					Start: b.CheckIn,
					End:   b.CheckOut,
				},
				ExpiresAt: time.Now().Add(consts.ExpireRoomLockMinutes * time.Minute),
			},
		)
	}

	if _, err = s.repo.CreateRoomLocks(ctx, tx, locks); err != nil {
		return models.Booking{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return models.Booking{}, err
	}

	return newBooking, nil
}

func (s *Service) GetBookings(
	ctx context.Context,
	bookingRef models.BookingRef,
	page uint64,
	limit uint64,
) (models.BookingList, error) {
	offset := (page - 1) * limit

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return models.BookingList{}, err
	}
	defer tx.Rollback(ctx)

	bookingList, err := s.repo.GetBookingsByHotelInfo(ctx, tx, bookingRef, limit, offset)
	if err != nil {
		return models.BookingList{}, err
	}

	bookingIDs := make([]uuid.UUID, len(bookingList.Bookings))
	for i, booking := range bookingList.Bookings {
		bookingIDs[i] = booking.ID
	}

	allRooms, err := s.repo.GetBookingRoomsInfoByBookingIDs(ctx, tx, bookingIDs)
	if err != nil {
		return models.BookingList{}, fmt.Errorf("get booking rooms: %w", err)
	}

	roomsByBookingID := make(map[uuid.UUID][]models.BookingRoomInfo)
	for _, room := range allRooms {
		roomsByBookingID[room.BookingID] = append(roomsByBookingID[room.BookingID], room)
	}

	for i := range bookingList.Bookings {
		bookingList.Bookings[i].BookingRooms = roomsByBookingID[bookingList.Bookings[i].ID]
	}

	return bookingList, nil
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
