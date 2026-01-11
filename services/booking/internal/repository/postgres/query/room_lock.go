package query

const (
	CreateRoomLock = `
		INSERT INTO room_lock (room_id, booking_id, stay_range, expires_at)
		VALUES ($1, $2, daterange($3::date, $4::date, '[)'), $5)
		RETURNING id, is_active, created_at;`

	GetRoomsLockByBookingID = `
		SELECT
		  id,
		  room_id,
		  booking_id,
		  stay_range,
		  is_active,
		  expires_at,
		  created_at
		FROM room_lock
		WHERE booking_id = $1
		ORDER BY created_at;`

	GetRoomLockByID = `
		SELECT
		  id,
		  room_id,
		  booking_id,
		  stay_range,
		  is_active,
		  expires_at,
		  created_at
		FROM room_lock
		WHERE id = $1;`

	UpdateRoomLockActivityByID = `
		UPDATE room_lock
		SET
		  is_active = $2,
		  expires_at = $3
		WHERE id = $1;`

	DeleteRoomLockByID = `
		DELETE FROM room_lock
		WHERE id = $1;`
)
