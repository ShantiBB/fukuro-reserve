package service

import (
	"context"
	"hotel/internal/repository/models"

	"github.com/gosimple/slug"
)

type HotelRepository interface {
	HotelCreate(ctx context.Context, h models.HotelCreate) (models.Hotel, error)
	HotelGetBySlug(ctx context.Context, h models.Hotel) (models.Hotel, error)
	HotelGetAll(ctx context.Context, filter models.HotelFilter) (models.HotelList, error)
	HotelUpdateBySlug(ctx context.Context, hotelSlug string, h models.HotelUpdate) error
	HotelDeleteBySlug(ctx context.Context, countryCode, citySlug, slug string) error
}

func (s *Service) HotelCreate(
	ctx context.Context,
	countryCode string,
	citySlug string,
	h models.HotelCreate,
) (models.Hotel, error) {
	h.CountryCode = countryCode
	h.CitySlug = citySlug
	h.Slug = slug.Make(h.Title)

	newHotel, err := s.repo.HotelCreate(ctx, h)
	if err != nil {
		return models.Hotel{}, err
	}

	return newHotel, nil
}

func (s *Service) HotelGetAll(
	ctx context.Context,
	countryCode string,
	citySlug string,
	sortField string,
	page uint64,
	limit uint64,
) (models.HotelList, error) {
	filter := models.HotelFilter{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		SortField:   sortField,
		Limit:       limit,
		Offset:      (page - 1) * limit,
	}

	hotelList, err := s.repo.HotelGetAll(ctx, filter)
	if err != nil {
		return models.HotelList{}, err
	}

	return hotelList, nil
}

func (s *Service) HotelGetBySlug(
	ctx context.Context,
	countryCode string,
	citySlug string,
	hotelSlug string,
) (models.Hotel, error) {
	fields := models.Hotel{
		Slug:        hotelSlug,
		CountryCode: countryCode,
		CitySlug:    citySlug,
	}

	h, err := s.repo.HotelGetBySlug(ctx, fields)
	if err != nil {
		return models.Hotel{}, err
	}

	return h, nil
}

func (s *Service) HotelUpdateBySlug(
	ctx context.Context,
	countryCode string,
	citySlug string,
	hotelSlug string,
	h models.HotelUpdate,
) error {
	h.CountryCode = countryCode
	h.CitySlug = citySlug
	h.Slug = slug.Make(h.Title)

	if err := s.repo.HotelUpdateBySlug(ctx, hotelSlug, h); err != nil {
		return err
	}

	return nil
}

func (s *Service) HotelDeleteBySlug(ctx context.Context, countryCode, citySlug, hotelSlug string) error {
	if err := s.repo.HotelDeleteBySlug(ctx, countryCode, citySlug, hotelSlug); err != nil {
		return err
	}

	return nil
}
