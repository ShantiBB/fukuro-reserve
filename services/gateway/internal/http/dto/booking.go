package dto

import (
	"time"

	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
)

// Booking DTOs.
type CreateBookingRoomRequest struct {
	RoomId        string `json:"room_id"`
	PricePerNight string `json:"price_per_night"`
	Adults        uint32 `json:"adults"`
	Children      uint32 `json:"children"`
}

type CreateBookingRequest struct {
	CheckIn             time.Time                   `json:"check_in"`
	CheckOut            time.Time                   `json:"check_out"`
	HotelId             string                      `json:"hotel_id"`
	GuestName           string                      `json:"guest_name"`
	GuestEmail          string                      `json:"guest_email,omitempty"`
	GuestPhone          string                      `json:"guest_phone,omitempty"`
	Currency            string                      `json:"currency"`
	ExpectedTotalAmount string                      `json:"expected_total_amount"`
	Rooms               []*CreateBookingRoomRequest `json:"rooms"`
	UserId              int64                       `json:"user_id"`
}

type BookingRoomResponse struct {
	Id            string `json:"id"`
	RoomId        string `json:"room_id"`
	PricePerNight string `json:"price_per_night"`
	Adults        uint32 `json:"adults"`
	Children      uint32 `json:"children"`
}

type BookingResponse struct {
	CheckOut            time.Time              `json:"check_out"`
	UpdatedAt           time.Time              `json:"updated_at"`
	CreatedAt           time.Time              `json:"created_at"`
	CheckIn             time.Time              `json:"check_in"`
	GuestName           string                 `json:"guest_name"`
	Status              string                 `json:"status"`
	Id                  string                 `json:"id"`
	GuestEmail          string                 `json:"guest_email,omitempty"`
	GuestPhone          string                 `json:"guest_phone,omitempty"`
	Currency            string                 `json:"currency"`
	ExpectedTotalAmount string                 `json:"expected_total_amount"`
	FinalTotalAmount    string                 `json:"final_total_amount,omitempty"`
	HotelId             string                 `json:"hotel_id"`
	BookingRooms        []*BookingRoomResponse `json:"booking_rooms"`
	UserId              int64                  `json:"user_id"`
}

type BookingShortResponse struct {
	CheckIn             time.Time              `json:"check_in"`
	CheckOut            time.Time              `json:"check_out"`
	GuestName           string                 `json:"guest_name"`
	HotelId             string                 `json:"hotel_id"`
	Status              string                 `json:"status"`
	Id                  string                 `json:"id"`
	GuestEmail          string                 `json:"guest_email,omitempty"`
	GuestPhone          string                 `json:"guest_phone,omitempty"`
	Currency            string                 `json:"currency"`
	ExpectedTotalAmount string                 `json:"expected_total_amount"`
	FinalTotalAmount    string                 `json:"final_total_amount,omitempty"`
	BookingRooms        []*BookingRoomResponse `json:"booking_rooms"`
	UserId              int64                  `json:"user_id"`
}

type BookingsResponse struct {
	Bookings []*BookingShortResponse `json:"bookings"`
}

func bookingRoomResponseFromProto(room *bookingv1.BookingRoom) *BookingRoomResponse {
	if room == nil {
		return nil
	}

	return &BookingRoomResponse{
		Id:            room.Id,
		RoomId:        room.RoomId,
		Adults:        room.Adults,
		Children:      room.Children,
		PricePerNight: room.PricePerNight,
	}
}

func bookingRoomWithLockResponseFromProto(room *bookingv1.BookingRoomWithLock) *BookingRoomResponse {
	if room == nil {
		return nil
	}

	return &BookingRoomResponse{
		Id:            room.Id,
		RoomId:        room.RoomId,
		Adults:        room.Adults,
		Children:      room.Children,
		PricePerNight: room.PricePerNight,
	}
}

func BookingResponseFromProto(booking *bookingv1.Booking) *BookingResponse {
	if booking == nil {
		return nil
	}

	resp := &BookingResponse{
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

	rooms := make([]*BookingRoomResponse, len(booking.BookingRooms))
	for i, r := range booking.BookingRooms {
		rooms[i] = bookingRoomWithLockResponseFromProto(r)
	}
	resp.BookingRooms = rooms

	return resp
}

func bookingShortResponseFromProto(booking *bookingv1.BookingShort) *BookingShortResponse {
	if booking == nil {
		return nil
	}

	resp := &BookingShortResponse{
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

	rooms := make([]*BookingRoomResponse, len(booking.BookingRooms))
	for i, r := range booking.BookingRooms {
		rooms[i] = bookingRoomResponseFromProto(r)
	}
	resp.BookingRooms = rooms

	return resp
}

func BookingsShortResponseFromProto(resp *bookingv1.GetBookingsResponse) *BookingsResponse {
	if resp == nil {
		return nil
	}

	bookings := make([]*BookingShortResponse, len(resp.Bookings))
	for i, b := range resp.Bookings {
		bookings[i] = bookingShortResponseFromProto(b)
	}

	return &BookingsResponse{Bookings: bookings}
}
