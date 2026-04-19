package dto

import "time"

type CreateRoomRequest struct {
	RoomNumber  string   `json:"room_number"`
	Type        string   `json:"type"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Price       string   `json:"price"`
	Amenities   []string `json:"amenities"`
	Images      []string `json:"images"`
	Floor       int64    `json:"floor"`
	Capacity    int64    `json:"capacity"`
	AreaSqm     float32  `json:"area_sqm"`
}

type UpdateRoomRequest struct {
	Title       string   `json:"title"`
	RoomNumber  string   `json:"room_number"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Price       string   `json:"price"`
	Amenities   []string `json:"amenities"`
	Images      []string `json:"images"`
	Capacity    int64    `json:"capacity"`
	Floor       int64    `json:"floor"`
	AreaSqm     float32  `json:"area_sqm"`
}

type UpdateRoomStatusRequest struct {
	Status string `json:"status"`
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
