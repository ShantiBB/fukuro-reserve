package mapper

import (
	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
)

func BookingResponseFromProto(booking *bookingv1.Booking) *dto.BookingResponse {
	if booking == nil {
		return nil
	}

	resp := &dto.BookingResponse{
		Id:                  booking.Id,
		UserId:              booking.UserId,
		HotelId:             booking.HotelId,
		Status:              booking.Status.String(),
		GuestName:           booking.GuestName,
		Currency:            booking.Currency,
		ExpectedTotalAmount: booking.ExpectedTotalAmount,
		FinalTotalAmount:    booking.FinalTotalAmount,
	}
	if booking.CheckIn != nil {
		resp.CheckIn = booking.CheckIn.AsTime()
	}
	if booking.CheckOut != nil {
		resp.CheckOut = booking.CheckOut.AsTime()
	}
	if booking.GuestEmail != nil {
		resp.GuestEmail = *booking.GuestEmail
	}
	if booking.GuestPhone != nil {
		resp.GuestPhone = *booking.GuestPhone
	}
	if booking.CreatedAt != nil {
		resp.CreatedAt = booking.CreatedAt.AsTime()
	}
	if booking.UpdatedAt != nil {
		resp.UpdatedAt = booking.UpdatedAt.AsTime()
	}

	rooms := make([]*dto.BookingRoomResponse, len(booking.BookingRooms))
	for i, room := range booking.BookingRooms {
		rooms[i] = bookingRoomWithLockResponseFromProto(room)
	}
	resp.BookingRooms = rooms

	return resp
}

func BookingsShortResponseFromProto(resp *bookingv1.GetBookingsResponse) *dto.BookingsResponse {
	if resp == nil {
		return nil
	}

	bookings := make([]*dto.BookingShortResponse, len(resp.Bookings))
	for i, booking := range resp.Bookings {
		bookings[i] = bookingShortResponseFromProto(booking)
	}

	return &dto.BookingsResponse{Bookings: bookings}
}

func QuoteBookingResponseFromProto(resp *bookingv1.QuoteBookingResponse) *dto.QuoteBookingResponse {
	if resp == nil {
		return nil
	}

	out := &dto.QuoteBookingResponse{
		Currency:    resp.Currency,
		TotalAmount: resp.TotalAmount,
		Nights:      resp.Nights,
		Rooms:       make([]*dto.QuoteBookingRoomResponse, len(resp.Rooms)),
	}
	if resp.CheckIn != nil {
		out.CheckIn = resp.CheckIn.AsTime()
	}
	if resp.CheckOut != nil {
		out.CheckOut = resp.CheckOut.AsTime()
	}
	for i, room := range resp.Rooms {
		out.Rooms[i] = quoteBookingRoomResponseFromProto(room)
	}

	return out
}

func bookingShortResponseFromProto(booking *bookingv1.BookingShort) *dto.BookingShortResponse {
	if booking == nil {
		return nil
	}

	resp := &dto.BookingShortResponse{
		Id:                  booking.Id,
		UserId:              booking.UserId,
		HotelId:             booking.HotelId,
		Status:              booking.Status.String(),
		GuestName:           booking.GuestName,
		Currency:            booking.Currency,
		ExpectedTotalAmount: booking.ExpectedTotalAmount,
		FinalTotalAmount:    booking.FinalTotalAmount,
	}
	if booking.CheckIn != nil {
		resp.CheckIn = booking.CheckIn.AsTime()
	}
	if booking.CheckOut != nil {
		resp.CheckOut = booking.CheckOut.AsTime()
	}
	if booking.GuestEmail != nil {
		resp.GuestEmail = *booking.GuestEmail
	}
	if booking.GuestPhone != nil {
		resp.GuestPhone = *booking.GuestPhone
	}

	rooms := make([]*dto.BookingRoomResponse, len(booking.BookingRooms))
	for i, room := range booking.BookingRooms {
		rooms[i] = bookingRoomResponseFromProto(room)
	}
	resp.BookingRooms = rooms

	return resp
}

func bookingRoomResponseFromProto(room *bookingv1.BookingRoom) *dto.BookingRoomResponse {
	if room == nil {
		return nil
	}

	return &dto.BookingRoomResponse{
		Id:            room.Id,
		RoomId:        room.RoomId,
		Adults:        room.Adults,
		Children:      room.Children,
		PricePerNight: room.PricePerNight,
	}
}

func quoteBookingRoomResponseFromProto(room *bookingv1.QuoteBookingRoom) *dto.QuoteBookingRoomResponse {
	if room == nil {
		return nil
	}

	return &dto.QuoteBookingRoomResponse{
		RoomId:        room.RoomId,
		Adults:        room.Adults,
		Children:      room.Children,
		PricePerNight: room.PricePerNight,
		TotalAmount:   room.TotalAmount,
	}
}

func bookingRoomWithLockResponseFromProto(room *bookingv1.BookingRoomWithLock) *dto.BookingRoomResponse {
	if room == nil {
		return nil
	}

	return &dto.BookingRoomResponse{
		Id:            room.Id,
		RoomId:        room.RoomId,
		Adults:        room.Adults,
		Children:      room.Children,
		PricePerNight: room.PricePerNight,
	}
}
