package mapper

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
	"booking/pkg/utils/consts"
)

func CreateBookingRoomsToDomain(rooms []*bookingv1.CreateBookingRoom) ([]models.CreateBookingRoom, error) {
	result := make([]models.CreateBookingRoom, 0, len(rooms))

	for _, r := range rooms {
		roomID, err := uuid.Parse(r.RoomId)
		if err != nil {
			return nil, consts.InvalidRoomID
		}

		price, err := decimal.NewFromString(r.PricePerNight)
		if err != nil {
			return nil, consts.InvalidPricePerNightID
		}

		result = append(
			result, models.CreateBookingRoom{
				RoomID:        roomID,
				Adults:        uint8(r.Adults),
				Children:      uint8(r.Children),
				PricePerNight: price,
			},
		)
	}

	return result, nil
}

func GetBookingRoomsRequestToDomain(req *bookingv1.GetBookingRoomsRequest) (uuid.UUID, error) {
	bookingID, err := uuid.Parse(req.BookingId)
	if err != nil {
		return uuid.UUID{}, err
	}

	return bookingID, nil
}
