package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"booking/internal/repository/models"
	"booking/internal/repository/postgres/query"
	"booking/pkg/utils/consts"
)

func (r *Repository) CreateBookingRooms(
	ctx context.Context,
	tx pgx.Tx,
	rooms []models.CreateBookingRoom,
) ([]models.BookingRoomInfo, error) {

	if len(rooms) == 0 {
		return nil, nil
	}

	db := r.executor(tx)

	bookingID := rooms[0].BookingID

	roomIDs := make([]uuid.UUID, 0, len(rooms))
	adults := make([]int32, 0, len(rooms))
	children := make([]int32, 0, len(rooms))
	prices := make([]string, 0, len(rooms))

	for _, r := range rooms {
		if r.BookingID != bookingID {
			return nil, errors.New("all rooms must have same booking_id")
		}

		roomIDs = append(roomIDs, r.RoomID)
		adults = append(adults, int32(r.Adults))
		children = append(children, int32(r.Children))
		prices = append(prices, r.PricePerNight.StringFixed(2))
	}

	rows, err := db.Query(ctx, query.CreateBookingRooms, bookingID, roomIDs, adults, children, prices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]models.BookingRoomInfo, 0, len(rooms))
	for rows.Next() {
		var br models.BookingRoomInfo

		var a, c int32

		if err := rows.Scan(
			&br.ID,
			&br.BookingID,
			&br.RoomID,
			&a,
			&c,
			&br.PricePerNight,
			&br.CreatedAt,
		); err != nil {
			return nil, err
		}

		br.Adults = uint8(a)
		br.Children = uint8(c)

		out = append(out, br)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func (r *Repository) GetBookingRoomsInfoByBookingIDs(
	ctx context.Context,
	tx pgx.Tx,
	bookingIDs []uuid.UUID,
) ([]models.BookingRoomInfo, error) {
	db := r.executor(tx)

	var bookingRoomList []models.BookingRoomInfo
	rows, err := db.Query(ctx, query.GetBookingRoomsInfoByBookingID, bookingIDs)
	if err != nil {
		return []models.BookingRoomInfo{}, err
	}

	var bRoom models.BookingRoomInfo
	for rows.Next() {
		err = rows.Scan(
			&bRoom.ID,
			&bRoom.BookingID,
			&bRoom.RoomID,
			&bRoom.Adults,
			&bRoom.Children,
			&bRoom.PricePerNight,
			&bRoom.CreatedAt,
		)
		if err != nil {
			return []models.BookingRoomInfo{}, err
		}

		bookingRoomList = append(bookingRoomList, bRoom)
	}

	return bookingRoomList, nil
}

func (r *Repository) GetBookingRoomsFullInfoByBookingIDs(
	ctx context.Context,
	tx pgx.Tx,
	bookingID uuid.UUID,
) ([]models.BookingRoomFullInfo, error) {
	db := r.executor(tx)

	var bookingRoomList []models.BookingRoomFullInfo
	rows, err := db.Query(ctx, query.GetBookingRoomsFullInfoByBookingID, bookingID)
	if err != nil {
		return []models.BookingRoomFullInfo{}, err
	}

	var bRoom models.BookingRoomFullInfo
	var stayRange *pgtype.Range[pgtype.Date]
	for rows.Next() {
		err = rows.Scan(
			&bRoom.ID,
			&bRoom.BookingID,
			&bRoom.RoomID,
			&bRoom.Adults,
			&bRoom.Children,
			&bRoom.PricePerNight,
			&bRoom.CreatedAt,
			&bRoom.RoomLock.ID,
			&stayRange,
			&bRoom.RoomLock.ISActive,
			&bRoom.RoomLock.ExpiresAt,
			&bRoom.RoomLock.CreatedAt,
		)
		if err != nil {
			return []models.BookingRoomFullInfo{}, err
		}

		bRoom.RoomLock.StayRange.Start = stayRange.Lower.Time
		bRoom.RoomLock.StayRange.End = stayRange.Upper.Time

		bookingRoomList = append(bookingRoomList, bRoom)
	}

	return bookingRoomList, nil
}

func (r *Repository) GetBookingRoomByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) (models.BookingRoomInfo, error) {
	db := r.executor(tx)

	var bRoom models.BookingRoomInfo
	scanArgs := []any{
		bRoom.ID,
		bRoom.BookingID,
		bRoom.RoomID,
		bRoom.Adults,
		bRoom.Children,
		bRoom.PricePerNight,
		bRoom.CreatedAt,
	}

	if err := db.QueryRow(ctx, query.GetBookingRoomByID, id).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.BookingRoomInfo{}, consts.BookingRoomNotFound
		}
		return models.BookingRoomInfo{}, err
	}

	return bRoom, nil
}

func (r *Repository) UpdateBookingRoomGuestCountsByID(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
	bRoom models.BookingRoomGuestCounts,
) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.UpdateBookingRoomGuestCountsByID, id, bRoom.Adults, bRoom.Children)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingRoomNotFound
	}

	return nil
}

func (r *Repository) DeleteBookingRoomByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.DeleteBookingRoomByID, id)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingRoomNotFound
	}

	return nil
}
