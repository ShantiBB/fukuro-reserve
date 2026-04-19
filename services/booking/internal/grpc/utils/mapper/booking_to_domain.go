package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/repository/models"
	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/utils/consts"
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
		CountryCode:         req.CountryCode,
		CitySlug:            req.CitySlug,
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
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		UserID:      req.UserId,
		Status:      BookingStatusToDomain(req.Status),
	}

	if req.HotelId == "" {
		if req.RoomId == "" {
			return bookingRef, nil
		}
	} else {
		hotelID, err := uuid.Parse(req.HotelId)
		if err != nil {
			return models.BookingRef{}, consts.ErrInvalidHotelID
		}
		bookingRef.HotelID = hotelID
	}

	if req.RoomId != "" {
		roomID, err := uuid.Parse(req.RoomId)
		if err != nil {
			return models.BookingRef{}, consts.ErrInvalidBookingRoomID
		}
		bookingRef.RoomID = roomID
	}

	return bookingRef, nil
}

func GetUnavailableRoomsRequestToDomain(req *bookingv1.GetUnavailableRoomsRequest) (models.BookingRef, time.Time, time.Time, error) {
	hotelID, err := uuid.Parse(req.HotelId)
	if err != nil {
		return models.BookingRef{}, time.Time{}, time.Time{}, consts.ErrInvalidHotelID
	}

	return models.BookingRef{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		HotelID:     hotelID,
	}, req.CheckIn.AsTime(), req.CheckOut.AsTime(), nil
}

func UpdateBookingGuestInfoRequestToDomain(req *bookingv1.UpdateBookingGuestInfoRequest) (models.BookingRef, uuid.UUID, *models.UpdateBooking, error) {
	bookingID, err := uuid.Parse(req.Id)
	if err != nil {
		return models.BookingRef{}, uuid.UUID{}, nil, consts.ErrInvalidBookingID
	}

	hotelID, err := uuid.Parse(req.HotelId)
	if err != nil {
		return models.BookingRef{}, uuid.UUID{}, nil, consts.ErrInvalidHotelID
	}

	bookingRef := models.BookingRef{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		HotelID:     hotelID,
	}
	booking := &models.UpdateBooking{
		GuestName:  req.GuestName,
		GuestEmail: req.GuestEmail,
		GuestPhone: req.GuestPhone,
	}

	return bookingRef, bookingID, booking, nil
}

type bookingLocationRefGetter interface {
	GetCountryCode() string
	GetCitySlug() string
}

func BookingLocationRefToDomain[T bookingLocationRefGetter](req T) models.BookingRef {
	return models.BookingRef{
		CountryCode: req.GetCountryCode(),
		CitySlug:    req.GetCitySlug(),
	}
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
