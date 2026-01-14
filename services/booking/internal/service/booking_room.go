package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"booking/internal/repository/models"
)

func (s *Service) GetBookingRooms(ctx context.Context, bookingID uuid.UUID) ([]models.BookingRoomFullInfo, error) {
	allRooms, err := s.repo.GetBookingRoomsFullInfoByBookingIDs(ctx, nil, bookingID)
	if err != nil {
		return []models.BookingRoomFullInfo{}, fmt.Errorf("get booking rooms: %w", err)
	}

	return allRooms, nil
}
