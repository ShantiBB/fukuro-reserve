package mapper

import (
	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
)

func BookingRoomInfoToProto(r *models.BookingRoomInfo) *bookingv1.BookingRoomInfo {
	return &bookingv1.BookingRoomInfo{
		RoomId:        r.RoomID.String(),
		Adults:        uint32(r.Adults),
		Children:      uint32(r.Children),
		PricePerNight: r.PricePerNight.String(),
	}
}

func BookingRoomsInfoToProto(rooms []models.BookingRoomInfo) []*bookingv1.BookingRoomInfo {
	result := make([]*bookingv1.BookingRoomInfo, 0, len(rooms))
	for _, r := range rooms {
		result = append(
			result, &bookingv1.BookingRoomInfo{
				Id:            r.ID.String(),
				RoomId:        r.RoomID.String(),
				Adults:        uint32(r.Adults),
				Children:      uint32(r.Children),
				PricePerNight: r.PricePerNight.String(),
			},
		)
	}
	return result
}

func BookingRoomsFullInfoToProto(rooms []models.BookingRoomFullInfo) []*bookingv1.BookingRoomFullInfo {
	result := make([]*bookingv1.BookingRoomFullInfo, 0, len(rooms))
	for _, r := range rooms {
		result = append(
			result, &bookingv1.BookingRoomFullInfo{
				Id:            r.ID.String(),
				RoomId:        r.RoomID.String(),
				Adults:        uint32(r.Adults),
				Children:      uint32(r.Children),
				PricePerNight: r.PricePerNight.String(),
				RoomLock:      RoomLockToProto(&r.RoomLock),
			},
		)
	}
	return result
}
