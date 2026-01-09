package response

import (
	"time"

	"github.com/google/uuid"

	"hotel/internal/http/utils/pagination"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type HotelCreate struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	OwnerID     int64     `json:"owner_id"`
	Description *string   `json:"description"`
	Address     string    `json:"address"`
	Location    Location  `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type HotelUpdate struct {
	Description *string  `json:"description"`
	Address     string   `json:"address"`
	Location    Location `json:"location"`
}

type HotelTitleUpdate struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type HotelShort struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Slug     string    `json:"slug"`
	OwnerID  int64     `json:"owner_id"`
	Address  string    `json:"address"`
	Rating   *float32  `json:"rating"`
	Location Location  `json:"location"`
}

type Hotel struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	OwnerID     int64     `json:"owner_id"`
	Description *string   `json:"description"`
	Address     string    `json:"address"`
	Rating      *float32  `json:"rating"`
	Location    Location  `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type HotelList struct {
	Hotels           []HotelShort     `json:"hotels"`
	CurrentPage      uint64           `json:"current_page"`
	Limit            uint64           `json:"limit"`
	Links            pagination.Links `json:"links"`
	TotalPageCount   uint64           `json:"total_page_count"`
	TotalHotelsCount uint64           `json:"total_rooms_count"`
}
