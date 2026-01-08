package request

type Location struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type HotelCreate struct {
	Title       string   `json:"title" validate:"required"`
	OwnerID     int64    `json:"owner_id" validate:"required"`
	Description *string  `json:"description"`
	Address     string   `json:"address" validate:"required"`
	Location    Location `json:"location" validate:"required"`
}

type HotelUpdate struct {
	Description *string  `json:"description"`
	Address     string   `json:"address" validate:"required"`
	Location    Location `json:"location" validate:"required"`
}

type HotelTitleUpdate struct {
	Title string `json:"title" validate:"required"`
}
