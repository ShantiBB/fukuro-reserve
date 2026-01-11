package models

import (
	"time"

	"github.com/google/uuid"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

type CreateRoomLock struct {
	RoomID    uuid.UUID
	BookingID uuid.UUID
	StayRange DateRange
	ExpiresAt time.Time
}

type UpdateRoomLockActivity struct {
	IsActive  bool
	ExpiresAt time.Time
}
type RoomLock struct {
	ID        uuid.UUID
	RoomID    uuid.UUID
	BookingID uuid.UUID
	StayRange DateRange
	ExpiresAt time.Time
	ISActive  bool
	CreatedAt time.Time
}

func (roomLock *CreateRoomLock) ToRead() RoomLock {
	return RoomLock{
		RoomID:    roomLock.RoomID,
		BookingID: roomLock.BookingID,
		StayRange: roomLock.StayRange,
		ExpiresAt: roomLock.ExpiresAt,
	}
}
