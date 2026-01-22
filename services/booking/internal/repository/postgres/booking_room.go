package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"booking/internal/repository/models"
	"booking/internal/repository/postgres/query"
	"booking/internal/utils/consts"
)

func (r *Repository) CreateBookingRooms(
	ctx context.Context,
	tx pgx.Tx,
	bookingID uuid.UUID,
	rooms []*models.CreateBookingRoom,
) ([]*models.BookingRoomWithLock, error) {
	db := r.executor(tx)

	roomIDs := make([]uuid.UUID, len(rooms))
	adults := make([]uint32, len(rooms))
	children := make([]uint32, len(rooms))
	prices := make([]string, len(rooms))

	for i, room := range rooms {
		if room.BookingID != bookingID {
			return nil, consts.ErrConflictBookingRooms
		}

		roomIDs[i] = room.RoomID
		adults[i] = room.Adults
		children[i] = room.Children
		prices[i] = room.PricePerNight.StringFixed(2)
	}

	rows, err := db.Query(ctx, query.CreateBookingRooms, bookingID, roomIDs, adults, children, prices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	values := make([]models.BookingRoomWithLock, len(rooms))
	var br models.BookingRoomWithLock
	var idx int
	for rows.Next() {
		if err = rows.Scan(&br.ID, &br.CreatedAt); err != nil {
			return nil, err
		}
		br.RoomID = rooms[idx].RoomID
		br.Adults = rooms[idx].Adults
		br.Children = rooms[idx].Children
		br.PricePerNight = rooms[idx].PricePerNight
		values[idx] = br
		idx++
	}
	values = values[:idx]

	if err = rows.Err(); err != nil {
		return nil, err
	}

	out := make([]*models.BookingRoomWithLock, len(values))
	for i := range values {
		out[i] = &values[i]
	}

	return out, nil
}

func (r *Repository) GetBookingRoomsByBookingIDs(
	ctx context.Context,
	tx pgx.Tx,
	bookingIDs []uuid.UUID,
) ([]*models.BookingRoom, error) {
	db := r.executor(tx)

	var values []models.BookingRoom
	rows, err := db.Query(ctx, query.GetBookingRoomsByBookingID, bookingIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bRoom models.BookingRoom
	for rows.Next() {
		err = rows.Scan(
			&bRoom.ID,
			&bRoom.BookingID,
			&bRoom.RoomID,
			&bRoom.Adults,
			&bRoom.Children,
			&bRoom.PricePerNight,
		)
		if err != nil {
			return nil, err
		}

		values = append(values, bRoom)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	out := make([]*models.BookingRoom, len(values))
	for i := range values {
		out[i] = &values[i]
	}

	return out, nil
}

func (r *Repository) GetBookingRoomsWithLockByBookingIDs(
	ctx context.Context,
	tx pgx.Tx,
	bookingIDs []uuid.UUID,
) ([]*models.BookingRoomWithLock, error) {
	db := r.executor(tx)

	var values []models.BookingRoomWithLock
	rows, err := db.Query(ctx, query.GetBookingRoomsWithLockByBookingID, bookingIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bRoom models.BookingRoomWithLock
	for rows.Next() {
		err = rows.Scan(
			&bRoom.ID,
			&bRoom.RoomID,
			&bRoom.Adults,
			&bRoom.Children,
			&bRoom.PricePerNight,
			&bRoom.CreatedAt,
			&bRoom.RoomLock.ID,
			&bRoom.RoomLock.ISActive,
			&bRoom.RoomLock.ExpiresAt,
			&bRoom.RoomLock.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		values = append(values, bRoom)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	out := make([]*models.BookingRoomWithLock, len(values))
	for i := range values {
		out[i] = &values[i]
	}

	return out, nil
}
