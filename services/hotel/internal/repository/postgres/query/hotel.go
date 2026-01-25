package query

const (
	CreateHotelQuery = `
		INSERT INTO hotel (country_code, 
						   city_slug, 
						   title, 
						   slug, 
						   owner_id, 
						   description, 
						   address, 
						   location)
		VALUES ($1, $2, $3, $4, $5, $6, $7, ST_SetSRID(ST_MakePoint($8, $9), 4326)::geography)
		RETURNING id, created_at, updated_at`

	GetHotelBySlug = `
		SELECT id,
			   title, 
			   owner_id, 
			   description, 
			   address,
			   longitude,
			   latitude,
			   rating, 
			   created_at, 
			   updated_at
		FROM hotel 
		WHERE country_code = $1 AND city_slug = $2 AND slug = $3`

	GetHotels = `
		SELECT id,
			   title,
			   slug,
			   owner_id, 
			   address, 
			   rating,
			   longitude,
			   latitude,
			   COUNT(*) OVER() as total_count
		FROM hotel
		WHERE country_code = $1 AND city_slug = $2
		ORDER BY
		    CASE WHEN $3 = 'title' THEN title END,
			CASE WHEN $3 = 'rating' THEN rating END DESC
		LIMIT $4 OFFSET $5;`

	UpdateHotelBySlug = `
		UPDATE hotel
		SET
		  description = $1,
		  address = $2,
		  location = ST_SetSRID(ST_MakePoint($3, $4), 4326)::geography,
		  updated_at = now()
		WHERE country_code = $5 AND city_slug = $6 AND slug = $7
		RETURNING id, slug;`

	UpdateHotelTitleBySlug = `
		UPDATE hotel
		SET
		  title = $1,
		  slug = $2,
		  updated_at = now()
		WHERE country_code = $3 AND city_slug = $4 AND slug = $5
		RETURNING id, slug;`

	DeleteHotelBySlug = `
		DELETE FROM hotel 
		WHERE country_code = $1 AND city_slug = $2 AND slug = $3;`

	UpdateHotelRating = `
		UPDATE hotel 
		SET rating = $1,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $2`
)
