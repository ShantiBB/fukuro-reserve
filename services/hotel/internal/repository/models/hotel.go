package models

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	Latitude  float64
	Longitude float64
}

type HotelCreate struct {
	CountryCode string
	CitySlug    string
	Title       string
	Slug        string
	OwnerID     int64
	Description *string
	Address     string
	Location    Location
}

type HotelUpdate struct {
	CountryCode string
	CitySlug    string
	Title       string
	Slug        string
	Description *string
	Address     string
	Location    Location
}

type HotelShort struct {
	ID          uuid.UUID
	CountryCode string
	CitySlug    string
	Title       string
	Slug        string
	OwnerID     int64
	Address     string
	Rating      *float32
	Location    Location
}

type Hotel struct {
	ID          uuid.UUID
	CountryCode string
	CitySlug    string
	Title       string
	Slug        string
	OwnerID     int64
	Description *string
	Address     string
	Rating      *float32
	Location    Location
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type HotelList struct {
	Hotels     []HotelShort
	TotalCount uint64
}

type HotelFilter struct {
	CountryCode string
	CitySlug    string
	Slug        string
	SortField   string
	Limit       uint64
	Offset      uint64
}

func (h *HotelCreate) ToRead() Hotel {
	return Hotel{
		CountryCode: h.CountryCode,
		CitySlug:    h.CitySlug,
		Title:       h.Title,
		Slug:        h.Slug,
		OwnerID:     h.OwnerID,
		Description: h.Description,
		Address:     h.Address,
		Location:    h.Location,
	}
}
