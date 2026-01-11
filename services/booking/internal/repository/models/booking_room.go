package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type BookingRoomCreate struct {
	BookingID     uuid.UUID
	RoomID        uuid.UUID
	Adults        uint8
	Children      uint8
	PricePerNight decimal.Decimal
}

type BookingRoomGuestCounts struct {
	Adults   uint8
	Children uint8
}

type BookingRoom struct {
	ID            uuid.UUID
	BookingID     uuid.UUID
	RoomID        uuid.UUID
	Adults        uint8
	Children      uint8
	PricePerNight decimal.Decimal
	CreatedAt     time.Time
}

type BookingRoomList struct {
	BookingRooms []BookingRoom
}

func (b *BookingRoomCreate) ToRead() BookingRoom {
	return BookingRoom{
		BookingID:     b.BookingID,
		RoomID:        b.RoomID,
		Adults:        b.Adults,
		Children:      b.Children,
		PricePerNight: b.PricePerNight,
	}
}
