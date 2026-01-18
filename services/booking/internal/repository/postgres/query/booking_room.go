package query

const (
	CreateBookingRooms = `
		WITH input AS (
		  SELECT
			$1::uuid AS booking_id,
			unnest($2::uuid[])    AS room_id,
			unnest($3::int[])     AS adults,
			unnest($4::int[])     AS children,
			unnest($5::numeric[]) AS price_per_night
		)
		INSERT INTO booking_room (booking_id, room_id, adults, children, price_per_night)
		SELECT booking_id, room_id, adults, children, price_per_night
		FROM input
		RETURNING id, created_at;`

	GetBookingRoomsByBookingID = `
		SELECT
			id::uuid,
			booking_id::uuid,
			room_id::uuid,
			adults,
			children,
			price_per_night
		FROM booking_room
		WHERE booking_id = ANY($1)
		ORDER BY created_at;`

	GetBookingRoomsWithLockByBookingID = `
		SELECT
			br.id::uuid as booking_room_id,
			br.room_id::uuid,
			br.adults,
			br.children,
			br.price_per_night,
			br.created_at,
			rl.id::uuid,
			rl.is_active,
			rl.expires_at,
			rl.created_at
		FROM booking_room br
		LEFT JOIN room_lock rl ON rl.booking_id = br.booking_id AND rl.room_id = br.room_id
		WHERE br.booking_id = ANY($1)
		ORDER BY br.created_at;`

	GetBookingRoomsWithLockDetailByBookingID = `
		SELECT
			br.id::uuid as booking_room_id,
			br.room_id::uuid,
			br.adults,
			br.children,
			br.price_per_night,
			br.created_at,
			rl.id::uuid,
			rl.stay_range,
			rl.is_active,
			rl.expires_at,
			rl.created_at
		FROM booking_room br
		LEFT JOIN room_lock rl ON rl.booking_id = br.booking_id AND rl.room_id = br.room_id
		WHERE br.booking_id = ANY($1)
		ORDER BY br.created_at;`

	GetBookingRoomWithLockByID = `
		SELECT
			br.id::uuid as booking_room_id,
			br.room_id::uuid,
			br.adults,
			br.children,
			br.price_per_night,
			br.created_at,
			rl.id::uuid,
			rl.is_active,
			rl.expires_at,
			rl.created_at
		FROM booking_room br
		LEFT JOIN room_lock rl ON rl.booking_id = br.booking_id AND rl.room_id = br.room_id
		WHERE br.id = $1;`

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
