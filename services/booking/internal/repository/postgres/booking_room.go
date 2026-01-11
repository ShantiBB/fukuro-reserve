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

func (r *Repository) CreateBookingRoom(ctx context.Context, tx pgx.Tx, bRoom models.CreateBookingRoom) (
	models.BookingRoom, error,
) {
	db := r.executor(tx)

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

	if err := db.QueryRow(ctx, query.CreateBookingRoom, insertArgs...).Scan(scanArgs...); err != nil {
		return models.BookingRoom{}, err
	}

	return newBookingRoom, nil
}

func (r *Repository) GetBookingRoomsByBookingID(
	ctx context.Context,
	tx pgx.Tx,
	bookingID uuid.UUID,
) ([]models.BookingRoom, error) {
	db := r.executor(tx)

	var bookingRoomList []models.BookingRoom
	rows, err := db.Query(ctx, query.GetBookingRoomsByBookingID, bookingID)
	if err != nil {
		return []models.BookingRoom{}, err
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
			return []models.BookingRoom{}, err
		}

		bookingRoomList = append(bookingRoomList, bRoom)
	}

	return bookingRoomList, nil
}

func (r *Repository) GetBookingRoomByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) (models.BookingRoom, error) {
	db := r.executor(tx)

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

	if err := db.QueryRow(ctx, query.GetBookingRoomByID, id).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.BookingRoom{}, consts.BookingRoomNotFound
		}
		return models.BookingRoom{}, err
	}

	return bRoom, nil
}

func (r *Repository) UpdateBookingRoomGuestCountsByID(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
	bRoom models.BookingRoomGuestCounts,
) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.UpdateBookingRoomGuestCountsByID, id, bRoom.Adults, bRoom.Children)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingRoomNotFound
	}

	return nil
}

func (r *Repository) DeleteBookingRoomByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.DeleteBookingRoomByID, id)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingRoomNotFound
	}

	return nil
}
