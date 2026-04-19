package helper

import (
	"github.com/google/uuid"

	"github.com/ShantiBB/fukuro-reserve/services/hotel/pkg/lib/utils/consts"
)

func ParseRoomID(roomID string) (uuid.UUID, error) {
	id, err := uuid.Parse(roomID)
	if err != nil {
		return uuid.UUID{}, consts.ErrInvalidRoomID
	}

	return id, nil
}

func ParseHotelID(hotelID string) (uuid.UUID, error) {
	id, err := uuid.Parse(hotelID)
	if err != nil {
		return uuid.UUID{}, consts.ErrInvalidHotelID
	}

	return id, nil
}
