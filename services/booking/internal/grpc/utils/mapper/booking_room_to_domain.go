package mapper

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
	"booking/pkg/utils/consts"
)

func CreateBookingRoomsToDomain(rooms []*bookingv1.CreateBookingRoom) ([]*models.CreateBookingRoom, error) {
	result := make([]*models.CreateBookingRoom, len(rooms))

	for i, r := range rooms {
		roomID, err := uuid.Parse(r.RoomId)
		if err != nil {
			return nil, consts.ErrInvalidBookingRoomID
		}

		price, err := decimal.NewFromString(r.PricePerNight)
		if err != nil {
			return nil, consts.ErrInvalidPricePerNightID
		}

		result[i] = &models.CreateBookingRoom{
			RoomID:        roomID,
			Adults:        r.Adults,
			Children:      r.Children,
			PricePerNight: price,
		}
	}

	return result, nil
}
