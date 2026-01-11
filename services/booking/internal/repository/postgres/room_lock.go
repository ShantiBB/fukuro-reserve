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

func (r *Repository) CreateRoomLock(
	ctx context.Context,
	roomLock models.CreateRoomLock,
) (models.RoomLock, error) {
	newRoomLock := roomLock.ToRead()
	insertArgs := []any{
		roomLock.RoomID,
		roomLock.BookingID,
		roomLock.StayRange,
		roomLock.ExpiresAt,
	}
	scanArgs := []any{
		&newRoomLock.ID,
		&newRoomLock.ISActive,
		&newRoomLock.CreatedAt,
	}

	if err := r.db.QueryRow(ctx, query.CreateRoomLock, insertArgs...).Scan(scanArgs...); err != nil {
		return models.RoomLock{}, err
	}

	return newRoomLock, nil
}

func (r *Repository) GetRoomsLockByBookingID(
	ctx context.Context,
	bookingID uuid.UUID,
) (models.RoomLockList, error) {
	var roomLockList models.RoomLockList
	rows, err := r.db.Query(ctx, query.GetRoomsLockByBookingID, bookingID)
	if err != nil {
		return models.RoomLockList{}, err
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
			return models.RoomLockList{}, err
		}

		roomLockList.RoomsLock = append(roomLockList.RoomsLock, roomLock)
	}

	return roomLockList, nil
}

func (r *Repository) GetRoomLockByID(ctx context.Context, id uuid.UUID) (models.RoomLock, error) {
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

	if err := r.db.QueryRow(ctx, query.GetRoomLockByID, id).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.RoomLock{}, consts.RoomLockNotFound
		}
		return models.RoomLock{}, err
	}

	return roomLock, nil
}

func (r *Repository) UpdateRoomLockActivityByID(
	ctx context.Context,
	id uuid.UUID,
	roomLock models.UpdateRoomLockActivity,
) error {
	row, err := r.db.Exec(ctx, query.UpdateRoomLockActivityByID, id, roomLock.IsActive, roomLock.ExpiresAt)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.RoomLockNotFound
	}

	return nil
}

func (r *Repository) DeleteRoomLockByID(ctx context.Context, id uuid.UUID) error {
	row, err := r.db.Exec(ctx, query.DeleteRoomLockByID, id)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.RoomLockNotFound
	}

	return nil
}
