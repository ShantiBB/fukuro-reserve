package response

import (
	"time"

	"github.com/google/uuid"

	"hotel/internal/http/utils/pagination"
)

type Location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type HotelCreate struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description *string   `json:"description"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Address     string    `json:"address"`
	OwnerID     int64     `json:"owner_id"`
	Location    Location  `json:"location"`
	ID          uuid.UUID `json:"id"`
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
	Rating   *float32  `json:"rating"`
	Title    string    `json:"title"`
	Slug     string    `json:"slug"`
	Address  string    `json:"address"`
	OwnerID  int64     `json:"owner_id"`
	Location Location  `json:"location"`
	ID       uuid.UUID `json:"id"`
}

type Hotel struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description *string   `json:"description"`
	Rating      *float32  `json:"rating"`
	Title       string    `json:"title"`
	Address     string    `json:"address"`
	OwnerID     int64     `json:"owner_id"`
	Location    Location  `json:"location"`
	ID          uuid.UUID `json:"id"`
}

type HotelList struct {
	Links            pagination.Links `json:"links"`
	Hotels           []HotelShort     `json:"hotels"`
	CurrentPage      uint64           `json:"current_page"`
	Limit            uint64           `json:"limit"`
	TotalPageCount   uint64           `json:"total_page_count"`
	TotalHotelsCount uint64           `json:"total_rooms_count"`
}
