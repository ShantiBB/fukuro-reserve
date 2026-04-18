package mapper

import (
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

func RoomResponseFromProto(room *hotelv1.Room) *dto.RoomResponse {
	if room == nil {
		return nil
	}

	resp := &dto.RoomResponse{
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

func RoomsShortResponseFromProto(resp *hotelv1.GetRoomsResponse) *dto.RoomsResponse {
	if resp == nil {
		return nil
	}

	rooms := make([]*dto.RoomShortResponse, len(resp.Rooms))
	for i, room := range resp.Rooms {
		rooms[i] = roomShortResponseFromProto(room)
	}

	return &dto.RoomsResponse{Rooms: rooms}
}

func UpdateRoomResponseFromProto(room *hotelv1.UpdateRoom) *dto.RoomResponse {
	if room == nil {
		return nil
	}

	return &dto.RoomResponse{
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

func roomShortResponseFromProto(room *hotelv1.RoomShort) *dto.RoomShortResponse {
	if room == nil {
		return nil
	}

	return &dto.RoomShortResponse{
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
