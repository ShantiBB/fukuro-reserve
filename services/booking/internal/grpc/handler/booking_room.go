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

func (h *Handler) GetBookingRooms(
	ctx context.Context,
	req *bookingv1.GetBookingRoomsRequest,
) (*bookingv1.GetBookingRoomsResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bookingId, err := mapper.GetBookingRoomsRequestToDomain(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bookingRooms, err := h.svc.GetBookingRooms(ctx, bookingId)
	if err != nil {
		slog.Error(err.Error())
		return nil, helper.DomainError(err)
	}

	return &bookingv1.GetBookingRoomsResponse{
		BookingRooms: mapper.BookingRoomsFullInfoToProto(bookingRooms),
	}, nil
}
