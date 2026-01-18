package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateBookingRoom struct {
	PricePerNight decimal.Decimal
	BookingID     uuid.UUID
	RoomID        uuid.UUID
	Adults        uint32
	Children      uint32
}

type BookingRoomGuestCounts struct {
	Adults   uint32
	Children uint32
}

type BookingRoom struct {
	PricePerNight decimal.Decimal
	BookingID     uuid.UUID
	ID            uuid.UUID
	RoomID        uuid.UUID
	Adults        uint32
	Children      uint32
}

type BookingRoomWithLock struct {
	CreatedAt     time.Time
	PricePerNight decimal.Decimal
	RoomLock      RoomLockShort
	ID            uuid.UUID
	RoomID        uuid.UUID
	Adults        uint32
	Children      uint32
}
