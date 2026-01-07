package request

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type HotelCreate struct {
	Title       string   `json:"title" validate:"required"`
	OwnerID     int64    `json:"owner_id" validate:"required"`
	Description *string  `json:"description"`
	Address     string   `json:"address"`
	Location    Location `json:"location"`
}

type HotelUpdate struct {
	Title       string   `json:"title" validate:"required"`
	Description *string  `json:"description"`
	Address     string   `json:"address"`
	Location    Location `json:"location"`
}
