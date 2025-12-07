package response

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type HotelCreate struct {
	Name        string   `json:"name"`
	OwnerID     int64    `json:"owner_id"`
	Description *string  `json:"description"`
	Address     string   `json:"address"`
	Location    Location `json:"location"`
}

type HotelUpdate struct {
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Address     string   `json:"address"`
	Location    Location `json:"location"`
}

type HotelShort struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	OwnerID  int64     `json:"owner_id"`
	Address  string    `json:"address"`
	Rating   *float32  `json:"rating"`
	Location Location  `json:"location"`
}

type Hotel struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	OwnerID     int64     `json:"owner_id"`
	Description *string   `json:"description"`
	Address     string    `json:"address"`
	Rating      *float32  `json:"rating"`
	Location    Location  `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
