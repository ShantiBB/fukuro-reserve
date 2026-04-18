package consts

const (
	FieldBookingHotelID      = "hotel_id"
	FieldBookingCurrency     = "currency"
	FieldExpectedTotalAmount = "expected_total_amount"
	FieldBookingGuestEmail   = "guest_email"
	ErrRoomsIndexedRequired  = "rooms[%d] is required"
	ErrRoomsIndexedWrap      = "rooms[%d]: %w"
)

const (
	ErrBookingGuestNameRequired = "guest_name is required"
	ErrBookingCheckInFuture     = "check_in must be in the future"
	ErrBookingCheckOutAfterIn   = "check_out must be after check_in"
	ErrBookingAtLeastOneRoom    = "at least one room is required"
)
