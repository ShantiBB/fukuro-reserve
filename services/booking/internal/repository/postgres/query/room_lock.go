package query

const (
	CreateRoomLocks = `
		WITH input AS (
		  SELECT *
		  FROM unnest(
			$1::uuid[],
			$2::uuid[],
			$3::date[],
			$4::date[],
			$5::timestamptz[]
		  ) AS t(room_id, booking_id, start_date, end_date, expires_at)
		)
		INSERT INTO room_lock (room_id, booking_id, stay_range, expires_at)
		SELECT
		  room_id,
		  booking_id,
		  daterange(start_date, end_date, '[)'),
		  expires_at
		FROM input
		RETURNING id, room_id, booking_id, is_active, created_at;`

	UpdateRoomLockActivityByID = `
		UPDATE room_lock
		SET
		  is_active = $2,
		  expires_at = COALESCE($3, expires_at)
		WHERE booking_id = $1;`

	DeleteRoomLockByID = `
		DELETE FROM room_lock
		WHERE id = $1;`
)
