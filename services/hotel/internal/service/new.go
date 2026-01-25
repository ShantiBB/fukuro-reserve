package service

import (
	"context"

	"github.com/google/uuid"

	"hotel/internal/repository/models"
)

type HotelRepository interface {
	InsertHotel(ctx context.Context, h *models.CreateHotel) (*models.Hotel, error)
	SelectHotels(
		ctx context.Context, hotelRef models.HotelRef, sortField string, limit, offset uint64,
	) (*models.HotelList, error)
	SelectHotelBySlug(ctx context.Context, ref models.HotelRef) (*models.Hotel, error)
	UpdateHotelBySlug(ctx context.Context, ref models.HotelRef, h models.UpdateHotel) error
	UpdateHotelTitleBySlug(ctx context.Context, ref models.HotelRef, h models.UpdateHotelTitle) error
	DeleteHotelBySlug(ctx context.Context, ref models.HotelRef) error
}

type RoomRepository interface {
	InsertRoom(ctx context.Context, hotelRef models.HotelRef, room *models.CreateRoom) (*models.Room, error)
	SelectRooms(ctx context.Context, hotelRef models.HotelRef, limit, offset uint64) (*models.RoomList, error)
	SelectRoomByID(ctx context.Context, roomID uuid.UUID) (*models.Room, error)
	UpdateRoomByID(ctx context.Context, roomID uuid.UUID, room *models.UpdateRoom) error
	UpdateRoomStatusByID(ctx context.Context, roomID uuid.UUID, room models.UpdateRoomStatus) error
	DeleteRoomByID(ctx context.Context, roomID uuid.UUID) error
}

type Repository interface {
	HotelRepository
	RoomRepository
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo}
}
