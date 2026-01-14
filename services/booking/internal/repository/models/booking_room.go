package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateBookingRoom struct {
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

type BookingRoomInfo struct {
	ID            uuid.UUID
	BookingID     uuid.UUID
	RoomID        uuid.UUID
	Adults        uint8
	Children      uint8
	PricePerNight decimal.Decimal
	CreatedAt     time.Time
}

type BookingRoomFullInfo struct {
	ID            uuid.UUID
	BookingID     uuid.UUID
	RoomID        uuid.UUID
	RoomLock      RoomLockShort
	Adults        uint8
	Children      uint8
	PricePerNight decimal.Decimal
	CreatedAt     time.Time
}

func (b *CreateBookingRoom) ToRead() BookingRoomInfo {
	return BookingRoomInfo{
		BookingID:     b.BookingID,
		RoomID:        b.RoomID,
		Adults:        b.Adults,
		Children:      b.Children,
		PricePerNight: b.PricePerNight,
	}
}
