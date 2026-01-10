package postgres

import (
	"context"
	"errors"

	"booking/internal/repository/models"
	"booking/internal/repository/postgres/query"
	"booking/pkg/utils/consts"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) BookingCreate(ctx context.Context, b models.BookingCreate) (models.Booking, error) {
	newBooking := b.ToRead()
	insertArgs := []any{
		b.UserID,
		b.HotelID,
		b.CheckIn,
		b.CheckOut,
		b.GuestName,
		b.GuestEmail,
		b.GuestPhone,
		b.Currency,
		b.TotalAmount,
	}
	scanArgs := []any{
		&newBooking.ID,
		&newBooking.CreatedAt,
		&newBooking.UpdatedAt,
	}

	if err := r.db.QueryRow(ctx, query.BookingCreate, insertArgs...).Scan(scanArgs...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Booking{}, consts.UniqueBookingField
		}
		return models.Booking{}, err
	}

	return newBooking, nil
}

func (r *Repository) BookingGetAll(
	ctx context.Context,
	bookingRef models.BookingRef,
	limit uint64,
	offset uint64,
) (models.BookingList, error) {
	var bookingList models.BookingList
	selectArgs := []any{
		bookingRef.UserID,
		bookingRef.HotelID,
		bookingRef.Status,
		limit,
		offset,
	}

	rows, err := r.db.Query(ctx, query.BookingGetAll, selectArgs...)
	if err != nil {
		return models.BookingList{}, err
	}

	var b models.BookingShort
	for rows.Next() {
		err = rows.Scan(
			b.ID,
			b.UserID,
			b.HotelID,
			b.CheckIn,
			b.CheckOut,
			b.Status,
			b.GuestName,
			b.GuestEmail,
			b.GuestPhone,
			b.Currency,
			b.TotalAmount)
		if err != nil {
			return models.BookingList{}, err
		}

		bookingList.Booking = append(bookingList.Booking, b)
	}

	if err = r.db.
		QueryRow(
			ctx,
			query.BookingGetCountRows,
			bookingRef.UserID,
			bookingRef.HotelID,
			bookingRef.Status,
		).
		Scan(&bookingList.TotalCount); err != nil {
		return models.BookingList{}, err
	}

	return bookingList, nil
}

func (r *Repository) BookingGetByID(ctx context.Context, id uuid.UUID) (models.Booking, error) {
	var b models.Booking
	scanArgs := []any{
		b.UserID,
		b.HotelID,
		b.CheckIn,
		b.CheckOut,
		b.Status,
		b.GuestName,
		b.GuestEmail,
		b.GuestPhone,
		b.Currency,
		b.TotalAmount,
		b.CreatedAt,
		b.UpdatedAt,
	}

	if err := r.db.QueryRow(ctx, query.BookingGetByID, id).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Booking{}, consts.BookingNotFound
		}
		return models.Booking{}, err
	}

	return b, nil
}

func (r *Repository) BookingUpdateByID(ctx context.Context, id uuid.UUID, b models.BookingUpdate) error {
	updateArgs := []any{
		id,
		b.GuestName,
		b.GuestEmail,
		b.GuestPhone,
		b.CheckIn,
		b.CheckOut,
		b.TotalAmount,
	}

	row, err := r.db.Exec(ctx, query.BookingUpdateByID, updateArgs...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return consts.UniqueBookingField
		}
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingNotFound
	}

	return nil
}

func (r *Repository) BookingStatusUpdateByID(ctx context.Context, id uuid.UUID, b models.BookingStatusUpdate) error {
	updateArgs := []any{
		id,
		b.Status,
	}

	row, err := r.db.Exec(ctx, query.BookingStatusUpdateByID, updateArgs...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return consts.UniqueBookingField
		}
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingNotFound
	}

	return nil
}

func (r *Repository) BookingDeleteByID(ctx context.Context, id uuid.UUID) error {
	row, err := r.db.Exec(ctx, query.BookingDeleteByID, id)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingNotFound
	}

	return nil
}
