package query

const (
	RoomCreateQuery = `
		INSERT INTO room (
			hotel_id,
			title,
			description,
			room_number,
			type,
			price,
			capacity,
			area_sqm,
			floor,
			amenities,
			images
		)
		SELECT h.id, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
		FROM hotel h
		WHERE h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3
		RETURNING id, status, created_at, updated_at;`

	RoomGetAll = `
		SELECT r.id,
			   r.title,
			   r.room_number,
			   r.type,
			   r.status,
			   r.price,
			   r.capacity,
			   r.area_sqm,
			   r.amenities,
			   r.images
		FROM room r
		JOIN hotel h ON h.id = r.hotel_id
		WHERE h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3
		ORDER BY r.room_number
		LIMIT $4 OFFSET $5;`

	RoomGetByID = `
		SELECT r.title,
			   r.description,
			   r.room_number,
			   r.type,
			   r.status,
			   r.price,
			   r.capacity,
			   r.area_sqm,
			   r.floor,
			   r.amenities,
			   r.images,
			   r.created_at,
			   r.updated_at
		FROM room r
		JOIN hotel h ON h.id = r.hotel_id
		WHERE h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3 AND r.id = $4;`

	RoomUpdateByID = `
		UPDATE room r
		SET title       = $5,
		    description = $6,
		    room_number = $7,
		    type        = $8,
		    price       = $9,
		    capacity    = $10,
		    area_sqm    = $11,
		    floor       = $12,
		    amenities   = $13,
		    images      = $14
		FROM hotel h
		WHERE h.id = r.hotel_id AND h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3 AND r.id = $4;`

	RoomDeleteByID = `
		DELETE FROM room r
		USING hotel h
		WHERE h.id = r.hotel_id AND h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3 AND r.id = $4;`

	RoomGetCountRows = `
		SELECT COUNT(*)
		FROM room r
		JOIN hotel h ON h.id = r.hotel_id
		WHERE h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3;`
)
