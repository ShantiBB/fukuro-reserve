package hotel

import (
	"context"
	"fmt"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/mapper"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

func (s *Service) CreateRoom(ctx context.Context, req dto.CreateRoomRequest) (*dto.RoomResponse, error) {
	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	if req.CountryCode == "" || req.CitySlug == "" || req.HotelSlug == "" {
		return nil, fmt.Errorf("country_code, city_slug, and hotel_slug are required")
	}

	var price float32
	if req.Price != "" {
		var p float64
		if _, err := fmt.Sscanf(req.Price, "%f", &p); err != nil {
			return nil, fmt.Errorf("invalid price")
		}
		price = float32(p)
	}

	resp, err := s.clients.Room.CreateRoom(ctx, &hotelv1.CreateRoomRequest{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		HotelSlug:   req.HotelSlug,
		Title:       req.Title,
		Description: description,
		RoomNumber:  req.RoomNumber,
		Type:        hotelv1.RoomType(hotelv1.RoomType_value[req.Type]),
		Price:       price,
		Capacity:    req.Capacity,
		AreaSqm:     req.AreaSqm,
		Floor:       req.Floor,
		Amenities:   req.Amenities,
		Images:      req.Images,
	})
	if err != nil {
		return nil, err
	}

	return mapper.RoomResponseFromProto(resp.Room), nil
}

func (s *Service) GetRooms(ctx context.Context, countryCode, citySlug, hotelSlug string, page, limit uint64) (*dto.RoomsResponse, error) {
	if page == 0 {
		page = s.pagination.DefaultPage
	}
	if limit == 0 {
		limit = s.pagination.DefaultPageSize
	}

	resp, err := s.clients.Room.GetRooms(ctx, &hotelv1.GetRoomsRequest{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		HotelSlug:   hotelSlug,
		Page:        page,
		Limit:       limit,
	})
	if err != nil {
		return nil, err
	}

	return mapper.RoomsShortResponseFromProto(resp), nil
}

func (s *Service) GetRoom(ctx context.Context, roomID string) (*dto.RoomResponse, error) {
	resp, err := s.clients.Room.GetRoom(ctx, &hotelv1.GetRoomRequest{Id: roomID})
	if err != nil {
		return nil, err
	}

	return mapper.RoomResponseFromProto(resp.Room), nil
}

func (s *Service) UpdateRoom(ctx context.Context, roomID string, req dto.UpdateRoomRequest) (*dto.RoomResponse, error) {
	resp, err := s.clients.Room.UpdateRoom(ctx, &hotelv1.UpdateRoomRequest{
		Id:          roomID,
		Title:       req.Title,
		RoomNumber:  req.RoomNumber,
		Type:        hotelv1.RoomType(hotelv1.RoomType_value[req.Type]),
		Description: req.Description,
		Price:       req.Price,
		Capacity:    req.Capacity,
		AreaSqm:     req.AreaSqm,
		Floor:       req.Floor,
		Amenities:   req.Amenities,
		Images:      req.Images,
	})
	if err != nil {
		return nil, err
	}

	return mapper.UpdateRoomResponseFromProto(resp.Room), nil
}

func (s *Service) UpdateRoomStatus(ctx context.Context, roomID string, req dto.UpdateRoomStatusRequest) (*dto.StatusResponse, error) {
	resp, err := s.clients.Room.UpdateRoomStatus(ctx, &hotelv1.UpdateRoomStatusRequest{
		Id:     roomID,
		Status: hotelv1.RoomStatus(hotelv1.RoomStatus_value[req.Status]),
	})
	if err != nil {
		return nil, err
	}

	return &dto.StatusResponse{Status: resp.Status.String()}, nil
}

func (s *Service) DeleteRoom(ctx context.Context, roomID string) error {
	_, err := s.clients.Room.DeleteRoom(ctx, &hotelv1.DeleteRoomRequest{Id: roomID})
	return err
}
