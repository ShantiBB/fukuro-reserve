package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/grpc/dto/response"
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
		return nil, response.DomainError(err)
	}

	return &bookingv1.CreateBookingResponse{
		Booking: mapper.BookingToProto(&created),
	}, nil
}
