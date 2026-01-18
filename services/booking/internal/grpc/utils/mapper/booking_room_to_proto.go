package mapper

import (
	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
)

func BookingRoomToProto(r *models.BookingRoom) *bookingv1.BookingRoom {
	return &bookingv1.BookingRoom{
		Id:            r.ID.String(),
		RoomId:        r.RoomID.String(),
		Adults:        r.Adults,
		Children:      r.Children,
		PricePerNight: r.PricePerNight.String(),
	}
}

func BookingRoomWithLockToProto(r *models.BookingRoomWithLock) *bookingv1.BookingRoomWithLock {
	return &bookingv1.BookingRoomWithLock{
		Id:            r.ID.String(),
		RoomId:        r.RoomID.String(),
		Adults:        r.Adults,
		Children:      r.Children,
		PricePerNight: r.PricePerNight.String(),
		RoomLock:      RoomLockToProto(&r.RoomLock),
	}
}

func BookingRoomsWithLockToProto(rooms []*models.BookingRoomWithLock) []*bookingv1.BookingRoomWithLock {
	result := make([]*bookingv1.BookingRoomWithLock, len(rooms))
	for i, r := range rooms {
		result[i] = BookingRoomWithLockToProto(r)
	}
	return result
}

func BookingRoomsToProto(rooms []*models.BookingRoom) []*bookingv1.BookingRoom {
	result := make([]*bookingv1.BookingRoom, len(rooms))
	for i, r := range rooms {
		result[i] = BookingRoomToProto(r)
	}
	return result
}
