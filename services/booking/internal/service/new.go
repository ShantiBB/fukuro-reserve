package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"booking/internal/repository/models"
)

type BookingTransactionRepository interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
}

type BookingRepository interface {
	CreateBooking(ctx context.Context, tx pgx.Tx, b *models.CreateBooking) (*models.Booking, error)
	GetBookingsByHotelInfo(
		ctx context.Context, tx pgx.Tx, bookingRef models.BookingRef, limit uint64, offset uint64,
	) (*models.BookingList, error)
	GetBookingByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) (*models.Booking, error)
	UpdateBookingGuestInfoByID(ctx context.Context, tx pgx.Tx, id uuid.UUID, b *models.UpdateBooking) error
	UpdateBookingStatusByID(ctx context.Context, tx pgx.Tx, id uuid.UUID, b *models.BookingStatusInfo) error
	DeleteBookingByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
}

type BookingRoomRepository interface {
	CreateBookingRooms(
		ctx context.Context, tx pgx.Tx, bookingID uuid.UUID, rooms []*models.CreateBookingRoom,
	) ([]*models.BookingRoomWithLock, error)
	GetBookingRoomsByBookingIDs(
		ctx context.Context, tx pgx.Tx, bookingIDs []uuid.UUID,
	) ([]*models.BookingRoom, error)
	GetBookingRoomsWithLockByBookingIDs(
		ctx context.Context, tx pgx.Tx, bookingIDs []uuid.UUID,
	) ([]*models.BookingRoomWithLock, error)
	UpdateBookingRoomGuestCountsByID(
		ctx context.Context, tx pgx.Tx, id uuid.UUID, bRoom *models.BookingRoomGuestCounts,
	) error
	DeleteBookingRoomByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
}

type RoomLockRepository interface {
	CreateRoomLocks(ctx context.Context, tx pgx.Tx, locks []*models.CreateRoomLock) ([]*models.RoomLockDetail, error)
	UpdateRoomLockActivityByID(
		ctx context.Context, tx pgx.Tx, id uuid.UUID, roomLock *models.UpdateRoomLockActivity,
	) error
	DeleteRoomLockByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
}

type Repository interface {
	BookingTransactionRepository
	BookingRepository
	BookingRoomRepository
	RoomLockRepository
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo}
}
