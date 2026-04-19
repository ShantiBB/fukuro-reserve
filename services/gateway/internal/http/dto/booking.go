package dto

import "time"

type CreateBookingRoomRequest struct {
	RoomId        string `json:"room_id" example:"1f8fad5b-d9cb-469f-a165-70867728950e"`
	PricePerNight string `json:"price_per_night" example:"120.00"`
	Adults        uint32 `json:"adults" example:"2"`
	Children      uint32 `json:"children" example:"1"`
}

type CreateBookingRequest struct {
	CheckIn             time.Time                   `json:"check_in" example:"2026-05-10T14:00:00Z"`
	CheckOut            time.Time                   `json:"check_out" example:"2026-05-13T11:00:00Z"`
	GuestName           string                      `json:"guest_name" example:"Ivan Petrov"`
	GuestEmail          string                      `json:"guest_email,omitempty" example:"ivan.petrov@example.com"`
	GuestPhone          string                      `json:"guest_phone,omitempty" example:"+79991234567"`
	Currency            string                      `json:"currency" example:"RUB"`
	ExpectedTotalAmount string                      `json:"expected_total_amount" example:"360.00"`
	Rooms               []*CreateBookingRoomRequest `json:"rooms"`
	UserId              int64                       `json:"user_id" example:"1"`
}

type QuoteBookingRequest struct {
	CheckIn  time.Time                   `json:"check_in" example:"2026-05-10T14:00:00Z"`
	CheckOut time.Time                   `json:"check_out" example:"2026-05-13T11:00:00Z"`
	Currency string                      `json:"currency" example:"RUB"`
	Rooms    []*CreateBookingRoomRequest `json:"rooms"`
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

type QuoteBookingRoomResponse struct {
	RoomId        string `json:"room_id"`
	PricePerNight string `json:"price_per_night"`
	TotalAmount   string `json:"total_amount"`
	Adults        uint32 `json:"adults"`
	Children      uint32 `json:"children"`
}

type QuoteBookingResponse struct {
	CheckIn     time.Time                   `json:"check_in"`
	CheckOut    time.Time                   `json:"check_out"`
	Currency    string                      `json:"currency"`
	TotalAmount string                      `json:"total_amount"`
	Rooms       []*QuoteBookingRoomResponse `json:"rooms"`
	Nights      uint32                      `json:"nights"`
}

type AvailabilityResponse struct {
	CheckIn  time.Time            `json:"check_in"`
	CheckOut time.Time            `json:"check_out"`
	Rooms    []*RoomShortResponse `json:"rooms"`
}
