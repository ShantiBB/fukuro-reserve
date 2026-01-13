package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateBooking struct {
	UserID              int64
	HotelID             uuid.UUID
	CheckIn             time.Time
	CheckOut            time.Time
	GuestName           string
	GuestEmail          *string
	GuestPhone          *string
	Currency            string
	ExpectedTotalAmount decimal.Decimal
	FinalTotalAmount    decimal.Decimal
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
	ID                  uuid.UUID
	UserID              int64
	HotelID             uuid.UUID
	BookingRooms        []BookingRoom
	CheckIn             time.Time
	CheckOut            time.Time
	Status              BookingStatus
	GuestName           string
	GuestEmail          *string
	GuestPhone          *string
	Currency            string
	ExpectedTotalAmount decimal.Decimal
	FinalTotalAmount    decimal.Decimal
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type BookingShort struct {
	ID                  string
	UserID              int64
	HotelID             uuid.UUID
	CheckIn             time.Time
	CheckOut            time.Time
	Status              BookingStatus
	GuestName           string
	GuestEmail          *string
	GuestPhone          *string
	Currency            string
	ExpectedTotalAmount decimal.Decimal
	FinalTotalAmount    decimal.Decimal
}

type BookingList struct {
	Bookings   []BookingShort
	TotalCount uint64
}

type BookingRef struct {
	UserID  int64
	HotelID uuid.UUID
	Status  BookingStatus
}

func (b *CreateBooking) ToRead() Booking {
	return Booking{
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
