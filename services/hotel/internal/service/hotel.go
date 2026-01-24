package service

import (
	"context"

	"hotel/internal/repository/models"

	"github.com/gosimple/slug"
)

func (s *Service) CreateHotel(ctx context.Context, h *models.CreateHotel) (*models.Hotel, error) {
	h.Slug = slug.Make(h.Title)

	newHotel, err := s.repo.CreateHotel(ctx, h)
	if err != nil {
		return nil, err
	}

	return newHotel, nil
}

func (s *Service) GetHotels(
	ctx context.Context,
	ref models.HotelRef,
	sort string,
	page uint64,
	limit uint64,
) (*models.HotelList, error) {
	offset := (page - 1) * limit
	hotelList, err := s.repo.GetHotels(ctx, ref, sort, limit, offset)
	if err != nil {
		return nil, err
	}

	return hotelList, nil
}

func (s *Service) GetHotelBySlug(ctx context.Context, ref models.HotelRef) (*models.Hotel, error) {
	h, err := s.repo.GetHotelBySlug(ctx, ref)
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (s *Service) UpdateHotelBySlug(ctx context.Context, ref models.HotelRef, h models.UpdateHotel) error {
	if err := s.repo.UpdateHotelBySlug(ctx, ref, h); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateHotelTitleBySlug(
	ctx context.Context,
	ref models.HotelRef,
	h models.UpdateHotelTitle,
) (models.UpdateHotelTitle, error) {
	h.Slug = slug.Make(h.Title)
	if err := s.repo.UpdateHotelTitleBySlug(ctx, ref, h); err != nil {
		return models.UpdateHotelTitle{}, err
	}

	return h, nil
}

func (s *Service) DeleteHotelBySlug(ctx context.Context, ref models.HotelRef) error {
	if err := s.repo.DeleteHotelBySlug(ctx, ref); err != nil {
		return err
	}

	return nil
}
