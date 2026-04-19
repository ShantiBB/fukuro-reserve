package service

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/mapper"
)

func (s *Booking) CreateBooking(ctx context.Context, countryCode, citySlug, hotelID string, req dto.CreateBookingRequest) (*dto.BookingResponse, error) {
	var guestEmail *string
	if req.GuestEmail != "" {
		guestEmail = &req.GuestEmail
	}

	var guestPhone *string
	if req.GuestPhone != "" {
		guestPhone = &req.GuestPhone
	}

	rooms := make([]*bookingv1.CreateBookingRoomRequest, len(req.Rooms))
	for i, room := range req.Rooms {
		rooms[i] = &bookingv1.CreateBookingRoomRequest{
			RoomId:        room.RoomId,
			Adults:        room.Adults,
			Children:      room.Children,
			PricePerNight: room.PricePerNight,
		}
	}

	resp, err := s.clients.Booking.CreateBooking(ctx, &bookingv1.CreateBookingRequest{
		UserId:              req.UserId,
		HotelId:             hotelID,
		CountryCode:         countryCode,
		CitySlug:            citySlug,
		CheckIn:             timestamppb.New(req.CheckIn),
		CheckOut:            timestamppb.New(req.CheckOut),
		GuestName:           req.GuestName,
		GuestEmail:          guestEmail,
		GuestPhone:          guestPhone,
		Currency:            req.Currency,
		ExpectedTotalAmount: req.ExpectedTotalAmount,
		Rooms:               rooms,
	})
	if err != nil {
		return nil, err
	}

	return mapper.BookingResponseFromProto(resp.Booking), nil
}

func (s *Booking) GetBookings(ctx context.Context, countryCode, citySlug string, userID int64, hotelID, status string, page, limit uint64) (*dto.BookingsResponse, error) {
	if page == 0 {
		page = s.pagination.DefaultPage
	}
	if limit == 0 {
		limit = s.pagination.DefaultPageSize
	}

	req := &bookingv1.GetBookingsRequest{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		Page:        page,
		Limit:       limit,
	}
	if userID != 0 {
		req.UserId = userID
	}
	if hotelID != "" {
		req.HotelId = hotelID
	}
	if status != "" {
		req.Status = bookingv1.BookingStatus(bookingv1.BookingStatus_value[status])
	}

	resp, err := s.clients.Booking.GetBookings(ctx, req)
	if err != nil {
		return nil, err
	}

	return mapper.BookingsShortResponseFromProto(resp), nil
}

func (s *Booking) GetBooking(ctx context.Context, countryCode, citySlug, bookingID string) (*dto.BookingResponse, error) {
	resp, err := s.clients.Booking.GetBooking(ctx, &bookingv1.GetBookingRequest{
		Id:          bookingID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
	})
	if err != nil {
		return nil, err
	}

	return mapper.BookingResponseFromProto(resp.Booking), nil
}

func (s *Booking) ConfirmBooking(ctx context.Context, countryCode, citySlug, bookingID string) (*dto.StatusResponse, error) {
	resp, err := s.clients.Booking.ConfirmBookingStatus(ctx, &bookingv1.ConfirmBookingStatusRequest{
		Id:          bookingID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
	})
	if err != nil {
		return nil, err
	}

	return &dto.StatusResponse{Status: resp.Status.String()}, nil
}

func (s *Booking) CancelBooking(ctx context.Context, countryCode, citySlug, bookingID string) (*dto.StatusResponse, error) {
	resp, err := s.clients.Booking.CancelBookingStatus(ctx, &bookingv1.CancelBookingStatusRequest{
		Id:          bookingID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
	})
	if err != nil {
		return nil, err
	}

	return &dto.StatusResponse{Status: resp.Status.String()}, nil
}

func (s *Booking) DeleteBooking(ctx context.Context, countryCode, citySlug, bookingID string) error {
	_, err := s.clients.Booking.DeleteBooking(ctx, &bookingv1.DeleteBookingRequest{
		Id:          bookingID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
	})
	return err
}
