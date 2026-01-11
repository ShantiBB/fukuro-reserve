package postgres

import (
	"context"
	"errors"

	"booking/internal/repository/models"
	"booking/internal/repository/postgres/query"
	"booking/pkg/utils/consts"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateBooking(ctx context.Context, b models.CreateBooking) (models.Booking, error) {
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

	if err := r.db.QueryRow(ctx, query.CreateBooking, insertArgs...).Scan(scanArgs...); err != nil {
		return models.Booking{}, err
	}

	return newBooking, nil
}

func (r *Repository) GetBookingsByHotelInfo(
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

	rows, err := r.db.Query(ctx, query.GetBookingsByHotelInfo, selectArgs...)
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
			b.TotalAmount,
		)
		if err != nil {
			return models.BookingList{}, err
		}

		bookingList.Bookings = append(bookingList.Bookings, b)
	}

	if err = r.db.
		QueryRow(
			ctx,
			query.GetBookingCountRows,
			bookingRef.UserID,
			bookingRef.HotelID,
			bookingRef.Status,
		).
		Scan(&bookingList.TotalCount); err != nil {
		return models.BookingList{}, err
	}

	return bookingList, nil
}

func (r *Repository) GetBookingByID(ctx context.Context, id uuid.UUID) (models.Booking, error) {
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

	if err := r.db.QueryRow(ctx, query.GetBookingByID, id).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Booking{}, consts.BookingNotFound
		}
		return models.Booking{}, err
	}

	return b, nil
}

func (r *Repository) UpdateBookingGuestInfoByID(ctx context.Context, id uuid.UUID, b models.UpdateBooking) error {
	updateArgs := []any{
		id,
		b.GuestName,
		b.GuestEmail,
		b.GuestPhone,
	}

	row, err := r.db.Exec(ctx, query.UpdateBookingGuestInfoByID, updateArgs...)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingNotFound
	}

	return nil
}

func (r *Repository) UpdateBookingStatusByID(ctx context.Context, id uuid.UUID, b models.BookingStatusInfo) error {
	updateArgs := []any{
		id,
		b.Status,
	}

	row, err := r.db.Exec(ctx, query.UpdateBookingStatusByID, updateArgs...)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingNotFound
	}

	return nil
}

func (r *Repository) DeleteBookingByID(ctx context.Context, id uuid.UUID) error {
	row, err := r.db.Exec(ctx, query.DeleteBookingByID, id)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingNotFound
	}

	return nil
}
