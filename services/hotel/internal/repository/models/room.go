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
	Price       decimal.Decimal
	Capacity    int
	AreaSqm     *float64
	Floor       *int
	Description *string
	Amenities   []string
	Images      []string
}

type RoomUpdate struct {
	RoomNumber  string
	Type        RoomType
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

type RoomShort struct {
	ID          uuid.UUID
	HotelID     uuid.UUID
	RoomNumber  string
	Type        RoomType
	Status      RoomStatus
	Price       decimal.Decimal
	Capacity    int
	AreaSqm     *float64
	Description *string
	Amenities   []string
	Images      []string
}
type RoomList struct {
	Rooms      []RoomShort
	TotalCount uint64
}

func (r *RoomCreate) ToRead() Room {
	return Room{
		HotelID:     r.HotelID,
		RoomNumber:  r.RoomNumber,
		Type:        r.Type,
		Price:       r.Price,
		Capacity:    r.Capacity,
		AreaSqm:     r.AreaSqm,
		Floor:       r.Floor,
		Description: r.Description,
		Amenities:   r.Amenities,
		Images:      r.Images,
	}
}
