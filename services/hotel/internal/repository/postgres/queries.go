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
		   ST_X(location::geometry) AS longitude,
       	   ST_Y(location::geometry) AS latitude,
	       rating, 
	       created_at, 
	       updated_at
	FROM hotel 
	WHERE name = $1`
)
