package postgres

import (
	"context"
	"errors"
	"hotel/internal/repository/models"
	"hotel/internal/repository/postgres/query"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"fukuro-reserve/pkg/utils/consts"
)

func (r *Repository) HotelCreate(ctx context.Context, h models.HotelCreate) (models.Hotel, error) {
	newHotel := h.ToRead()
	insertArgs := []any{
		h.CountryCode,
		h.CitySlug,
		h.Title,
		h.Slug,
		h.OwnerID,
		h.Description,
		h.Address,
		h.Location.Longitude,
		h.Location.Latitude,
	}
	scanArgs := []any{
		&newHotel.ID,
		&newHotel.CreatedAt,
		&newHotel.UpdatedAt,
	}

	if err := r.db.QueryRow(ctx, query.HotelCreateQuery, insertArgs...).Scan(scanArgs...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Hotel{}, consts.UniqueHotelField
		}
		return models.Hotel{}, err
	}

	return newHotel, nil
}

func (r *Repository) HotelGetBySlug(ctx context.Context, h models.Hotel) (models.Hotel, error) {
	insertArgs := []any{
		h.CountryCode,
		h.CitySlug,
		h.Slug,
	}
	scanArgs := []any{
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
	}

	if err := r.db.QueryRow(ctx, query.HotelGetBySlug, insertArgs...).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Hotel{}, consts.HotelNotFound
		}
		return models.Hotel{}, err
	}

	return h, nil
}

func (r *Repository) HotelGetAll(ctx context.Context, filter models.HotelFilter) (models.HotelList, error) {
	var hotelList models.HotelList
	selectArgs := []any{
		filter.CountryCode,
		filter.CitySlug,
		filter.SortField,
		filter.Limit,
		filter.Offset,
	}

	rows, err := r.db.Query(ctx, query.HotelGetAll, selectArgs...)
	if err != nil {
		return models.HotelList{}, err
	}

	var h models.HotelShort
	for rows.Next() {
		err = rows.Scan(
			&h.ID,
			&h.Title,
			&h.Slug,
			&h.OwnerID,
			&h.Address,
			&h.Rating,
			&h.Location.Longitude,
			&h.Location.Latitude,
		)
		if err != nil {
			return models.HotelList{}, err
		}

		hotelList.Hotels = append(hotelList.Hotels, h)
	}

	if err = r.db.QueryRow(ctx, query.HotelGetCountRows).Scan(&hotelList.TotalCount); err != nil {
		return models.HotelList{}, err
	}

	return hotelList, nil
}

func (r *Repository) HotelUpdateBySlug(ctx context.Context, hotelSlug string, h models.HotelUpdate) error {
	updateArgs := []any{
		h.Title,
		h.Slug,
		h.Description,
		h.Address,
		h.Location.Longitude,
		h.Location.Latitude,
		h.CountryCode,
		h.CitySlug,
		hotelSlug,
	}

	row, err := r.db.Exec(ctx, query.HotelUpdateBySlug, updateArgs...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return consts.UniqueHotelField
		}
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.HotelNotFound
	}

	return nil
}

func (r *Repository) HotelDeleteBySlug(ctx context.Context, countryCode, citySlug, slug string) error {
	row, err := r.db.Exec(ctx, query.HotelDeleteBySlug, countryCode, citySlug, slug)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.HotelNotFound
	}

	return nil
}
