package response

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type BookingRoom struct {
	RoomID        uuid.UUID       `json:"room_id"`
	Adults        uint8           `json:"adults"`
	Children      uint8           `json:"children"`
	PricePerNight decimal.Decimal `json:"price_per_night"`
}
