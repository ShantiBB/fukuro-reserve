package mapper

import (
	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
	"hotel/internal/repository/models"
)

func RoomCreateRequestToEntity(req request.RoomCreate) models.CreateRoom {
	return models.CreateRoom{
		Title:       *req.Title,
		Description: req.Description,
		RoomNumber:  *req.RoomNumber,
		Type:        *req.Type,
		Price:       *req.Price,
		Capacity:    *req.Capacity,
		AreaSqm:     *req.AreaSqm,
		Floor:       *req.Floor,
		Amenities:   req.Amenities,
		Images:      req.Images,
	}
}

func RoomUpdateRequestToEntity(req request.RoomUpdate) models.UpdateRoom {
	return models.UpdateRoom{
		Title:       *req.Title,
		Description: *req.Description,
		RoomNumber:  *req.RoomNumber,
		Type:        *req.Type,
		Price:       *req.Price,
		Capacity:    *req.Capacity,
		AreaSqm:     *req.AreaSqm,
		Floor:       *req.Floor,
		Amenities:   req.Amenities,
		Images:      req.Images,
	}
}

func RoomStatusUpdateRequestToEntity(req request.RoomStatusUpdate) models.UpdateRoomStatus {
	return models.UpdateRoomStatus{
		Status: req.Status,
	}
}

func RoomEntityToResponse(req models.Room) response.Room {
	return response.Room{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		RoomNumber:  req.RoomNumber,
		Type:        req.Type,
		Status:      req.Status,
		Price:       req.Price,
		Capacity:    req.Capacity,
		AreaSqm:     req.AreaSqm,
		Floor:       req.Floor,
		Amenities:   req.Amenities,
		Images:      req.Images,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}
}

func RoomShortEntityToShortResponse(req models.RoomShort) response.RoomShort {
	return response.RoomShort{
		ID:         req.ID,
		Title:      req.Title,
		RoomNumber: req.RoomNumber,
		Type:       req.Type,
		Status:     req.Status,
		Price:      req.Price,
		Capacity:   req.Capacity,
		AreaSqm:    req.AreaSqm,
		Amenities:  req.Amenities,
		Images:     req.Images,
	}
}

func RoomUpdateEntityToResponse(req models.UpdateRoom) response.RoomUpdate {
	return response.RoomUpdate{
		Title:       req.Title,
		Description: &req.Description,
		RoomNumber:  req.RoomNumber,
		Type:        req.Type,
		Price:       req.Price,
		Capacity:    req.Capacity,
		AreaSqm:     req.AreaSqm,
		Floor:       req.Floor,
		Amenities:   req.Amenities,
		Images:      req.Images,
	}
}

func RoomStatusUpdateEntityToResponse(req models.UpdateRoomStatus) response.RoomStatusUpdate {
	return response.RoomStatusUpdate{
		Status: req.Status,
	}
}
