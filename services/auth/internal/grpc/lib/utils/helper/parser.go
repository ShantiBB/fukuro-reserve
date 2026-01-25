package helper

import (
	"github.com/google/uuid"

	"hotel/pkg/lib/utils/consts"
)

func ParseRoomID(roomID string) (uuid.UUID, error) {
	id, err := uuid.Parse(roomID)
	if err != nil {
		return uuid.UUID{}, consts.ErrInvalidRoomID
	}

	return id, nil
}
