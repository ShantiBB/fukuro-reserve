package postgres

const (
	hotelCreateQuery = `
	INSERT INTO hotel (name, owner_id, description, address, location)
	VALUES ($1, $2, $3, $4, ST_SetSRID(ST_MakePoint($5, $6), 4326))
	RETURNING id, created_at, updated_at`

	hotelGetByID = `
	SELECT id,
	       name, 
	       owner_id, 
	       description, 
	       address,
		   ST_X(location::geometry) AS longitude,
       	   ST_Y(location::geometry) AS latitude,
	       rating, 
	       created_at, 
	       updated_at
	FROM hotel 
	WHERE id = $1`

	hotelGetByName = `
	SELECT id,
	       name,
	       owner_id, 
	       description, 
	       address,
		   st_x(location::geometry) AS longitude,
       	   st_y(location::geometry) AS latitude,
	       rating, 
	       created_at, 
	       updated_at
	FROM hotel 
	WHERE name = $1`

	hotelGetAll = `
	SELECT id, name, owner_id, address, rating,
	       st_x(location::geometry) AS longitude,
	       st_y(location::geometry) AS latitude
	FROM hotel
	ORDER BY name
	LIMIT $1 OFFSET $2;`

	hotelUpdateByID = `
	UPDATE hotel 
	SET name = $1, description = $2, address = $3, 
	    location = ST_SetSRID(ST_MakePoint($4, $5), 4326)
	WHERE id = $6;`

	hotelDeleteByID = `
	DELETE FROM hotel 
	WHERE id = $1;`
)
