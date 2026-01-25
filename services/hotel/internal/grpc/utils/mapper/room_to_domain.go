package mapper

import (
	"github.com/shopspring/decimal"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/repository/models"
	"hotel/pkg/lib/utils/consts"
)

func roomTypeToDomain(status hotelv1.RoomType) models.RoomType {
	var s models.RoomType
	switch status {
	case hotelv1.RoomType_ROOM_TYPE_SINGLE:
		s = models.RoomTypeSingle
	case hotelv1.RoomType_ROOM_TYPE_DOUBLE:
		s = models.RoomTypeDouble
	case hotelv1.RoomType_ROOM_TYPE_SUITE:
		s = models.RoomTypeSuite
	case hotelv1.RoomType_ROOM_TYPE_FAMILY:
		s = models.RoomTypeFamily
	case hotelv1.RoomType_ROOM_TYPE_PRESIDENTIAL:
		s = models.RoomTypePresidential
	default:
		s = models.RoomTypeUnspecified
	}
	return s
}

func roomStatusToDomain(status hotelv1.RoomStatus) models.RoomStatus {
	var s models.RoomStatus
	switch status {
	case hotelv1.RoomStatus_ROOM_STATUS_AVAILABLE:
		s = models.RoomStatusAvailable
	case hotelv1.RoomStatus_ROOM_STATUS_OCCUPIED:
		s = models.RoomStatusOccupied
	case hotelv1.RoomStatus_ROOM_STATUS_MAINTENANCE:
		s = models.RoomStatusMaintenance
	case hotelv1.RoomStatus_ROOM_STATUS_CLEANING:
		s = models.RoomStatusCleaning
	default:
		s = models.RoomStatusUnspecified
	}
	return s
}

func CreateRoomRequestToDomain(req *hotelv1.CreateRoomRequest) *models.CreateRoom {
	return &models.CreateRoom{
		Description: req.Description,
		Title:       req.Title,
		RoomNumber:  req.RoomNumber,
		Type:        roomTypeToDomain(req.Type),
		Price:       decimal.NewFromFloat(float64(req.Price)),
		Amenities:   req.Amenities,
		Images:      req.Images,
		Capacity:    int(req.Capacity),
		AreaSqm:     float64(req.AreaSqm),
		Floor:       int(req.Floor),
	}
}

func UpdateRoomRequestToDomain(req *hotelv1.UpdateRoomRequest) (*models.UpdateRoom, error) {
	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		return nil, consts.ErrInvalidPrice
	}

	return &models.UpdateRoom{
		Description: req.Description,
		Title:       req.Title,
		RoomNumber:  req.RoomNumber,
		Type:        roomTypeToDomain(req.Type),
		Price:       price,
		Amenities:   req.Amenities,
		Images:      req.Images,
		Capacity:    int(req.Capacity),
		AreaSqm:     float64(req.AreaSqm),
		Floor:       int(req.Floor),
	}, nil
}

func UpdateRoomStatusRequestToDomain(req *hotelv1.UpdateRoomStatusRequest) models.UpdateRoomStatus {
	return models.UpdateRoomStatus{
		Status: roomStatusToDomain(req.Status),
	}
}
