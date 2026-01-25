package service

import (
	"context"

	"hotel/internal/repository/models"

	"github.com/google/uuid"
)

func (s *Service) CreateRoom(ctx context.Context, hotel models.HotelRef, room *models.CreateRoom) (
	*models.Room, error,
) {
	newRoom, err := s.repo.InsertRoom(ctx, hotel, room)
	if err != nil {
		return nil, err
	}

	return newRoom, nil
}

func (s *Service) GetRooms(ctx context.Context, hotel models.HotelRef, page, limit uint64) (*models.RoomList, error) {
	offset := (page - 1) * limit
	roomList, err := s.repo.SelectRooms(ctx, hotel, limit, offset)
	if err != nil {
		return nil, err
	}

	return roomList, nil
}

func (s *Service) GetRoomByID(ctx context.Context, roomID uuid.UUID) (*models.Room, error) {
	room, err := s.repo.SelectRoomByID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *Service) UpdateRoomByID(ctx context.Context, roomID uuid.UUID, room *models.UpdateRoom) error {
	if err := s.repo.UpdateRoomByID(ctx, roomID, room); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateRoomStatusByID(ctx context.Context, roomID uuid.UUID, room models.UpdateRoomStatus) error {
	if err := s.repo.UpdateRoomStatusByID(ctx, roomID, room); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteRoomByID(ctx context.Context, roomID uuid.UUID) error {
	if err := s.repo.DeleteRoomByID(ctx, roomID); err != nil {
		return err
	}

	return nil
}
