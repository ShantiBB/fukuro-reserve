package query

const (
	HotelCreateQuery = `
		INSERT INTO hotel (country_code, 
						   city_slug, 
						   title, 
						   slug, 
						   owner_id, 
						   description, 
						   address, 
						   location)
		VALUES ($1, $2, $3, $4, $5, $6, $7, ST_SetSRID(ST_MakePoint($8, $9), 4326))
		RETURNING id, created_at, updated_at`

	HotelGetBySlug = `
		SELECT id,
			   title, 
			   owner_id, 
			   description, 
			   address,
			   ST_X(location::geometry) AS longitude,
			   ST_Y(location::geometry) AS latitude,
			   rating, 
			   created_at, 
			   updated_at
		FROM hotel 
		WHERE country_code = $1 AND city_slug = $2 AND slug = $3`

	HotelGetAll = `
		SELECT id,
			   title,
			   slug,
			   owner_id, 
			   address, 
			   rating,
			   st_x(location::geometry) AS longitude,
			   st_y(location::geometry) AS latitude
		FROM hotel
		WHERE country_code = $1 AND city_slug = $2
		ORDER BY
		    CASE WHEN $3 = 'title' THEN title END,
			CASE WHEN $3 = 'rating' THEN rating END DESC
		LIMIT $4 OFFSET $5;`

	HotelUpdateBySlug = `
		UPDATE hotel
		SET
		  title = $1::text,
		  slug = CASE WHEN title IS DISTINCT FROM $1::text THEN $2 ELSE slug END,
		  description = $3,
		  address = $4,
		  location = ST_SetSRID(ST_MakePoint($5, $6), 4326),
		  updated_at = now()
		WHERE country_code = $7 AND city_slug = $8 AND slug = $9
		RETURNING id, slug;`

	HotelDeleteBySlug = `
		DELETE FROM hotel 
		WHERE country_code = $1 AND city_slug = $2 AND slug = $3;`

	HotelGetCountRows = `SELECT COUNT(*) FROM hotel;`

	HotelUpdateRating = `
		UPDATE hotel 
		SET rating = $1,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $2`
)
