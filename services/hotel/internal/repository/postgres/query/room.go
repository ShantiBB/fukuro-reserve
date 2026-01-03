package query

const (
	RoomCreateQuery = `
	INSERT INTO room (hotel_id,
	                  description,
	                  room_number,
	                  type,
	                  price,
	                  capacity,
	                  area_sqm,
	                  floor,
	                  amenities,
	                  images)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id, status, created_at, updated_at`

	RoomGetByID = `
	SELECT id,
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
	WHERE hotel_id = $1 AND id = $2`

	RoomGetAll = `
	SELECT id,
	       description,
		   room_number,
		   type,
		   status,
		   price,
		   capacity,
		   area_sqm,
		   amenities,
		   images
	FROM room
	WHERE hotel_id = $1
	ORDER BY room_number
	LIMIT $2 OFFSET $3;`

	RoomUpdateByID = `
	UPDATE room 
	SET description = $1,
	    room_number = $2,
	    type = $3,
	    price = $4,
		capacity = $5,
		area_sqm = $6,
		floor = $7,
		amenities = $8,
		images = $9
	WHERE hotel_id = $10 AND id = $11;`

	RoomDeleteByID = `
	DELETE FROM room 
	WHERE hotel_id = $1 AND id = $2;`

	RoomGetCountRows = `SELECT COUNT(*) FROM room;`
)
