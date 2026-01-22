package mapper

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
	"booking/pkg/utils/consts"
)

func CreateBookingRequestToDomain(req *bookingv1.CreateBookingRequest) (*models.CreateBooking, error) {
	hotelID, err := uuid.Parse(req.HotelId)
	if err != nil {
		return nil, consts.ErrInvalidHotelID
	}

	expectedTotalAmount, err := decimal.NewFromString(req.ExpectedTotalAmount)
	if err != nil {
		return nil, consts.ErrInvalidExpectedTotalAmountID
	}

	b := &models.CreateBooking{
		UserID:              req.UserId,
		HotelID:             hotelID,
		CheckIn:             req.CheckIn.AsTime(),
		CheckOut:            req.CheckOut.AsTime(),
		GuestName:           req.GuestName,
		GuestEmail:          req.GuestEmail,
		GuestPhone:          req.GuestPhone,
		Currency:            req.Currency,
		ExpectedTotalAmount: expectedTotalAmount,
	}

	return b, nil
}

func GetBookingsRequestToDomain(req *bookingv1.GetBookingsRequest) (models.BookingRef, error) {
	bookingRef := models.BookingRef{
		UserID: req.UserId,
		Status: BookingStatusToDomain(req.Status),
	}

	hotelID, err := uuid.Parse(req.HotelId)
	if err != nil {
		return models.BookingRef{}, consts.ErrInvalidHotelID
	}
	bookingRef.HotelID = hotelID

	return bookingRef, nil
}

func GetBookingRequestToDomain(idStr string) (uuid.UUID, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.UUID{}, consts.ErrInvalidBookingID
	}

	return id, nil
}

func BookingStatusToDomain(status bookingv1.BookingStatus) models.BookingStatus {
	var s models.BookingStatus
	switch status {
	case bookingv1.BookingStatus_BOOKING_STATUS_PENDING:
		s = models.BookingStatusPending
	case bookingv1.BookingStatus_BOOKING_STATUS_CONFIRMED:
		s = models.BookingStatusConfirmed
	case bookingv1.BookingStatus_BOOKING_STATUS_CANCELLED:
		s = models.BookingStatusCancelled
	default:
		s = models.BookingStatusUnspecified
	}
	return s
}
