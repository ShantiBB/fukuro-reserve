package models

type RoomType string
type RoomStatus string

const (
	RoomTypeUnspecified  RoomType = "ROOM_TYPE_UNSPECIFIED"
	RoomTypeSingle       RoomType = "ROOM_TYPE_SINGLE"
	RoomTypeDouble       RoomType = "ROOM_TYPE_DOUBLE"
	RoomTypeSuite        RoomType = "ROOM_TYPE_SUITE"
	RoomTypeDeluxe       RoomType = "ROOM_TYPE_DELUXE"
	RoomTypeFamily       RoomType = "ROOM_TYPE_FAMILY"
	RoomTypePresidential RoomType = "ROOM_TYPE_PRESIDENTIAL"

	RoomStatusUnspecified RoomStatus = "ROOM_STATUS_UNSPECIFIED"
	RoomStatusAvailable   RoomStatus = "ROOM_STATUS_AVAILABLE"
	RoomStatusOccupied    RoomStatus = "ROOM_STATUS_OCCUPIED"
	RoomStatusMaintenance RoomStatus = "ROOM_STATUS_MAINTENANCE"
	RoomStatusCleaning    RoomStatus = "ROOM_STATUS_CLEANING"
)

var RoomTypeValues = []RoomType{
	RoomTypeSingle,
	RoomTypeDouble,
	RoomTypeSuite,
	RoomTypeDeluxe,
	RoomTypeFamily,
	RoomTypePresidential,
}

var RoomStatusValues = []RoomStatus{
	RoomStatusAvailable,
	RoomStatusOccupied,
	RoomStatusMaintenance,
	RoomStatusCleaning,
}
