package consts

const (
	FieldRoomPrice         = "price"
	FieldRoomID            = "room_id"
	FieldRoomPricePerNight = "price_per_night"
)

const (
	ErrRoomNumberRequired = "room_number is required"
	ErrRoomTypeRequired   = "type is required"
	ErrRoomHotelSlugReq   = "hotel_slug is required"
	ErrRoomCapacityRange  = "capacity must be between 1 and 20"
	ErrRoomFloorNonNeg    = "floor must be greater than or equal to 0"
	ErrRoomAreaPositive   = "area_sqm must be greater than 0"
	ErrRoomStatusRequired = "status is required"
	ErrAdultsRange        = "adults must be between 1 and 20"
	ErrChildrenRange      = "children must be 20 or less"
)
