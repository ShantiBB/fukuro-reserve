package query

const (
	CreateBooking = `
		INSERT INTO booking (
			country_code,
			city_slug,
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
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
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
			  AND ($4::text = '' OR country_code = $4)
			  AND ($5::text = '' OR city_slug = $5)
			  AND (
			    $6::uuid IS NULL
			    OR EXISTS (
			      SELECT 1
			      FROM booking_room br
			      WHERE br.booking_id = booking.id AND br.room_id = $6
			    )
			  )
			ORDER BY created_at DESC
			LIMIT $7 OFFSET $8;`

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
		WHERE id = $1 AND country_code = $2 AND city_slug = $3;`

	UpdateBookingGuestInfoByID = `
		UPDATE booking
		SET
			guest_name = COALESCE($5, guest_name),
			guest_email = COALESCE($6, guest_email),
			guest_phone = COALESCE($7, guest_phone)
		WHERE id = $1
		  AND country_code = $2
		  AND city_slug = $3
		  AND hotel_id = $4;`

	UpdateBookingStatusByID = `
		UPDATE booking
		SET status = $2
		WHERE id = $1 AND country_code = $3 AND city_slug = $4
		RETURNING check_out`

	DeleteBookingByID = `
		DELETE FROM booking
		WHERE id = $1 AND country_code = $2 AND city_slug = $3;`

	GetBookingCountRows = `
		SELECT COUNT(*)
		FROM booking
		WHERE ($1::bigint IS NULL OR user_id = $1)
		  AND ($2::uuid IS NULL OR hotel_id = $2)
		  AND ($3::booking_status IS NULL OR status = $3)
		  AND ($4::text = '' OR country_code = $4)
		  AND ($5::text = '' OR city_slug = $5)
		  AND (
		    $6::uuid IS NULL
		    OR EXISTS (
		      SELECT 1
		      FROM booking_room br
		      WHERE br.booking_id = booking.id AND br.room_id = $6
		    )
		  );`
)
