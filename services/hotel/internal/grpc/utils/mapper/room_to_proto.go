package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/repository/models"
)

func roomStatusToProto(status models.RoomStatus) hotelv1.RoomStatus {
	var s hotelv1.RoomStatus
	switch status {
	case models.RoomStatusAvailable:
		s = hotelv1.RoomStatus_ROOM_STATUS_AVAILABLE
	case models.RoomStatusOccupied:
		s = hotelv1.RoomStatus_ROOM_STATUS_OCCUPIED
	case models.RoomStatusMaintenance:
		s = hotelv1.RoomStatus_ROOM_STATUS_MAINTENANCE
	case models.RoomStatusCleaning:
		s = hotelv1.RoomStatus_ROOM_STATUS_CLEANING
	default:
		s = hotelv1.RoomStatus_ROOM_STATUS_UNSPECIFIED
	}
	return s
}

func roomTypeToProto(status models.RoomType) hotelv1.RoomType {
	var s hotelv1.RoomType
	switch status {
	case models.RoomTypeSingle:
		s = hotelv1.RoomType_ROOM_TYPE_SINGLE
	case models.RoomTypeDouble:
		s = hotelv1.RoomType_ROOM_TYPE_DOUBLE
	case models.RoomTypeSuite:
		s = hotelv1.RoomType_ROOM_TYPE_SUITE
	case models.RoomTypeFamily:
		s = hotelv1.RoomType_ROOM_TYPE_FAMILY
	case models.RoomTypePresidential:
		s = hotelv1.RoomType_ROOM_TYPE_PRESIDENTIAL
	default:
		s = hotelv1.RoomType_ROOM_TYPE_UNSPECIFIED
	}
	return s
}

func RoomResponseToProto(resp *models.Room) *hotelv1.Room {
	return &hotelv1.Room{
		Id:          resp.ID.String(),
		CreatedAt:   timestamppb.New(resp.CreatedAt),
		UpdatedAt:   timestamppb.New(resp.UpdatedAt),
		Description: resp.Description,
		Price:       resp.Price.String(),
		Type:        roomTypeToProto(resp.Type),
		Status:      roomStatusToProto(resp.Status),
		RoomNumber:  resp.RoomNumber,
		Title:       resp.Title,
		Amenities:   resp.Amenities,
		Images:      resp.Images,
		Capacity:    int64(resp.Capacity),
		AreaSqm:     float32(resp.AreaSqm),
		Floor:       int64(resp.Floor),
	}
}

func RoomShortResponseToProto(resp *models.RoomShort) *hotelv1.RoomShort {
	return &hotelv1.RoomShort{
		Id:         resp.ID.String(),
		Title:      resp.Title,
		RoomNumber: resp.RoomNumber,
		Type:       roomTypeToProto(resp.Type),
		Status:     roomStatusToProto(resp.Status),
		Price:      resp.Price.String(),
		Amenities:  resp.Amenities,
		Images:     resp.Images,
		Capacity:   int64(resp.Capacity),
		AreaSqm:    float32(resp.AreaSqm),
	}
}

func RoomsResponseToProto(resp []*models.RoomShort) []*hotelv1.RoomShort {
	rooms := make([]*hotelv1.RoomShort, len(resp))
	for i, h := range resp {
		rooms[i] = RoomShortResponseToProto(h)
	}

	return rooms
}

func UpdateRoomResponseToProto(resp *models.UpdateRoom) *hotelv1.UpdateRoom {
	return &hotelv1.UpdateRoom{
		Description: resp.Description,
		Title:       resp.Title,
		RoomNumber:  resp.RoomNumber,
		Price:       resp.Price.String(),
		Type:        roomTypeToProto(resp.Type),
		Amenities:   resp.Amenities,
		Images:      resp.Images,
		Capacity:    int64(resp.Capacity),
		AreaSqm:     float32(resp.AreaSqm),
		Floor:       int64(resp.Floor),
	}
}

func UpdateRoomStatusResponseToProto(resp models.UpdateRoomStatus) hotelv1.RoomStatus {
	return roomStatusToProto(resp.Status)
}
