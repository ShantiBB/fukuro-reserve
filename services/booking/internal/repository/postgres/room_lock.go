package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"booking/internal/repository/models"
	"booking/internal/repository/postgres/query"
	"booking/pkg/utils/consts"
)

func (r *Repository) CreateRoomLocks(
	ctx context.Context,
	tx pgx.Tx,
	locks []models.CreateRoomLock,
) ([]models.RoomLock, error) {

	if len(locks) == 0 {
		return nil, nil
	}

	db := r.executor(tx)

	roomIDs := make([]uuid.UUID, 0, len(locks))
	bookingIDs := make([]uuid.UUID, 0, len(locks))
	startDates := make([]time.Time, 0, len(locks))
	endDates := make([]time.Time, 0, len(locks))
	expiresAts := make([]time.Time, 0, len(locks))

	for _, l := range locks {
		roomIDs = append(roomIDs, l.RoomID)
		bookingIDs = append(bookingIDs, l.BookingID)
		startDates = append(startDates, l.StayRange.Start)
		endDates = append(endDates, l.StayRange.End)
		expiresAts = append(expiresAts, l.ExpiresAt)
	}

	rows, err := db.Query(ctx, query.CreateRoomLocks, roomIDs, bookingIDs, startDates, endDates, expiresAts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]models.RoomLock, 0, len(locks))
	for rows.Next() {
		var rl models.RoomLock
		if err = rows.Scan(
			&rl.ID,
			&rl.RoomID,
			&rl.BookingID,
			&rl.ISActive,
			&rl.CreatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, rl)
	}

	if err = rows.Err(); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23P01" {
			return nil, consts.RoomLockAlreadyExist
		}
		return nil, err
	}

	return out, nil
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
