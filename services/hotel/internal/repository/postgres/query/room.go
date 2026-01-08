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
	WHERE country_code = $1
	  AND city_slug    = $2
	  AND slug         = $3
	RETURNING id, status, created_at, updated_at;`

	RoomGetByID = `
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
	FROM room r
	JOIN hotel h ON h.id = r.hotel_id
	WHERE h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3 AND r.id = $4`

	RoomGetAll = `
	SELECT id,
	       title,
		   room_number,
		   type,
		   status,
		   price,
		   capacity,
		   area_sqm,
		   amenities,
		   images
	FROM room r
	JOIN hotel h ON h.id = r.hotel_id
	WHERE h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3
	ORDER BY room_number
	LIMIT $4 OFFSET $5;`

	RoomUpdateByID = `
	UPDATE room r
	SET title = $5,
	    description = $6,
	    room_number = $7,
	    type = $8,
	    price = $9,
		capacity = $10,
		area_sqm = $11,
		floor = $12,
		amenities = $13,
		images = $14
	JOIN hotel h ON h.id = r.hotel_id
	WHERE h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3 AND r.id = $4;`

	RoomDeleteByID = `
	DELETE FROM room r
	JOIN hotel h ON h.id = r.hotel_id
	WHERE h.country_code = $1 AND h.city_slug = $2 AND h.slug = $3 AND r.id = $4;`

	RoomGetCountRows = `SELECT COUNT(*) FROM room;`
)
