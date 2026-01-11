package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"booking/internal/repository/models"
	"booking/internal/repository/postgres/query"
	"booking/pkg/utils/consts"
)

func (r *Repository) CreateBookingRoom(ctx context.Context, bRoom models.BookingRoomCreate) (
	models.BookingRoom, error,
) {
	newBookingRoom := bRoom.ToRead()
	insertArgs := []any{
		bRoom.BookingID,
		bRoom.RoomID,
		bRoom.Adults,
		bRoom.Children,
		bRoom.PricePerNight,
	}
	scanArgs := []any{
		&newBookingRoom.ID,
		&newBookingRoom.CreatedAt,
	}

	if err := r.db.QueryRow(ctx, query.CreateBookingRoom, insertArgs...).Scan(scanArgs...); err != nil {
		return models.BookingRoom{}, err
	}

	return newBookingRoom, nil
}

func (r *Repository) GetBookingRoomsByBookingID(
	ctx context.Context,
	bookingID uuid.UUID,
) (models.BookingRoomList, error) {
	var bookingRoomList models.BookingRoomList
	rows, err := r.db.Query(ctx, query.GetBookingRoomsByBookingID, bookingID)
	if err != nil {
		return models.BookingRoomList{}, err
	}

	var bRoom models.BookingRoom
	for rows.Next() {
		err = rows.Scan(
			bRoom.ID,
			bRoom.BookingID,
			bRoom.RoomID,
			bRoom.Adults,
			bRoom.Children,
			bRoom.PricePerNight,
			bRoom.CreatedAt,
		)
		if err != nil {
			return models.BookingRoomList{}, err
		}

		bookingRoomList.BookingRooms = append(bookingRoomList.BookingRooms, bRoom)
	}

	return bookingRoomList, nil
}

func (r *Repository) GetBookingRoomByID(ctx context.Context, id uuid.UUID) (models.BookingRoom, error) {
	var bRoom models.BookingRoom
	scanArgs := []any{
		bRoom.ID,
		bRoom.BookingID,
		bRoom.RoomID,
		bRoom.Adults,
		bRoom.Children,
		bRoom.PricePerNight,
		bRoom.CreatedAt,
	}

	if err := r.db.QueryRow(ctx, query.GetBookingRoomByID, id).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.BookingRoom{}, consts.BookingRoomNotFound
		}
		return models.BookingRoom{}, err
	}

	return bRoom, nil
}

func (r *Repository) UpdateBookingRoomGuestCountsByID(
	ctx context.Context,
	id uuid.UUID,
	bRoom models.BookingRoomGuestCounts,
) error {
	row, err := r.db.Exec(ctx, query.UpdateBookingRoomGuestCountsByID, id, bRoom.Adults, bRoom.Children)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingRoomNotFound
	}

	return nil
}

func (r *Repository) DeleteBookingRoomByID(ctx context.Context, id uuid.UUID) error {
	row, err := r.db.Exec(ctx, query.DeleteBookingRoomByID, id)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingRoomNotFound
	}

	return nil
}
