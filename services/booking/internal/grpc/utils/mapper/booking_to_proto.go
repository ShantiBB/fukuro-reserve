package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
)

func BookingToProto(b *models.Booking) *bookingv1.Booking {
	p := &bookingv1.Booking{
		Id:                  b.ID.String(),
		UserId:              b.UserID,
		HotelId:             b.HotelID.String(),
		CheckIn:             timestamppb.New(b.CheckIn),
		CheckOut:            timestamppb.New(b.CheckOut),
		Status:              BookingStatusToProto(b.Status),
		GuestName:           b.GuestName,
		GuestEmail:          b.GuestEmail,
		GuestPhone:          b.GuestPhone,
		Currency:            b.Currency,
		ExpectedTotalAmount: b.ExpectedTotalAmount.String(),
		FinalTotalAmount:    b.FinalTotalAmount.String(),
		CreatedAt:           timestamppb.New(b.CreatedAt),
		UpdatedAt:           timestamppb.New(b.UpdatedAt),
		BookingRooms:        BookingRoomsWithLockToProto(b.BookingRooms),
	}

	return p
}

func BookingListToProto(bookings []*models.BookingShort) []*bookingv1.BookingShort {
	result := make([]*bookingv1.BookingShort, len(bookings))
	for i, b := range bookings {
		result[i] = BookingShortToProto(b)
	}
	return result
}

func BookingShortToProto(b *models.BookingShort) *bookingv1.BookingShort {
	return &bookingv1.BookingShort{
		Id:                  b.ID.String(),
		UserId:              b.UserID,
		HotelId:             b.HotelID.String(),
		CheckIn:             timestamppb.New(b.CheckIn),
		CheckOut:            timestamppb.New(b.CheckOut),
		Status:              BookingStatusToProto(b.Status),
		GuestName:           b.GuestName,
		GuestEmail:          b.GuestEmail,
		GuestPhone:          b.GuestPhone,
		Currency:            b.Currency,
		ExpectedTotalAmount: b.ExpectedTotalAmount.String(),
		FinalTotalAmount:    b.FinalTotalAmount.String(),
		BookingRooms:        BookingRoomsToProto(b.BookingRooms),
	}
}

func BookingStatusToProto(s models.BookingStatus) bookingv1.BookingStatus {
	switch s {
	case models.BookingStatusPending:
		return bookingv1.BookingStatus_BOOKING_STATUS_PENDING
	case models.BookingStatusConfirmed:
		return bookingv1.BookingStatus_BOOKING_STATUS_CONFIRMED
	case models.BookingStatusCancelled:
		return bookingv1.BookingStatus_BOOKING_STATUS_CANCELLED
	default:
		return bookingv1.BookingStatus_BOOKING_STATUS_UNSPECIFIED
	}
}
