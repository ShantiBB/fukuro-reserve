package handler

import (
	"context"

	"buf.build/go/protovalidate"
	"github.com/google/uuid"

	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
	"github.com/ShantiBB/fukuro-reserve/services/hotel/internal/repository/models"
)

type HotelService interface {
	CreateHotel(ctx context.Context, h *models.CreateHotel) (*models.Hotel, error)
	GetHotels(ctx context.Context, ref models.HotelRef, sort string, page, limit uint64) (*models.HotelList, error)
	GetHotelByID(ctx context.Context, ref models.HotelRef, id uuid.UUID) (*models.Hotel, error)
	GetHotelBySlug(ctx context.Context, ref models.HotelRef) (*models.Hotel, error)
	UpdateHotelByID(ctx context.Context, ref models.HotelRef, id uuid.UUID, h models.UpdateHotel) error
	UpdateHotelBySlug(ctx context.Context, ref models.HotelRef, h models.UpdateHotel) error
	UpdateHotelTitleByID(
		ctx context.Context, ref models.HotelRef, id uuid.UUID, h models.UpdateHotelTitle,
	) (models.UpdateHotelTitle, error)
	UpdateHotelTitleBySlug(
		ctx context.Context, ref models.HotelRef, h models.UpdateHotelTitle,
	) (models.UpdateHotelTitle, error)
	DeleteHotelByID(ctx context.Context, ref models.HotelRef, id uuid.UUID) error
	DeleteHotelBySlug(ctx context.Context, ref models.HotelRef) error
}

type RoomService interface {
	CreateRoom(ctx context.Context, hotelRef models.HotelRef, room *models.CreateRoom) (*models.Room, error)
	CreateRoomByHotelID(ctx context.Context, hotelRef models.HotelRef, hotelID uuid.UUID, room *models.CreateRoom) (*models.Room, error)
	GetRooms(ctx context.Context, hotelRef models.HotelRef, page, limit uint64) (*models.RoomList, error)
	GetRoomsByHotelID(ctx context.Context, hotelRef models.HotelRef, hotelID uuid.UUID, page, limit uint64) (*models.RoomList, error)
	GetRoomByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID) (*models.Room, error)
	UpdateRoomByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID, room *models.UpdateRoom) error
	UpdateRoomStatusByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID, room models.UpdateRoomStatus) error
	DeleteRoomByID(ctx context.Context, hotelRef models.HotelRef, roomID uuid.UUID) error
}

type Service interface {
	HotelService
	RoomService
}

type Handler struct {
	hotelv1.UnimplementedHotelServiceServer
	hotelv1.UnimplementedRoomServiceServer
	svc       Service
	validator protovalidate.Validator
}

func New(svc Service, validator protovalidate.Validator) *Handler {
	return &Handler{svc: svc, validator: validator}
}
