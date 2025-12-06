package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"hotel/internal/repository/models"
)

func (r *Repository) HotelCreate(ctx context.Context, h *models.HotelCreate) (*models.Hotel, error) {
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

	if err := r.db.QueryRow(ctx, hotelCreateQuery, insertArgs...).Scan(scanArgs...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, errors.New("username or email already exists")
		}
		return nil, err
	}

	return &newHotel, nil
}

func (r *Repository) HotelGetByIDOrName(ctx context.Context, field any) (*models.Hotel, error) {
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

	var query string
	switch v := field.(type) {
	case uuid.UUID:
		query = hotelGetByID
	case string:
		query = hotelGetByName
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}

	if err := r.db.QueryRow(ctx, query, field).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("hotel not found")
		}
		return nil, err
	}

	return &hotel, nil
}

func (r *Repository) HotelGetAll(ctx context.Context, page, pageSize int) ([]models.HotelShort, error) {
	var hotels []models.HotelShort

	offset := (page - 1) * pageSize
	rows, err := r.db.Query(ctx, hotelGetAll, pageSize, offset)
	if err != nil {
		return nil, err
	}

	var h models.HotelShort
	for rows.Next() {
		err = rows.Scan(
			&h.ID, &h.Name, &h.OwnerID, &h.Address, &h.Rating, &h.Location.Longitude, &h.Location.Latitude,
		)
		if err != nil {
			return nil, err
		}

		hotels = append(hotels, h)
	}

	return hotels, nil
}
