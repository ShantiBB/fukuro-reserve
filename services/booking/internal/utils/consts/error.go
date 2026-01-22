package consts

import "errors"

const (
	MsgNilObject                    = "object cannot be nil"
	MsgBookingNotFound              = "booking not found"
	MsgBookingRoomNotFound          = "booking room not found"
	MsgRoomLockNotFound             = "room lock room not found"
	MsgRoomLockAlreadyExist         = "room lock already exists"
	MsgInvalidDates                 = "invalid booking dates"
	MsgPriceChanged                 = "expected total amount does not match calculated total"
	MsgConflictBookingRooms         = "all rooms must have same booking_id"
	MsgInvalidHotelID               = "invalid hotel ID"
	MsgInvalidBookingID             = "invalid booking ID"
	MsgInvalidBookingRoomID         = "invalid booking room ID"
	MsgInvalidBookingStatus         = "invalid booking status"
	MsgInvalidPricePerNightID       = "invalid price per night. example: 123.45"
	MsgInvalidExpectedTotalAmountID = "invalid expected total amount. example: 123.45"
	MsgInternalServer               = "internal server error"
)

var (
	ErrNilObject                    = errors.New(MsgNilObject)
	ErrBookingNotFound              = errors.New(MsgBookingNotFound)
	ErrBookingRoomNotFound          = errors.New(MsgBookingRoomNotFound)
	ErrRoomLockNotFound             = errors.New(MsgRoomLockNotFound)
	ErrRoomLockAlreadyExist         = errors.New(MsgRoomLockAlreadyExist)
	ErrInvalidDates                 = errors.New(MsgInvalidDates)
	ErrPriceChanged                 = errors.New(MsgPriceChanged)
	ErrConflictBookingRooms         = errors.New(MsgConflictBookingRooms)
	ErrInvalidHotelID               = errors.New(MsgInvalidHotelID)
	ErrInvalidBookingID             = errors.New(MsgInvalidBookingID)
	ErrInvalidBookingRoomID         = errors.New(MsgInvalidBookingRoomID)
	ErrInvalidBookingStatus         = errors.New(MsgInvalidBookingStatus)
	ErrInvalidPricePerNightID       = errors.New(MsgInvalidPricePerNightID)
	ErrInvalidExpectedTotalAmountID = errors.New(MsgInvalidExpectedTotalAmountID)
	ErrInternalServer               = errors.New(MsgInternalServer)
)
