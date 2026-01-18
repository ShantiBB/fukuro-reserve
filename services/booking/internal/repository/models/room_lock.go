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
	StayRange DateRange
	ExpiresAt time.Time
	RoomID    uuid.UUID
	BookingID uuid.UUID
}

type UpdateRoomLockActivity struct {
	ExpiresAt time.Time
	IsActive  bool
}

type RoomLock struct {
	StayRange DateRange
	ExpiresAt time.Time
	CreatedAt time.Time
	ID        uuid.UUID
	ISActive  bool
}

type RoomLockDetail struct {
	StayRange DateRange
	ExpiresAt time.Time
	CreatedAt time.Time
	ID        uuid.UUID
	RoomID    uuid.UUID
	BookingID uuid.UUID
	ISActive  bool
}

type RoomLockShort struct {
	ExpiresAt time.Time
	CreatedAt time.Time
	ID        uuid.UUID
	ISActive  bool
}

func (roomLock *CreateRoomLock) ToRead() RoomLockDetail {
	return RoomLockDetail{
		RoomID:    roomLock.RoomID,
		BookingID: roomLock.BookingID,
		StayRange: roomLock.StayRange,
		ExpiresAt: roomLock.ExpiresAt,
	}
}
