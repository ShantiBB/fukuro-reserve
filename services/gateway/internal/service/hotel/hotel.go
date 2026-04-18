package hotel

import (
	"context"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/mapper"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

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
