package query

const (
	InsertRoomQuery = `
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

	SelectRooms = `
		SELECT r.id,
			   r.title,
			   r.room_number,
			   r.type,
			   r.status,
			   r.price,
			   r.capacity,
			   r.area_sqm,
			   r.amenities,
			   r.images,
			   COUNT(*) OVER() as total_count
		FROM room r
		JOIN hotel h ON h.id = r.hotel_id
		WHERE h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3
		ORDER BY r.room_number
		LIMIT $4 OFFSET $5;`

	SelectRoomByID = `
		SELECT title,
			   description,
			   room_number,
			   type,
			   status,
			   price,
			   capacity,
			   area_sqm,
			   floor,
			   amenities,
			   images,
			   created_at,
			   updated_at
		FROM room
		WHERE id = $1;`

	UpdateRoomByID = `
		UPDATE room
		SET title       = $2,
		    description = $3,
		    room_number = $4,
		    type        = $5,
		    price       = $6,
		    capacity    = $7,
		    area_sqm    = $8,
		    floor       = $9,
		    amenities   = $10,
		    images      = $11
		WHERE id = $1;`

	UpdateRoomStatusByID = `
		UPDATE room
		SET status = $2
		WHERE id = $1;`

	DeleteRoomByID = `
		DELETE FROM room
		WHERE id = $1;`
)
