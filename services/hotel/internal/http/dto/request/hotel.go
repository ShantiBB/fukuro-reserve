package request

type Location struct {
	Latitude  *float32 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude *float32 `json:"longitude" validate:"required,gte=-180,lte=180"`
}

type HotelCreate struct {
	Title       *string   `json:"title" validate:"required,min=3,max=100"`
	OwnerID     *int64    `json:"owner_id" validate:"required,gt=0"`
	Description *string   `json:"description" validate:"omitempty,max=2000"`
	Address     *string   `json:"address" validate:"required,min=5,max=500"`
	Location    *Location `json:"location" validate:"required"`
}

type HotelUpdate struct {
	Description *string   `json:"description" validate:"omitempty,max=2000"`
	Address     *string   `json:"address" validate:"required,min=5,max=500"`
	Location    *Location `json:"location" validate:"required"`
}

type HotelTitleUpdate struct {
	Title *string `json:"title" validate:"required,min=3,max=100"`
}

type HotelPathParams struct {
	CountryCode string `json:"country_code" validate:"required,len=2,lowercase,alpha"`
	CitySlug    string `json:"city_slug" validate:"required,min=1,max=100,slug_format"`
	HotelSlug   string `json:"hotel_slug" validate:"omitempty,min=1,max=100,slug_format"`
}
