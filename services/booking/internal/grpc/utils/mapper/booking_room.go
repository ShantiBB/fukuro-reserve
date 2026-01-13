package mapper

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
	"booking/pkg/utils/consts"
)

func BookingRoomToProto(r *models.BookingRoom) *bookingv1.CreateBookingRoomResponse {
	return &bookingv1.CreateBookingRoomResponse{
		RoomId:        r.RoomID.String(),
		Adults:        uint32(r.Adults),
		Children:      uint32(r.Children),
		PricePerNight: r.PricePerNight.String(),
	}
}

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
