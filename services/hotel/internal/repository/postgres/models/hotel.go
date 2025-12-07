package models

import (
	"time"

	"github.com/google/uuid"
)

type HotelCreate struct {
	Name        string
	OwnerID     int64
	Description *string
	Address     string
	Location    Location
}

type HotelUpdate struct {
	Name        string
	Description *string
	Address     string
	Location    Location
}

type Hotel struct {
	ID          uuid.UUID
	Name        string
	OwnerID     int64
	Description *string
	Address     string
	Rating      *float32
	Location    Location
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type HotelShort struct {
	ID       uuid.UUID
	Name     string
	OwnerID  int64
	Address  string
	Rating   *float32
	Location Location
}

type Location struct {
	Latitude  float64
	Longitude float64
}

func (h *HotelCreate) ToRead() Hotel {
	return Hotel{
		Name:        h.Name,
		OwnerID:     h.OwnerID,
		Description: h.Description,
		Address:     h.Address,
		Location:    h.Location,
	}
}
