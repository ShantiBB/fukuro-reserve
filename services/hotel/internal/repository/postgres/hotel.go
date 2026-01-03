package postgres

import (
	"context"
	"errors"
	"fmt"
	"hotel/internal/repository/models"
	"hotel/internal/repository/postgres/query"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"fukuro-reserve/pkg/utils/consts"
)

func (r *Repository) HotelCreate(ctx context.Context, h models.HotelCreate) (models.Hotel, error) {
	newHotel := h.ToRead()
	insertArgs := []any{
		h.Name,
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
			return models.Hotel{}, errors.New("username or email already exists")
		}
		return models.Hotel{}, err
	}

	return newHotel, nil
}

func (r *Repository) HotelGetByIDOrName(ctx context.Context, field any) (models.Hotel, error) {
	var hotel models.Hotel
	scanArgs := []any{
		&hotel.ID,
		&hotel.Name,
		&hotel.OwnerID,
		&hotel.Description,
		&hotel.Address,
		&hotel.Location.Longitude,
		&hotel.Location.Latitude,
		&hotel.Rating,
		&hotel.CreatedAt,
		&hotel.UpdatedAt,
	}

	var q string
	switch v := field.(type) {
	case uuid.UUID:
		q = query.HotelGetByID
	case string:
		q = query.HotelGetByName
	default:
		return models.Hotel{}, fmt.Errorf("unsupported type %T", v)
	}

	if err := r.db.QueryRow(ctx, q, field).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Hotel{}, errors.New("hotel not found")
		}
		return models.Hotel{}, err
	}

	return hotel, nil
}

func (r *Repository) HotelGetAll(ctx context.Context, limit, offset uint64) (models.HotelList, error) {
	var hotelList models.HotelList

	rows, err := r.db.Query(ctx, query.HotelGetAll, limit, offset)
	if err != nil {
		return models.HotelList{}, err
	}

	var h models.HotelShort
	for rows.Next() {
		err = rows.Scan(
			&h.ID, &h.Name, &h.OwnerID, &h.Address, &h.Rating, &h.Location.Longitude, &h.Location.Latitude,
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

func (r *Repository) HotelUpdateByID(ctx context.Context, id uuid.UUID, h models.HotelUpdate) error {
	row, err := r.db.Exec(
		ctx, query.HotelUpdateByID,
		h.Name,
		h.Description,
		h.Address,
		h.Location.Longitude,
		h.Location.Latitude,
		id,
	)
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

func (r *Repository) HotelDeleteByID(ctx context.Context, id uuid.UUID) error {
	row, err := r.db.Exec(ctx, query.HotelDeleteByID, id)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.HotelNotFound
	}

	return nil
}
