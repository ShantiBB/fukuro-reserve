package dto

import (
	"time"

	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

// Room DTOs.
type CreateRoomRequest struct {
	RoomNumber  string   `json:"room_number"`
	Type        string   `json:"type"`
	HotelSlug   string   `json:"hotel_slug"`
	HotelId     string   `json:"hotel_id,omitempty"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Price       string   `json:"price"`
	CountryCode string   `json:"country_code"`
	CitySlug    string   `json:"city_slug"`
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

func RoomResponseFromProto(room *hotelv1.Room) *RoomResponse {
	if room == nil {
		return nil
	}

	resp := &RoomResponse{
		Id:         room.Id,
		Title:      room.Title,
		RoomNumber: room.RoomNumber,
		Status:     room.Status.String(),
		Type:       room.Type.String(),
		Price:      room.Price,
		Capacity:   room.Capacity,
		AreaSqm:    room.AreaSqm,
		Floor:      room.Floor,
		Amenities:  room.Amenities,
		Images:     room.Images,
	}

	if room.Description != nil {
		resp.Description = *room.Description
	}
	if room.CreatedAt != nil {
		resp.CreatedAt = room.CreatedAt.AsTime()
	}
	if room.UpdatedAt != nil {
		resp.UpdatedAt = room.UpdatedAt.AsTime()
	}

	return resp
}

func roomShortResponseFromProto(room *hotelv1.RoomShort) *RoomShortResponse {
	if room == nil {
		return nil
	}

	return &RoomShortResponse{
		Id:         room.Id,
		Title:      room.Title,
		RoomNumber: room.RoomNumber,
		Status:     room.Status.String(),
		Type:       room.Type.String(),
		Price:      room.Price,
		Capacity:   room.Capacity,
		AreaSqm:    room.AreaSqm,
		Amenities:  room.Amenities,
		Images:     room.Images,
	}
}

func RoomsShortResponseFromProto(resp *hotelv1.GetRoomsResponse) *RoomsResponse {
	if resp == nil {
		return nil
	}

	rooms := make([]*RoomShortResponse, len(resp.Rooms))
	for i, r := range resp.Rooms {
		rooms[i] = roomShortResponseFromProto(r)
	}

	return &RoomsResponse{Rooms: rooms}
}

func UpdateRoomResponseFromProto(room *hotelv1.UpdateRoom) *RoomResponse {
	if room == nil {
		return nil
	}

	return &RoomResponse{
		Title:       room.Title,
		Description: room.Description,
		RoomNumber:  room.RoomNumber,
		Type:        room.Type.String(),
		Price:       room.Price,
		Capacity:    room.Capacity,
		AreaSqm:     room.AreaSqm,
		Floor:       room.Floor,
		Amenities:   room.Amenities,
		Images:      room.Images,
	}
}
