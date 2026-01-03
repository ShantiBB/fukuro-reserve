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
	RETURNING id, created_at, updated_at`

	RoomGetByID = `
	SELECT id,
	       hotel_id,
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
	WHERE id = $1`

	RoomGetAll = `
	SELECT id,
	       hotel_id,
	       description,
		   room_number,
		   type,
		   price,
		   capacity,
		   area_sqm,
		   floor,
		   amenities,
		   images
	FROM room
	ORDER BY name
	LIMIT $1 OFFSET $2;`

	RoomUpdateByID = `
	UPDATE room 
	SET description = $5,
	    price = $1,
		capacity = $2,
		area_sqm = $3,
		floor = $4,
		amenities = $6,
		images = $7
	WHERE id = $6;`

	RoomDeleteByID = `
	DELETE FROM room 
	WHERE id = $1;`

	RoomGetCountRows = `SELECT COUNT(*) FROM room;`
)
