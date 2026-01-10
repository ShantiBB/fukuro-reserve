package query

const (
	BookingCreate = `
		INSERT INTO booking (
			user_id,
			hotel_id,
			check_in,
			check_out,
			guest_name,
			guest_email,
			guest_phone,
			currency,
			total_amount
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at;`

	BookingGetAll = `
		SELECT
			id,
			user_id,
			hotel_id,
			check_in,
			check_out,
			status,
			guest_name,
			guest_email,
			guest_phone,
			currency,
			total_amount,
			created_at,
			updated_at
		FROM booking
		WHERE ($1::bigint IS NULL OR user_id = $1)
		  AND ($2::uuid IS NULL OR hotel_id = $2)
		  AND ($3::booking_status IS NULL OR status = $3)
		ORDER BY created_at DESC
		LIMIT $4 OFFSET $5;`

	BookingGetByID = `
		SELECT
			id,
			user_id,
			hotel_id,
			check_in,
			check_out,
			status,
			guest_name,
			guest_email,
			guest_phone,
			currency,
			total_amount,
			created_at,
			updated_at
		FROM booking
		WHERE id = $1;`

	BookingUpdate = `
		UPDATE booking
		SET
			guest_name = $2,
			guest_email = $3,
			guest_phone = $4,
			check_in = $5,
			check_out = $6,
			total_amount = $7
		WHERE id = $1`

	BookingUpdateStatus = `
		UPDATE booking
		SET status = $2
		WHERE id = $1`

	BookingDelete = `
		DELETE FROM booking
		WHERE id = $1;`
)
