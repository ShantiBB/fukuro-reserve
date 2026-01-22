package postgres

import (
	"context"
	"errors"

	"booking/internal/repository/models"
	"booking/internal/repository/postgres/query"
	"booking/pkg/lib/utils/consts"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateBooking(ctx context.Context, tx pgx.Tx, b *models.CreateBooking) (*models.Booking, error) {
	if b == nil {
		return nil, consts.ErrNilObject
	}

	db := r.executor(tx)

	newBooking := b.ToRead()
	err := db.QueryRow(
		ctx,
		query.CreateBooking,
		b.UserID,
		b.HotelID,
		b.CheckIn,
		b.CheckOut,
		b.GuestName,
		b.GuestEmail,
		b.GuestPhone,
		b.Currency,
		b.ExpectedTotalAmount,
		b.FinalTotalAmount,
	).Scan(
		&newBooking.ID,
		&newBooking.Status,
		&newBooking.CreatedAt,
		&newBooking.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return newBooking, nil
}

func (r *Repository) GetBookingsByHotelInfo(
	ctx context.Context,
	tx pgx.Tx,
	bookingRef models.BookingRef,
	limit uint64,
	offset uint64,
) (*models.BookingList, error) {
	db := r.executor(tx)

	rows, err := db.Query(
		ctx,
		query.GetBookingsByHotelInfo,
		bookingRef.UserID,
		bookingRef.HotelID,
		bookingRef.Status,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	values := make([]models.BookingShort, limit)
	var b models.BookingShort
	var idx int
	for rows.Next() {
		err = rows.Scan(
			&b.ID,
			&b.UserID,
			&b.HotelID,
			&b.CheckIn,
			&b.CheckOut,
			&b.Status,
			&b.GuestName,
			&b.GuestEmail,
			&b.GuestPhone,
			&b.Currency,
			&b.ExpectedTotalAmount,
			&b.FinalTotalAmount,
		)
		if err != nil {
			return nil, err
		}
		values[idx] = b
		idx++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	values = values[:idx]

	bookingList := &models.BookingList{
		Bookings: make([]*models.BookingShort, len(values)),
	}
	for i := range values {
		bookingList.Bookings[i] = &values[i]
	}

	if err = db.QueryRow(
		ctx,
		query.GetBookingCountRows,
		bookingRef.UserID,
		bookingRef.HotelID,
		bookingRef.Status,
	).Scan(&bookingList.TotalCount); err != nil {
		return nil, err
	}

	return bookingList, nil
}

func (r *Repository) GetBookingByID(ctx context.Context, tx pgx.Tx, bookingID uuid.UUID) (*models.Booking, error) {
	db := r.executor(tx)

	var b models.Booking
	err := db.QueryRow(ctx, query.GetBookingByID, bookingID).Scan(
		&b.ID,
		&b.UserID,
		&b.HotelID,
		&b.CheckIn,
		&b.CheckOut,
		&b.Status,
		&b.GuestName,
		&b.GuestEmail,
		&b.GuestPhone,
		&b.Currency,
		&b.ExpectedTotalAmount,
		&b.FinalTotalAmount,
		&b.CreatedAt,
		&b.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, consts.ErrBookingNotFound
		}
		return nil, err
	}

	return &b, nil
}

func (r *Repository) UpdateBookingGuestInfoByID(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
	b *models.UpdateBooking,
) error {
	db := r.executor(tx)

	row, err := db.Exec(
		ctx,
		query.UpdateBookingGuestInfoByID,
		id,
		b.GuestName,
		b.GuestEmail,
		b.GuestPhone,
	)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.ErrBookingNotFound
	}

	return nil
}

func (r *Repository) UpdateBookingStatusByID(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
	status models.BookingStatus,
) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.UpdateBookingStatusByID, id, status)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return consts.ErrBookingNotFound
	}

	return nil
}

func (r *Repository) DeleteBookingByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.DeleteBookingByID, id)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return consts.ErrBookingNotFound
	}

	return nil
}
