package dto

import "time"

type CreateRoomRequest struct {
	RoomNumber  string   `json:"room_number" example:"101"`
	Type        string   `json:"type" example:"ROOM_TYPE_SINGLE"`
	Title       string   `json:"title" example:"Standard Single Room"`
	Description string   `json:"description,omitempty" example:"Quiet room with city view"`
	Price       string   `json:"price" example:"150.00"`
	Amenities   []string `json:"amenities" example:"wifi,tv,air-conditioning"`
	Images      []string `json:"images" example:"https://example.com/rooms/101.jpg"`
	Floor       int64    `json:"floor" example:"1"`
	Capacity    int64    `json:"capacity" example:"2"`
	AreaSqm     float32  `json:"area_sqm" example:"18.5"`
}

type UpdateRoomRequest struct {
	Title       string   `json:"title" example:"Standard Double Room"`
	RoomNumber  string   `json:"room_number" example:"101"`
	Type        string   `json:"type" example:"ROOM_TYPE_DOUBLE"`
	Description string   `json:"description" example:"Refreshed interior and upgraded bed"`
	Price       string   `json:"price" example:"175.00"`
	Amenities   []string `json:"amenities" example:"wifi,tv,air-conditioning,coffee-machine"`
	Images      []string `json:"images" example:"https://example.com/rooms/101-updated.jpg"`
	Capacity    int64    `json:"capacity" example:"3"`
	Floor       int64    `json:"floor" example:"1"`
	AreaSqm     float32  `json:"area_sqm" example:"20.0"`
}

type UpdateRoomStatusRequest struct {
	Status string `json:"status" example:"ROOM_STATUS_MAINTENANCE"`
}

type RoomResponse struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Price       string    `json:"price"`
	RoomNumber  string    `json:"room_number"`
	Status      string    `json:"status"`
	Type        string    `json:"type"`
	Id          string    `json:"id"`
	Description string    `json:"description,omitempty"`
	Title       string    `json:"title"`
	Amenities   []string  `json:"amenities"`
	Images      []string  `json:"images"`
	Capacity    int64     `json:"capacity"`
	Floor       int64     `json:"floor"`
	AreaSqm     float32   `json:"area_sqm"`
}

type RoomShortResponse struct {
	Id         string   `json:"id"`
	Title      string   `json:"title"`
	RoomNumber string   `json:"room_number"`
	Status     string   `json:"status"`
	Type       string   `json:"type"`
	Price      string   `json:"price"`
	Amenities  []string `json:"amenities"`
	Images     []string `json:"images"`
	Capacity   int64    `json:"capacity"`
	AreaSqm    float32  `json:"area_sqm"`
}

type RoomsResponse struct {
	Rooms []*RoomShortResponse `json:"rooms"`
}
