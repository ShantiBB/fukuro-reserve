package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateBooking struct {
	CheckIn             time.Time
	CheckOut            time.Time
	GuestEmail          *string
	GuestPhone          *string
	GuestName           string
	Currency            string
	ExpectedTotalAmount decimal.Decimal
	FinalTotalAmount    decimal.Decimal
	UserID              int64
	HotelID             uuid.UUID
}

type UpdateBooking struct {
	GuestName   *string
	GuestEmail  *string
	GuestPhone  *string
	CheckIn     *time.Time
	CheckOut    *time.Time
	TotalAmount *string
}

type BookingStatusInfo struct {
	Status BookingStatus
}

type Booking struct {
	CreatedAt           time.Time
	CheckIn             time.Time
	UpdatedAt           time.Time
	CheckOut            time.Time
	GuestPhone          *string
	GuestEmail          *string
	Status              BookingStatus
	Currency            string
	GuestName           string
	ExpectedTotalAmount decimal.Decimal
	FinalTotalAmount    decimal.Decimal
	BookingRooms        []*BookingRoomWithLock
	UserID              int64
	ID                  uuid.UUID
	HotelID             uuid.UUID
}

type BookingShort struct {
	CheckIn             time.Time
	CheckOut            time.Time
	GuestPhone          *string
	GuestEmail          *string
	Status              BookingStatus
	GuestName           string
	Currency            string
	ExpectedTotalAmount decimal.Decimal
	FinalTotalAmount    decimal.Decimal
	BookingRooms        []*BookingRoom
	UserID              int64
	ID                  uuid.UUID
	HotelID             uuid.UUID
}

type BookingList struct {
	Bookings   []*BookingShort
	TotalCount uint64
}

type BookingRef struct {
	Status  BookingStatus
	UserID  int64
	HotelID uuid.UUID
}

func (b *CreateBooking) ToRead() *Booking {
	return &Booking{
		UserID:              b.UserID,
		HotelID:             b.HotelID,
		CheckIn:             b.CheckIn,
		CheckOut:            b.CheckOut,
		GuestName:           b.GuestName,
		GuestEmail:          b.GuestEmail,
		GuestPhone:          b.GuestPhone,
		Currency:            b.Currency,
		ExpectedTotalAmount: b.ExpectedTotalAmount,
		FinalTotalAmount:    b.FinalTotalAmount,
	}
}
