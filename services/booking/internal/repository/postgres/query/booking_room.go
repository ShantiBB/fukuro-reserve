package query

const (
	CreateBookingRoom = `
		INSERT INTO booking_room (booking_id, room_id, adults, children, price_per_night)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at;`

	GetBookingRoomsByBookingID = `
		SELECT
		  id,
		  booking_id,
		  room_id,
		  adults,
		  children,
		  price_per_night,
		  created_at
		FROM booking_room
		WHERE booking_id = $1
		ORDER BY created_at;`

	GetBookingRoomByID = `
		SELECT
		  id,
		  booking_id,
		  room_id,
		  adults,
		  children,
		  price_per_night,
		  created_at
		FROM booking_room
		WHERE id = $1;`

	UpdateBookingRoomGuestCountsByID = `
		UPDATE booking_room
		SET
		  adults = $2,
		  children = $3
		WHERE id = $1;`

	DeleteBookingRoomByID = `
		DELETE FROM booking_room
		WHERE id = $1;`
)
