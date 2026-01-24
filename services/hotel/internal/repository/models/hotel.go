package models

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	Latitude  float32
	Longitude float32
}

type CreateHotel struct {
	Description *string
	CountryCode string
	CitySlug    string
	Title       string
	Slug        string
	Address     string
	OwnerID     int64
	Location    Location
}

type UpdateHotel struct {
	Description *string
	Address     string
	Location    Location
}

type UpdateHotelTitle struct {
	Title string
	Slug  string
}

type HotelShort struct {
	Rating   *float32
	Title    string
	Slug     string
	Address  string
	OwnerID  int64
	Location Location
	ID       uuid.UUID
}

type Hotel struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description *string
	Rating      *float32
	Title       string
	Address     string
	Slug        string
	OwnerID     int64
	Location    Location
	ID          uuid.UUID
}

type HotelList struct {
	Hotels     []*HotelShort
	TotalCount uint64
}

func (h *CreateHotel) ToRead() *Hotel {
	return &Hotel{
		Title:       h.Title,
		Slug:        h.Slug,
		OwnerID:     h.OwnerID,
		Description: h.Description,
		Address:     h.Address,
		Location:    h.Location,
	}
}

type HotelRef struct {
	CountryCode string
	CitySlug    string
	HotelSlug   string
}
