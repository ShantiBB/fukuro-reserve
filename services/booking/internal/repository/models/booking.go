package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateBooking struct {
	UserID      int64
	HotelID     string
	CheckIn     string
	CheckOut    string
	GuestName   string
	GuestEmail  *string
	GuestPhone  *string
	Currency    string
	TotalAmount string
}

type UpdateBooking struct {
	GuestName   *string
	GuestEmail  *string
	GuestPhone  *string
	CheckIn     *string
	CheckOut    *string
	TotalAmount *string
}

type BookingStatusInfo struct {
	Status BookingStatus
}

type Booking struct {
	ID          string
	UserID      int64
	HotelID     string
	CheckIn     string
	CheckOut    string
	Status      BookingStatus
	GuestName   string
	GuestEmail  *string
	GuestPhone  *string
	Currency    string
	TotalAmount string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookingShort struct {
	ID          string
	UserID      int64
	HotelID     string
	CheckIn     string
	CheckOut    string
	Status      BookingStatus
	GuestName   string
	GuestEmail  *string
	GuestPhone  *string
	Currency    string
	TotalAmount string
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
		UserID:      b.UserID,
		HotelID:     b.HotelID,
		CheckIn:     b.CheckIn,
		CheckOut:    b.CheckOut,
		GuestName:   b.GuestName,
		GuestEmail:  b.GuestEmail,
		GuestPhone:  b.GuestPhone,
		Currency:    b.Currency,
		TotalAmount: b.TotalAmount,
	}
}
