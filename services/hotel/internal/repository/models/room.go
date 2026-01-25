package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateRoom struct {
	Description *string
	Title       string
	RoomNumber  string
	Type        RoomType
	Price       decimal.Decimal
	Amenities   []string
	Images      []string
	Capacity    int
	AreaSqm     float64
	Floor       int
}

type UpdateRoom struct {
	Description string
	Title       string
	RoomNumber  string
	Type        RoomType
	Price       decimal.Decimal
	Amenities   []string
	Images      []string
	Capacity    int
	AreaSqm     float64
	Floor       int
}

type UpdateRoomStatus struct {
	Status RoomStatus
}

type Room struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description *string
	Price       decimal.Decimal
	Type        RoomType
	Status      RoomStatus
	RoomNumber  string
	Title       string
	Amenities   []string
	Images      []string
	Capacity    int
	AreaSqm     float64
	Floor       int
	ID          uuid.UUID
}

type RoomShort struct {
	Title      string
	RoomNumber string
	Type       RoomType
	Status     RoomStatus
	Price      decimal.Decimal
	Amenities  []string
	Images     []string
	Capacity   int
	AreaSqm    float64
	ID         uuid.UUID
}
type RoomList struct {
	Rooms      []*RoomShort
	TotalCount uint64
}

func (r *CreateRoom) ToRead() *Room {
	return &Room{
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
