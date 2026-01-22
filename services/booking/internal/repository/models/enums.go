package models

type BookingStatus string

const (
	BookingStatusPending     BookingStatus = "BOOKING_STATUS_PENDING"
	BookingStatusConfirmed   BookingStatus = "BOOKING_STATUS_CONFIRMED"
	BookingStatusCancelled   BookingStatus = "BOOKING_STATUS_CANCELLED"
	BookingStatusUnspecified BookingStatus = "BOOKING_STATUS_UNSPECIFIED"
)
