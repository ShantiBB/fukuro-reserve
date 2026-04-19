package service

import (
	"context"

	"github.com/ShantiBB/fukuro-reserve/services/hotel/internal/repository/models"

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

func (s *Service) CreateRoomByHotelID(
	ctx context.Context,
	hotelRef models.HotelRef,
	hotelID uuid.UUID,
	room *models.CreateRoom,
) (*models.Room, error) {
	if _, err := s.repo.SelectHotelByID(ctx, hotelRef, hotelID); err != nil {
		return nil, err
	}

	newRoom, err := s.repo.InsertRoomByHotelID(ctx, hotelRef, hotelID, room)
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

func (s *Service) GetRoomsByHotelID(
	ctx context.Context,
	hotelRef models.HotelRef,
	hotelID uuid.UUID,
	page,
	limit uint64,
) (*models.RoomList, error) {
	if _, err := s.repo.SelectHotelByID(ctx, hotelRef, hotelID); err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	roomList, err := s.repo.SelectRoomsByHotelID(ctx, hotelRef, hotelID, limit, offset)
	if err != nil {
		return nil, err
	}

	return roomList, nil
}

func (s *Service) GetRoomByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID) (*models.Room, error) {
	room, err := s.repo.SelectRoomByID(ctx, hotelRef, roomID)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *Service) UpdateRoomByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID, room *models.UpdateRoom) error {
	if err := s.repo.UpdateRoomByID(ctx, hotelRef, roomID, room); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateRoomStatusByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID, room models.UpdateRoomStatus) error {
	if err := s.repo.UpdateRoomStatusByID(ctx, hotelRef, roomID, room); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteRoomByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID) error {
	if err := s.repo.DeleteRoomByID(ctx, hotelRef, roomID); err != nil {
		return err
	}

	return nil
}
