package handler

import (
	"context"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
)

type authService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.TokenResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.TokenResponse, error)
	RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (*dto.TokenResponse, error)
	GetUsers(ctx context.Context, page, limit uint64) (*dto.UsersResponse, error)
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUser(ctx context.Context, id int64) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id int64, req dto.UpdateUserRequest) (*dto.UserResponse, error)
	UpdateUserActivity(ctx context.Context, id int64, req dto.UpdateUserActivityRequest) (*dto.UpdateUserActivityResponse, error)
	UpdateUserRole(ctx context.Context, id int64, req dto.UpdateUserRoleRequest) (*dto.UpdateUserRoleResponse, error)
	DeleteUser(ctx context.Context, id int64) error
}

type hotelService interface {
	CreateHotel(ctx context.Context, req dto.CreateHotelRequest) (*dto.HotelResponse, error)
	GetHotels(ctx context.Context, countryCode, citySlug, sortBy string, page, limit uint64) (*dto.HotelsResponse, error)
	GetHotelBySlug(ctx context.Context, countryCode, citySlug, hotelSlug string) (*dto.HotelResponse, error)
	GetHotelByID(ctx context.Context, hotelID string) (*dto.HotelResponse, error)
	UpdateHotelByID(ctx context.Context, hotelID string, req dto.UpdateHotelRequest) (*dto.HotelResponse, error)
	UpdateHotelTitleByID(ctx context.Context, hotelID string, req dto.UpdateHotelTitleRequest) (*dto.HotelResponse, error)
	DeleteHotelByID(ctx context.Context, hotelID string) error
	CreateRoomByHotelID(ctx context.Context, hotelID string, req dto.CreateRoomRequest) (*dto.RoomResponse, error)
	GetRooms(ctx context.Context, countryCode, citySlug, hotelSlug string, page, limit uint64) (*dto.RoomsResponse, error)
	GetRoomsByHotelID(ctx context.Context, hotelID string, page, limit uint64) (*dto.RoomsResponse, error)
	GetRoom(ctx context.Context, roomID string) (*dto.RoomResponse, error)
	UpdateRoom(ctx context.Context, roomID string, req dto.UpdateRoomRequest) (*dto.RoomResponse, error)
	UpdateRoomStatus(ctx context.Context, roomID string, req dto.UpdateRoomStatusRequest) (*dto.StatusResponse, error)
	DeleteRoom(ctx context.Context, roomID string) error
}

type bookingService interface {
	CreateBooking(ctx context.Context, hotelID string, req dto.CreateBookingRequest) (*dto.BookingResponse, error)
	GetBookings(ctx context.Context, userID int64, hotelID, status string, page, limit uint64) (*dto.BookingsResponse, error)
	GetBooking(ctx context.Context, bookingID string) (*dto.BookingResponse, error)
	ConfirmBooking(ctx context.Context, bookingID string) (*dto.StatusResponse, error)
	CancelBooking(ctx context.Context, bookingID string) (*dto.StatusResponse, error)
	DeleteBooking(ctx context.Context, bookingID string) error
}

type AuthHandler struct {
	service authService
}

type HotelHandler struct {
	service hotelService
}

type BookingHandler struct {
	service bookingService
}

func NewAuthHandler(svc authService) *AuthHandler {
	return &AuthHandler{service: svc}
}

func NewHotelHandler(svc hotelService) *HotelHandler {
	return &HotelHandler{service: svc}
}

func NewBookingHandler(svc bookingService) *BookingHandler {
	return &BookingHandler{service: svc}
}
