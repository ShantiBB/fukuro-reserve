package postgres

import (
	"context"
	"errors"

	"hotel/internal/repository/models"
	"hotel/internal/repository/postgres/query"
	"hotel/pkg/lib/utils/consts"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) InsertRoom(
	ctx context.Context,
	hotelRef models.HotelRef,
	room *models.CreateRoom,
) (*models.Room, error) {
	newRoom := room.ToRead()
	err := r.db.QueryRow(
		ctx, query.InsertRoomQuery,
		hotelRef.CountryCode,
		hotelRef.CitySlug,
		hotelRef.HotelSlug,
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
	).Scan(
		&newRoom.ID,
		&newRoom.Status,
		&newRoom.CreatedAt,
		&newRoom.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, consts.ErrUniqueRoomField
		}
		return nil, err
	}

	return newRoom, nil
}

func (r *Repository) SelectRooms(
	ctx context.Context,
	hotelRef models.HotelRef,
	limit uint64,
	offset uint64,
) (*models.RoomList, error) {
	rows, err := r.db.Query(
		ctx, query.SelectRooms,
		hotelRef.CountryCode,
		hotelRef.CitySlug,
		hotelRef.HotelSlug,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	values := make([]models.RoomShort, limit)
	var room models.RoomShort
	var totalCount, idx uint64
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
			&room.Images,
			&totalCount,
		)
		if err != nil {
			return nil, err
		}

		values[idx] = room
		idx++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	values = values[:idx]

	roomList := &models.RoomList{
		Rooms:      make([]*models.RoomShort, len(values)),
		TotalCount: totalCount,
	}
	for i := range values {
		roomList.Rooms[i] = &values[i]
	}

	return roomList, nil
}

func (r *Repository) SelectRoomByID(ctx context.Context, roomID uuid.UUID) (*models.Room, error) {
	room := &models.Room{ID: roomID}
	err := r.db.QueryRow(ctx, query.SelectRoomByID, roomID).Scan(
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
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, consts.ErrRoomNotFound
		}
		return nil, err
	}

	return room, nil
}

func (r *Repository) UpdateRoomByID(ctx context.Context, roomID uuid.UUID, room *models.UpdateRoom) error {
	row, err := r.db.Exec(
		ctx, query.UpdateRoomByID,
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
			return consts.ErrUniqueRoomField
		}
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.ErrRoomNotFound
	}

	return nil
}

func (r *Repository) UpdateRoomStatusByID(ctx context.Context, roomID uuid.UUID, room models.UpdateRoomStatus) error {
	row, err := r.db.Exec(ctx, query.UpdateRoomStatusByID, roomID, room.Status)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.ErrRoomNotFound
	}

	return nil
}

func (r *Repository) DeleteRoomByID(ctx context.Context, roomID uuid.UUID) error {
	row, err := r.db.Exec(ctx, query.DeleteRoomByID, roomID)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.ErrRoomNotFound
	}

	return nil
}
