package service

import (
	"context"
	"fmt"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/mapper"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	int32Min = -1 << 31
	int32Max = 1<<31 - 1
)

func (s *Hotel) CreateHotel(ctx context.Context, req dto.CreateHotelRequest) (*dto.HotelResponse, error) {
	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	var location *hotelv1.CreateHotelLocationRequest
	if req.Location != nil {
		location = &hotelv1.CreateHotelLocationRequest{
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		}
	}

	resp, err := s.clients.Hotel.CreateHotel(ctx, &hotelv1.CreateHotelRequest{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		Title:       req.Title,
		OwnerId:     req.OwnerId,
		Description: description,
		Address:     req.Address,
		Location:    location,
	})
	if err != nil {
		return nil, err
	}

	return mapper.HotelResponseFromProto(resp.Hotel), nil
}

func (s *Hotel) GetHotels(ctx context.Context, countryCode, citySlug, sortBy string, page, limit uint64) (*dto.HotelsResponse, error) {
	if page == 0 {
		page = s.pagination.DefaultPage
	}
	if limit == 0 {
		limit = s.pagination.DefaultPageSize
	}

	resp, err := s.clients.Hotel.GetHotels(ctx, &hotelv1.GetHotelsRequest{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		SortBy:      sortBy,
		Page:        page,
		Limit:       limit,
	})
	if err != nil {
		return nil, err
	}

	return mapper.HotelsShortResponseFromProto(resp), nil
}

func (s *Hotel) GetHotelByID(ctx context.Context, countryCode, citySlug, hotelID string) (*dto.HotelResponse, error) {
	resp, err := s.clients.Hotel.GetHotelByID(ctx, &hotelv1.GetHotelByIDRequest{
		Id:          hotelID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
	})
	if err != nil {
		return nil, err
	}

	return mapper.HotelDetailByIDResponseFromProto(resp), nil
}

func (s *Hotel) GetHotelBySlug(ctx context.Context, countryCode, citySlug, hotelSlug string) (*dto.HotelResponse, error) {
	resp, err := s.clients.Hotel.GetHotel(ctx, &hotelv1.GetHotelRequest{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		HotelSlug:   hotelSlug,
	})
	if err != nil {
		return nil, err
	}

	return mapper.HotelDetailResponseFromProto(resp), nil
}

func (s *Hotel) UpdateHotelByID(ctx context.Context, countryCode, citySlug, hotelID string, req dto.UpdateHotelRequest) (*dto.HotelResponse, error) {
	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	var location *hotelv1.UpdateHotelLocationRequest
	if req.Location != nil {
		location = &hotelv1.UpdateHotelLocationRequest{
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		}
	}

	resp, err := s.clients.Hotel.UpdateHotelByID(ctx, &hotelv1.UpdateHotelByIDRequest{
		Id:          hotelID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
		Description: description,
		Address:     req.Address,
		Location:    location,
	})
	if err != nil {
		return nil, err
	}

	return mapper.UpdateHotelResponseFromProto(resp.Hotel), nil
}

func (s *Hotel) UpdateHotelTitleByID(ctx context.Context, countryCode, citySlug, hotelID string, req dto.UpdateHotelTitleRequest) (*dto.HotelResponse, error) {
	resp, err := s.clients.Hotel.UpdateHotelTitleByID(ctx, &hotelv1.UpdateHotelTitleByIDRequest{
		Id:          hotelID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
		Title:       req.Title,
	})
	if err != nil {
		return nil, err
	}

	return mapper.UpdateHotelTitleResponseFromProto(resp.Hotel), nil
}

func (s *Hotel) DeleteHotelByID(ctx context.Context, countryCode, citySlug, hotelID string) error {
	_, err := s.clients.Hotel.DeleteHotelByID(ctx, &hotelv1.DeleteHotelByIDRequest{
		Id:          hotelID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
	})
	return err
}

func (s *Hotel) CreateRoomByHotelID(ctx context.Context, countryCode, citySlug, hotelID string, req dto.CreateRoomRequest) (*dto.RoomResponse, error) {
	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	var price float32
	if req.Price != "" {
		var p float64
		if _, err := fmt.Sscanf(req.Price, "%f", &p); err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid price")
		}
		price = float32(p)
	}

	capacity, err := toInt32(req.Capacity, "capacity")
	if err != nil {
		return nil, err
	}
	floor, err := toInt32(req.Floor, "floor")
	if err != nil {
		return nil, err
	}

	resp, err := s.clients.Room.CreateRoomByHotelID(ctx, &hotelv1.CreateRoomByHotelIDRequest{
		HotelId:     hotelID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
		Title:       req.Title,
		Description: description,
		RoomNumber:  req.RoomNumber,
		Type:        hotelv1.RoomType(hotelv1.RoomType_value[req.Type]),
		Price:       price,
		Capacity:    capacity,
		AreaSqm:     float64(req.AreaSqm),
		Floor:       floor,
		Amenities:   req.Amenities,
		Images:      req.Images,
	})
	if err != nil {
		return nil, err
	}

	return mapper.RoomResponseFromProto(resp.Room), nil
}

func (s *Hotel) GetRooms(ctx context.Context, countryCode, citySlug, hotelSlug string, page, limit uint64) (*dto.RoomsResponse, error) {
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

func (s *Hotel) GetRoomsByHotelID(ctx context.Context, countryCode, citySlug, hotelID string, page, limit uint64) (*dto.RoomsResponse, error) {
	if page == 0 {
		page = s.pagination.DefaultPage
	}
	if limit == 0 {
		limit = s.pagination.DefaultPageSize
	}

	resp, err := s.clients.Room.GetRoomsByHotelID(ctx, &hotelv1.GetRoomsByHotelIDRequest{
		HotelId:     hotelID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
		Page:        page,
		Limit:       limit,
	})
	if err != nil {
		return nil, err
	}

	return mapper.RoomsShortByHotelIDResponseFromProto(resp), nil
}

func (s *Hotel) GetRoom(ctx context.Context, countryCode, citySlug, roomID string) (*dto.RoomResponse, error) {
	resp, err := s.clients.Room.GetRoom(ctx, &hotelv1.GetRoomRequest{
		Id:          roomID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
	})
	if err != nil {
		return nil, err
	}

	return mapper.RoomResponseFromProto(resp.Room), nil
}

func (s *Hotel) UpdateRoom(ctx context.Context, countryCode, citySlug, roomID string, req dto.UpdateRoomRequest) (*dto.RoomResponse, error) {
	resp, err := s.clients.Room.UpdateRoom(ctx, &hotelv1.UpdateRoomRequest{
		Id:          roomID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
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

func (s *Hotel) UpdateRoomStatus(ctx context.Context, countryCode, citySlug, roomID string, req dto.UpdateRoomStatusRequest) (*dto.StatusResponse, error) {
	resp, err := s.clients.Room.UpdateRoomStatus(ctx, &hotelv1.UpdateRoomStatusRequest{
		Id:          roomID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
		Status:      hotelv1.RoomStatus(hotelv1.RoomStatus_value[req.Status]),
	})
	if err != nil {
		return nil, err
	}

	return &dto.StatusResponse{Status: resp.Status.String()}, nil
}

func (s *Hotel) DeleteRoom(ctx context.Context, countryCode, citySlug, roomID string) error {
	_, err := s.clients.Room.DeleteRoom(ctx, &hotelv1.DeleteRoomRequest{
		Id:          roomID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
	})
	return err
}

func toInt32(v int64, field string) (int32, error) {
	if v < int32Min || v > int32Max {
		return 0, status.Errorf(codes.InvalidArgument, "%s is out of range", field)
	}
	return int32(v), nil
}
