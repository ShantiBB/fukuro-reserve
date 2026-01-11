package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateRoomLock struct {
	RoomID    uuid.UUID
	BookingID uuid.UUID
	StayRange time.Time
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
	StayRange time.Time
	ExpiresAt time.Time
	ISActive  bool
	CreatedAt time.Time
}

type RoomLockList struct {
	RoomsLock []RoomLock
}

func (roomLock *CreateRoomLock) ToRead() RoomLock {
	return RoomLock{
		RoomID:    roomLock.RoomID,
		BookingID: roomLock.BookingID,
		StayRange: roomLock.StayRange,
		ExpiresAt: roomLock.ExpiresAt,
	}
}
