package models

type RoomType string
type RoomStatus string

const (
	RoomTypeSingle       RoomType = "single"
	RoomTypeDouble       RoomType = "double"
	RoomTypeSuite        RoomType = "suite"
	RoomTypeDeluxe       RoomType = "deluxe"
	RoomTypeFamily       RoomType = "family"
	RoomTypePresidential RoomType = "presidential"

	RoomStatusAvailable   RoomStatus = "available"
	RoomStatusOccupied    RoomStatus = "occupied"
	RoomStatusMaintenance RoomStatus = "maintenance"
	RoomStatusCleaning    RoomStatus = "cleaning"
)
