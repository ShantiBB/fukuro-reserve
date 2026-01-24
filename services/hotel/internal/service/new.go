package service

import (
	"context"

	"github.com/google/uuid"

	"hotel/internal/repository/models"
)

type HotelRepository interface {
	CreateHotel(ctx context.Context, h *models.CreateHotel) (*models.Hotel, error)
	GetHotels(
		ctx context.Context, hotelRef models.HotelRef, sortField string, limit, offset uint64,
	) (*models.HotelList, error)
	GetHotelBySlug(ctx context.Context, ref models.HotelRef) (*models.Hotel, error)
	UpdateHotelBySlug(ctx context.Context, ref models.HotelRef, h models.UpdateHotel) error
	UpdateHotelTitleBySlug(ctx context.Context, ref models.HotelRef, h models.UpdateHotelTitle) error
	DeleteHotelBySlug(ctx context.Context, ref models.HotelRef) error
}

type RoomRepository interface {
	RoomCreate(ctx context.Context, hotel models.HotelRef, room models.RoomCreate) (models.Room, error)
	RoomGetByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID) (models.Room, error)
	RoomGetAll(ctx context.Context, hotel models.HotelRef, limit, offset uint64) (models.RoomList, error)
	RoomUpdateByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID, room models.RoomUpdate) error
	RoomStatusUpdateByID(
		ctx context.Context, hotel models.HotelRef, roomID uuid.UUID, room models.RoomStatusUpdate,
	) error
	RoomDeleteByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID) error
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
