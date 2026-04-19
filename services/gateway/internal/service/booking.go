package service

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/mapper"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
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

func (s *Booking) QuoteBooking(ctx context.Context, countryCode, citySlug, hotelID string, req dto.QuoteBookingRequest) (*dto.QuoteBookingResponse, error) {
	rooms := make([]*bookingv1.CreateBookingRoomRequest, len(req.Rooms))
	for i, room := range req.Rooms {
		rooms[i] = &bookingv1.CreateBookingRoomRequest{
			RoomId:        room.RoomId,
			Adults:        room.Adults,
			Children:      room.Children,
			PricePerNight: room.PricePerNight,
		}
	}

	resp, err := s.clients.Booking.QuoteBooking(ctx, &bookingv1.QuoteBookingRequest{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		HotelId:     hotelID,
		CheckIn:     timestamppb.New(req.CheckIn),
		CheckOut:    timestamppb.New(req.CheckOut),
		Currency:    req.Currency,
		Rooms:       rooms,
	})
	if err != nil {
		return nil, err
	}

	return mapper.QuoteBookingResponseFromProto(resp), nil
}

func (s *Booking) GetBookings(ctx context.Context, countryCode, citySlug string, userID int64, hotelID, status string, page, limit uint64) (*dto.BookingsResponse, error) {
	return s.getBookings(ctx, countryCode, citySlug, userID, hotelID, "", status, page, limit)
}

func (s *Booking) GetRoomBookings(ctx context.Context, countryCode, citySlug, hotelID, roomID, status string, page, limit uint64) (*dto.BookingsResponse, error) {
	return s.getBookings(ctx, countryCode, citySlug, 0, hotelID, roomID, status, page, limit)
}

func (s *Booking) getBookings(ctx context.Context, countryCode, citySlug string, userID int64, hotelID, roomID, status string, page, limit uint64) (*dto.BookingsResponse, error) {
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
	if roomID != "" {
		req.RoomId = roomID
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

func (s *Booking) GetAvailability(ctx context.Context, countryCode, citySlug, hotelID string, checkIn, checkOut time.Time, page, limit uint64) (*dto.AvailabilityResponse, error) {
	if page == 0 {
		page = s.pagination.DefaultPage
	}
	if limit == 0 {
		limit = s.pagination.DefaultPageSize
	}

	roomsResp, err := s.clients.Room.GetRoomsByHotelID(ctx, &hotelv1.GetRoomsByHotelIDRequest{
		HotelId:     hotelID,
		CountryCode: countryCode,
		CitySlug:    citySlug,
		Page:        page,
		Limit:       limit,
	})
	if err != nil {
		return nil, err
	}

	unavailableResp, err := s.clients.Booking.GetUnavailableRooms(ctx, &bookingv1.GetUnavailableRoomsRequest{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		HotelId:     hotelID,
		CheckIn:     timestamppb.New(checkIn),
		CheckOut:    timestamppb.New(checkOut),
	})
	if err != nil {
		return nil, err
	}

	unavailable := make(map[string]struct{}, len(unavailableResp.GetRoomIds()))
	for _, roomID := range unavailableResp.GetRoomIds() {
		unavailable[roomID] = struct{}{}
	}

	rooms := mapper.RoomsShortByHotelIDResponseFromProto(roomsResp)
	if rooms == nil {
		return &dto.AvailabilityResponse{CheckIn: checkIn, CheckOut: checkOut}, nil
	}

	availableRooms := make([]*dto.RoomShortResponse, 0, len(rooms.Rooms))
	for _, room := range rooms.Rooms {
		if room == nil {
			continue
		}
		if _, ok := unavailable[room.Id]; ok {
			continue
		}
		availableRooms = append(availableRooms, room)
	}

	return &dto.AvailabilityResponse{
		CheckIn:  checkIn,
		CheckOut: checkOut,
		Rooms:    availableRooms,
	}, nil
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
