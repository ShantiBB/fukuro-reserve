package dto

import "time"

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
