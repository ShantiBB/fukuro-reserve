package postgres

import (
	"context"
	"errors"
	"hotel/internal/repository/models"
	"hotel/internal/repository/postgres/query"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"fukuro-reserve/pkg/utils/consts"
)

func (r *Repository) RoomCreate(
	ctx context.Context,
	hotel models.HotelRef,
	room models.RoomCreate,
) (models.Room, error) {
	newRoom := room.ToRead()
	insertArgs := []any{
		hotel.CountryCode,
		hotel.CitySlug,
		hotel.HotelSlug,
		room.Title,
		room.Description,
		room.RoomNumber,
		room.Type,
		room.Price,
		room.Capacity,
		room.AreaSqm,
		room.Floor,
		room.Amenities,
		room.Images,
	}
	scanArgs := []any{
		&newRoom.ID,
		&newRoom.Status,
		&newRoom.CreatedAt,
		&newRoom.UpdatedAt,
	}

	if err := r.db.QueryRow(ctx, query.RoomCreateQuery, insertArgs...).Scan(scanArgs...); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.Room{}, consts.UniqueRoomField
		}
		return models.Room{}, err
	}

	return newRoom, nil
}

func (r *Repository) RoomGetByID(
	ctx context.Context,
	hotel models.HotelRef,
	roomID uuid.UUID,
) (models.Room, error) {
	var room models.Room
	insertArgs := []any{
		hotel.CountryCode,
		hotel.CitySlug,
		hotel.HotelSlug,
		roomID,
	}
	scanArgs := []any{
		&room.ID,
		&room.Title,
		&room.Description,
		&room.RoomNumber,
		&room.Type,
		&room.Status,
		&room.Price,
		&room.Capacity,
		&room.AreaSqm,
		&room.Floor,
		&room.Amenities,
		&room.Images,
		&room.CreatedAt,
		&room.UpdatedAt,
	}

	if err := r.db.QueryRow(ctx, query.RoomGetByID, insertArgs...).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Room{}, consts.RoomNotFound
		}
		return models.Room{}, err
	}

	return room, nil
}

func (r *Repository) RoomGetAll(
	ctx context.Context,
	hotel models.HotelRef,
	limit,
	offset uint64,
) (models.RoomList, error) {
	rows, err := r.db.Query(
		ctx,
		query.RoomGetAll,
		hotel.CountryCode,
		hotel.CitySlug,
		hotel.HotelSlug,
		limit,
		offset,
	)
	if err != nil {
		return models.RoomList{}, err
	}

	var roomList models.RoomList
	var room models.RoomShort

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&room.ID,
			&room.Title,
			&room.RoomNumber,
			&room.Type,
			&room.Status,
			&room.Price,
			&room.Capacity,
			&room.AreaSqm,
			&room.Amenities,
			&room.Images)
		if err != nil {
			return models.RoomList{}, err
		}

		roomList.Rooms = append(roomList.Rooms, room)
	}

	if err = rows.Err(); err != nil {
		return models.RoomList{}, err
	}

	if err = r.db.QueryRow(ctx, query.RoomGetCountRows).Scan(&roomList.TotalCount); err != nil {
		return models.RoomList{}, err
	}

	return roomList, nil
}

func (r *Repository) RoomUpdateByID(
	ctx context.Context,
	hotel models.HotelRef,
	roomID uuid.UUID,
	room models.RoomUpdate,
) error {
	row, err := r.db.Exec(
		ctx, query.RoomUpdateByID,
		hotel.CountryCode,
		hotel.CitySlug,
		hotel.HotelSlug,
		roomID,
		room.Title,
		room.Description,
		room.RoomNumber,
		room.Type,
		room.Price,
		room.Capacity,
		room.AreaSqm,
		room.Floor,
		room.Amenities,
		room.Images,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return consts.UniqueRoomField
		}
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.RoomNotFound
	}

	return nil
}

func (r *Repository) RoomDeleteByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID) error {
	row, err := r.db.Exec(
		ctx,
		query.RoomDeleteByID,
		hotel.CountryCode,
		hotel.CitySlug,
		hotel.HotelSlug,
		roomID,
	)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.RoomNotFound
	}

	return nil
}
