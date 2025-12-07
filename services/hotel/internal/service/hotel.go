package service

import (
	"context"

	"hotel/internal/repository/postgres/models"
)

type HotelRepository interface {
	HotelCreate(ctx context.Context, h models.HotelCreate) (models.Hotel, error)
	HotelGetByIDOrName(ctx context.Context, field any) (models.Hotel, error)
	HotelGetAll(ctx context.Context, limit, offset uint64) ([]models.HotelShort, error)
	HotelUpdateByID(ctx context.Context, id int64, h models.HotelUpdate) error
	HotelDeleteByID(ctx context.Context, id int64) error
}

func (s *Service) HotelCreate(ctx context.Context, h models.HotelCreate) (models.Hotel, error) {
	newHotel, err := s.repo.HotelCreate(ctx, h)
	if err != nil {
		return models.Hotel{}, err
	}

	return newHotel, nil
}

func (s *Service) HotelGetByIDOrName(ctx context.Context, field any) (models.Hotel, error) {
	h, err := s.repo.HotelGetByIDOrName(ctx, field)
	if err != nil {
		return models.Hotel{}, err
	}

	return h, nil
}

func (s *Service) HotelGetAll(ctx context.Context, page, limit uint64) ([]models.HotelShort, error) {
	offset := (page - 1) * limit
	hotels, err := s.repo.HotelGetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func (s *Service) HotelUpdateByID(ctx context.Context, id int64, h models.HotelUpdate) error {
	if err := s.repo.HotelUpdateByID(ctx, id, h); err != nil {
		return err
	}

	return nil
}

func (s *Service) HotelDeleteByID(ctx context.Context, id int64) error {
	if err := s.repo.HotelDeleteByID(ctx, id); err != nil {
		return err
	}

	return nil
}
