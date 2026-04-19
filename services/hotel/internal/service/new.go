package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ShantiBB/fukuro-reserve/services/hotel/internal/repository/models"
)

type HotelRepository interface {
	InsertHotel(ctx context.Context, h *models.CreateHotel) (*models.Hotel, error)
	SelectHotels(
		ctx context.Context, hotelRef models.HotelRef, sortField string, limit, offset uint64,
	) (*models.HotelList, error)
	SelectHotelBySlug(ctx context.Context, ref models.HotelRef) (*models.Hotel, error)
	SelectHotelByID(ctx context.Context, ref models.HotelRef, id uuid.UUID) (*models.Hotel, error)
	UpdateHotelBySlug(ctx context.Context, ref models.HotelRef, h models.UpdateHotel) error
	UpdateHotelByID(ctx context.Context, ref models.HotelRef, id uuid.UUID, h models.UpdateHotel) error
	UpdateHotelTitleBySlug(ctx context.Context, ref models.HotelRef, h models.UpdateHotelTitle) error
	UpdateHotelTitleByID(ctx context.Context, ref models.HotelRef, id uuid.UUID, h models.UpdateHotelTitle) error
	DeleteHotelBySlug(ctx context.Context, ref models.HotelRef) error
	DeleteHotelByID(ctx context.Context, ref models.HotelRef, id uuid.UUID) error
}

type RoomRepository interface {
	InsertRoom(ctx context.Context, hotelRef models.HotelRef, room *models.CreateRoom) (*models.Room, error)
	InsertRoomByHotelID(ctx context.Context, hotelRef models.HotelRef, hotelID uuid.UUID, room *models.CreateRoom) (*models.Room, error)
	SelectRooms(ctx context.Context, hotelRef models.HotelRef, limit, offset uint64) (*models.RoomList, error)
	SelectRoomsByHotelID(ctx context.Context, hotelRef models.HotelRef, hotelID uuid.UUID, limit, offset uint64) (*models.RoomList, error)
	SelectRoomByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID) (*models.Room, error)
	UpdateRoomByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID, room *models.UpdateRoom) error
	UpdateRoomStatusByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID, room models.UpdateRoomStatus) error
	DeleteRoomByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID) error
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
