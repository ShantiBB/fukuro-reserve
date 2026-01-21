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
	locks []*models.CreateRoomLock,
) ([]*models.RoomLockDetail, error) {
	db := r.executor(tx)

	roomIDs := make([]uuid.UUID, len(locks))
	bookingIDs := make([]uuid.UUID, len(locks))
	startDates := make([]time.Time, len(locks))
	endDates := make([]time.Time, len(locks))
	expiresAts := make([]time.Time, len(locks))

	for i, l := range locks {
		roomIDs[i] = l.RoomID
		bookingIDs[i] = l.BookingID
		startDates[i] = l.StayRange.Start
		endDates[i] = l.StayRange.End
		expiresAts[i] = l.ExpiresAt
	}

	rows, err := db.Query(ctx, query.CreateRoomLocks, roomIDs, bookingIDs, startDates, endDates, expiresAts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	values := make([]models.RoomLockDetail, len(locks))
	var rl models.RoomLockDetail
	var idx int32
	for rows.Next() {
		if err = rows.Scan(
			&rl.ID,
			&rl.RoomID,
			&rl.BookingID,
			&rl.ISActive,
			&rl.CreatedAt,
		); err != nil {
			return nil, err
		}

		rl.StayRange = locks[idx].StayRange
		rl.ExpiresAt = locks[idx].ExpiresAt

		values[idx] = rl
		idx++
	}
	values = values[:idx]

	if err = rows.Err(); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23P01" {
			return nil, consts.ErrRoomLockAlreadyExist
		}
		return nil, err
	}

	out := make([]*models.RoomLockDetail, len(values))
	for i := range values {
		out[i] = &values[i]
	}

	return out, nil
}

func (r *Repository) UpdateRoomLockActivityByID(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
	roomLock *models.RoomLockActivity,
) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.UpdateRoomLockActivityByID, id, roomLock.IsActive, roomLock.ExpiresAt)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return consts.ErrRoomLockNotFound
	}

	return nil
}

func (r *Repository) DeleteRoomLockByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.DeleteRoomLockByID, id)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return consts.ErrRoomLockNotFound
	}

	return nil
}
