package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"booking/internal/repository/models"
	"booking/internal/repository/postgres/query"
	"booking/pkg/utils/consts"
)

func (r *Repository) CreateRoomLock(
	ctx context.Context,
	tx pgx.Tx,
	roomLock models.CreateRoomLock,
) (models.RoomLock, error) {
	db := r.executor(tx)

	newRoomLock := roomLock.ToRead()
	insertArgs := []any{
		roomLock.RoomID,
		roomLock.BookingID,
		roomLock.StayRange.Start,
		roomLock.StayRange.End,
		roomLock.ExpiresAt,
	}
	scanArgs := []any{
		&newRoomLock.ID,
		&newRoomLock.ISActive,
		&newRoomLock.CreatedAt,
	}

	if err := db.QueryRow(ctx, query.CreateRoomLock, insertArgs...).Scan(scanArgs...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23P01" {
			return models.RoomLock{}, consts.RoomLockAlreadyExist
		}
		return models.RoomLock{}, err
	}

	return newRoomLock, nil
}

func (r *Repository) GetRoomsLockByBookingID(
	ctx context.Context,
	tx pgx.Tx,
	bookingID uuid.UUID,
) ([]models.RoomLock, error) {
	db := r.executor(tx)

	var roomLockList []models.RoomLock
	rows, err := db.Query(ctx, query.GetRoomsLockByBookingID, bookingID)
	if err != nil {
		return []models.RoomLock{}, err
	}

	var roomLock models.RoomLock
	for rows.Next() {
		err = rows.Scan(
			roomLock.ID,
			roomLock.RoomID,
			roomLock.BookingID,
			roomLock.StayRange,
			roomLock.ExpiresAt,
			roomLock.ISActive,
			roomLock.CreatedAt,
		)
		if err != nil {
			return []models.RoomLock{}, err
		}

		roomLockList = append(roomLockList, roomLock)
	}

	return roomLockList, nil
}

func (r *Repository) GetRoomLockByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) (models.RoomLock, error) {
	db := r.executor(tx)

	var roomLock models.RoomLock
	scanArgs := []any{
		roomLock.ID,
		roomLock.RoomID,
		roomLock.BookingID,
		roomLock.StayRange,
		roomLock.ExpiresAt,
		roomLock.ISActive,
		roomLock.CreatedAt,
	}

	if err := db.QueryRow(ctx, query.GetRoomLockByID, id).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.RoomLock{}, consts.RoomLockNotFound
		}
		return models.RoomLock{}, err
	}

	return roomLock, nil
}

func (r *Repository) UpdateRoomLockActivityByID(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
	roomLock models.UpdateRoomLockActivity,
) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.UpdateRoomLockActivityByID, id, roomLock.IsActive, roomLock.ExpiresAt)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.RoomLockNotFound
	}

	return nil
}

func (r *Repository) DeleteRoomLockByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.DeleteRoomLockByID, id)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.RoomLockNotFound
	}

	return nil
}
