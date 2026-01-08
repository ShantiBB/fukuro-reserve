package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type RoomCreate struct {
	Title       string
	Description *string
	RoomNumber  string
	Type        RoomType
	Price       decimal.Decimal
	Capacity    int
	AreaSqm     float64
	Floor       int
	Amenities   []string
	Images      []string
}

type RoomUpdate struct {
	Title       string
	Description *string
	RoomNumber  string
	Type        RoomType
	Price       decimal.Decimal
	Capacity    int
	AreaSqm     float64
	Floor       int
	Amenities   []string
	Images      []string
}

type Room struct {
	ID          uuid.UUID
	Title       string
	Description *string
	RoomNumber  string
	Type        RoomType
	Status      RoomStatus
	Price       decimal.Decimal
	Capacity    int
	AreaSqm     float64
	Floor       int
	Amenities   []string
	Images      []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RoomShort struct {
	ID         uuid.UUID
	Title      string
	RoomNumber string
	Type       RoomType
	Status     RoomStatus
	Price      decimal.Decimal
	Capacity   int
	AreaSqm    float64
	Amenities  []string
	Images     []string
}
type RoomList struct {
	Rooms      []RoomShort
	TotalCount uint64
}

func (r *RoomCreate) ToRead() Room {
	return Room{
		Title:       r.Title,
		Description: r.Description,
		RoomNumber:  r.RoomNumber,
		Type:        r.Type,
		Price:       r.Price,
		Capacity:    r.Capacity,
		AreaSqm:     r.AreaSqm,
		Floor:       r.Floor,
		Amenities:   r.Amenities,
		Images:      r.Images,
	}
}
