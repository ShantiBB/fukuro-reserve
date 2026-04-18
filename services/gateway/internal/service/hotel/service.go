package hotel

import (
	"context"
	"fmt"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/mapper"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

type Service struct {
	clients    *clients.Clients
	pagination config.PaginationConfig
}

func New(clients *clients.Clients, pagination config.PaginationConfig) *Service {
	return &Service{clients: clients, pagination: pagination}
}

func (s *Service) CreateHotel(ctx context.Context, req dto.CreateHotelRequest) (*dto.HotelResponse, error) {
	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	location := &hotelv1.CreateHotelLocationRequest{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
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

func (s *Service) GetHotels(ctx context.Context, countryCode, citySlug, sortBy string, page, limit uint64) (*dto.HotelsResponse, error) {
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

func (s *Service) GetHotel(ctx context.Context, countryCode, citySlug, hotelSlug string) (*dto.HotelResponse, error) {
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

func (s *Service) UpdateHotel(ctx context.Context, countryCode, citySlug, hotelSlug string, req dto.UpdateHotelRequest) (*dto.HotelResponse, error) {
	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	location := &hotelv1.UpdateHotelLocationRequest{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}

	resp, err := s.clients.Hotel.UpdateHotel(ctx, &hotelv1.UpdateHotelRequest{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		HotelSlug:   hotelSlug,
		Description: description,
		Address:     req.Address,
		Location:    location,
	})
	if err != nil {
		return nil, err
	}

	return mapper.UpdateHotelResponseFromProto(resp.Hotel), nil
}

func (s *Service) UpdateHotelTitle(ctx context.Context, countryCode, citySlug, hotelSlug string, req dto.UpdateHotelTitleRequest) (*dto.HotelResponse, error) {
	resp, err := s.clients.Hotel.UpdateHotelTitle(ctx, &hotelv1.UpdateHotelTitleRequest{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		HotelSlug:   hotelSlug,
		Title:       req.Title,
	})
	if err != nil {
		return nil, err
	}

	return mapper.UpdateHotelTitleResponseFromProto(resp.Hotel), nil
}

func (s *Service) DeleteHotel(ctx context.Context, countryCode, citySlug, hotelSlug string) error {
	_, err := s.clients.Hotel.DeleteHotel(ctx, &hotelv1.DeleteHotelRequest{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		HotelSlug:   hotelSlug,
	})
	return err
}

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
