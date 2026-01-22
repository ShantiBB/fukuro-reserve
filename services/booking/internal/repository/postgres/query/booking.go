package query

const (
	CreateBooking = `
		INSERT INTO booking (
			user_id,
			hotel_id,
			check_in,
			check_out,
			guest_name,
			guest_email,
			guest_phone,
			currency,
			expected_total_amount,
		    final_total_amount
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, status, created_at, updated_at;`

	GetBookingsByHotelInfo = `
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
				expected_total_amount,
				final_total_amount
			FROM booking
			WHERE ($1::bigint IS NULL OR user_id = $1)
			  AND ($2::uuid IS NULL OR hotel_id = $2)
			  AND ($3::booking_status IS NULL OR status = $3)
			ORDER BY created_at DESC
			LIMIT $4 OFFSET $5;`

	GetBookingByID = `
		SELECT
		    id,
			user_id,
			hotel_id::uuid,
			check_in,
			check_out,
			status,
			guest_name,
			guest_email,
			guest_phone,
			currency,
			expected_total_amount,
			final_total_amount,
			created_at,
			updated_at
		FROM booking
		WHERE id = $1;`

	UpdateBookingGuestInfoByID = `
		UPDATE booking
		SET
			guest_name = $2,
			guest_email = $3,
			guest_phone = $4
		WHERE id = $1`

	UpdateBookingStatusByID = `
		UPDATE booking
		SET status = $2
		WHERE id = $1
		RETURNING check_out`

	DeleteBookingByID = `
		DELETE FROM booking
		WHERE id = $1;`

	GetBookingCountRows = `
		SELECT COUNT(*)
		FROM booking
		WHERE ($1::bigint IS NULL OR user_id = $1)
		  AND ($2::uuid IS NULL OR hotel_id = $2)
		  AND ($3::booking_status IS NULL OR status = $3);`
)
