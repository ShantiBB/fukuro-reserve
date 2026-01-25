package postgres

import (
	"context"
	"errors"

	"hotel/internal/repository/models"
	"hotel/internal/repository/postgres/query"
	"hotel/pkg/lib/utils/consts"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) InsertHotel(ctx context.Context, h *models.CreateHotel) (*models.Hotel, error) {
	newHotel := h.ToRead()
	err := r.db.QueryRow(
		ctx, query.CreateHotelQuery,
		h.CountryCode,
		h.CitySlug,
		h.Title,
		h.HotelSlug,
		h.OwnerID,
		h.Description,
		h.Address,
		h.Location.Longitude,
		h.Location.Latitude,
	).Scan(
		&newHotel.ID,
		&newHotel.CreatedAt,
		&newHotel.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, consts.ErrUniqueHotelField
		}
		return nil, err
	}

	return newHotel, nil
}

func (r *Repository) SelectHotels(
	ctx context.Context,
	ref models.HotelRef,
	sortField string,
	limit uint64,
	offset uint64,
) (*models.HotelList, error) {
	rows, err := r.db.Query(
		ctx,
		query.GetHotels,
		ref.CountryCode,
		ref.CitySlug,
		sortField,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	values := make([]models.HotelShort, limit)
	var h models.HotelShort
	var totalCount, idx uint64
	for rows.Next() {
		err = rows.Scan(
			&h.ID,
			&h.Title,
			&h.HotelSlug,
			&h.OwnerID,
			&h.Address,
			&h.Rating,
			&h.Location.Longitude,
			&h.Location.Latitude,
			&totalCount,
		)
		if err != nil {
			return nil, err
		}

		values[idx] = h
		idx++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	values = values[:idx]

	hotelList := &models.HotelList{
		Hotels:     make([]*models.HotelShort, len(values)),
		TotalCount: totalCount,
	}
	for i := range values {
		hotelList.Hotels[i] = &values[i]
	}

	return hotelList, nil
}

func (r *Repository) SelectHotelBySlug(ctx context.Context, ref models.HotelRef) (*models.Hotel, error) {
	var h models.Hotel
	err := r.db.QueryRow(
		ctx, query.GetHotelBySlug,
		ref.CountryCode,
		ref.CitySlug,
		ref.HotelSlug,
	).Scan(
		&h.ID,
		&h.Title,
		&h.OwnerID,
		&h.Description,
		&h.Address,
		&h.Location.Longitude,
		&h.Location.Latitude,
		&h.Rating,
		&h.CreatedAt,
		&h.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, consts.ErrHotelNotFound
		}
		return nil, err
	}

	return &h, nil
}

func (r *Repository) UpdateHotelBySlug(ctx context.Context, ref models.HotelRef, h models.UpdateHotel) error {
	row, err := r.db.Exec(
		ctx, query.UpdateHotelBySlug,
		h.Description,
		h.Address,
		h.Location.Longitude,
		h.Location.Latitude,
		ref.CountryCode,
		ref.CitySlug,
		ref.HotelSlug,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return consts.ErrUniqueHotelField
		}
		return err
	}
	if row.RowsAffected() == 0 {
		return consts.ErrHotelNotFound
	}

	return nil
}

func (r *Repository) UpdateHotelTitleBySlug(
	ctx context.Context,
	ref models.HotelRef,
	h models.UpdateHotelTitle,
) error {
	row, err := r.db.Exec(
		ctx, query.UpdateHotelTitleBySlug,
		h.Title,
		h.HotelSlug,
		ref.CountryCode,
		ref.CitySlug,
		ref.HotelSlug,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return consts.ErrUniqueHotelField
		}
		return err
	}
	if row.RowsAffected() == 0 {
		return consts.ErrHotelNotFound
	}

	return nil
}

func (r *Repository) DeleteHotelBySlug(ctx context.Context, ref models.HotelRef) error {
	row, err := r.db.Exec(ctx, query.DeleteHotelBySlug, ref.CountryCode, ref.CitySlug, ref.HotelSlug)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return consts.ErrHotelNotFound
	}

	return nil
}
