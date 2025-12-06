package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type RoomCreate struct {
	HotelID     uuid.UUID
	RoomNumber  string
	Type        RoomType
	Status      RoomStatus
	Price       decimal.Decimal
	Capacity    int
	AreaSqm     *float64
	Floor       *int
	Description *string
	Amenities   []string
	Images      []string
}

type Room struct {
	ID          uuid.UUID
	HotelID     uuid.UUID
	RoomNumber  string
	Type        RoomType
	Status      RoomStatus
	Price       decimal.Decimal
	Capacity    int
	AreaSqm     *float64
	Floor       *int
	Description *string
	Amenities   []string
	Images      []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (r *RoomCreate) ToRead() Room {
	return Room{
		HotelID:     r.HotelID,
		RoomNumber:  r.RoomNumber,
		Type:        r.Type,
		Status:      r.Status,
		Price:       r.Price,
		Capacity:    r.Capacity,
		AreaSqm:     r.AreaSqm,
		Floor:       r.Floor,
		Description: r.Description,
		Amenities:   r.Amenities,
		Images:      r.Images,
	}
}
