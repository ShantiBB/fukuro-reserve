package handler

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/grpc/utils/helper"
	"booking/internal/grpc/utils/mapper"
)

func (h *Handler) CreateBooking(
	ctx context.Context,
	req *bookingv1.CreateBookingRequest,
) (*bookingv1.CreateBookingResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	booking, err := mapper.CreateBookingRequestToDomain(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	rooms, err := mapper.CreateBookingRoomsToDomain(req.Rooms)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	created, err := h.svc.BookingCreate(ctx, *booking, rooms)
	if err != nil {
		slog.Error(err.Error())
		return nil, helper.DomainError(err)
	}

	return &bookingv1.CreateBookingResponse{
		Booking: mapper.BookingToProto(&created),
	}, nil
}

func (h *Handler) GetBookings(
	ctx context.Context,
	req *bookingv1.GetBookingsRequest,
) (*bookingv1.GetBookingsResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bookingRef, err := mapper.GetBookingsRequestToDomain(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bookingList, err := h.svc.GetBookings(ctx, bookingRef, req.Page, req.Limit)
	if err != nil {
		slog.Error(err.Error())
		return nil, helper.DomainError(err)
	}

	return &bookingv1.GetBookingsResponse{
		Bookings:   mapper.BookingListToProto(bookingList.Bookings),
		TotalCount: bookingList.TotalCount,
		Page:       req.Page,
		Limit:      req.Limit,
	}, nil
}
