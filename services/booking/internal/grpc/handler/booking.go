package handler

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/grpc/utils/helper"
	"booking/internal/grpc/utils/mapper"
	"booking/internal/repository/models"
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

	created, err := h.svc.BookingCreate(ctx, booking, rooms)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &bookingv1.CreateBookingResponse{
		Booking: mapper.BookingToProto(created),
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
		return nil, errInvalidHotelID
	}

	bookingList, err := h.svc.GetBookings(ctx, bookingRef, req.Page, req.Limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &bookingv1.GetBookingsResponse{
		Bookings:   mapper.BookingListToProto(bookingList.Bookings),
		TotalCount: bookingList.TotalCount,
		Page:       req.Page,
		Limit:      req.Limit,
	}, nil
}

func (h *Handler) GetBooking(
	ctx context.Context,
	req *bookingv1.GetBookingRequest,
) (*bookingv1.GetBookingResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bookingId, err := mapper.GetBookingRequestToDomain(req.Id)
	if err != nil {
		return nil, errInvalidBookingID
	}

	booking, err := h.svc.GetBookingById(ctx, bookingId)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &bookingv1.GetBookingResponse{
		Booking: mapper.BookingToProto(booking),
	}, nil
}

func (h *Handler) ConfirmBookingStatus(
	ctx context.Context,
	req *bookingv1.ConfirmBookingStatusRequest,
) (*bookingv1.ConfirmBookingStatusResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bookingId, err := mapper.GetBookingRequestToDomain(req.Id)
	if err != nil {
		return nil, errInvalidBookingID
	}

	if err = h.svc.UpdateBookingStatus(ctx, bookingId, models.BookingStatusConfirmed); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &bookingv1.ConfirmBookingStatusResponse{
		Status: mapper.BookingStatusToProto(models.BookingStatusConfirmed),
	}, nil
}

func (h *Handler) CancelBookingStatus(
	ctx context.Context,
	req *bookingv1.CancelBookingStatusRequest,
) (*bookingv1.CancelBookingStatusResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bookingId, err := mapper.GetBookingRequestToDomain(req.Id)
	if err != nil {
		return nil, errInvalidBookingID
	}

	if err = h.svc.UpdateBookingStatus(ctx, bookingId, models.BookingStatusCancelled); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &bookingv1.CancelBookingStatusResponse{
		Status: mapper.BookingStatusToProto(models.BookingStatusCancelled),
	}, nil
}

func (h *Handler) DeleteBooking(
	ctx context.Context,
	req *bookingv1.DeleteBookingRequest,
) (*bookingv1.DeleteBookingResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bookingId, err := mapper.GetBookingRequestToDomain(req.Id)
	if err != nil {
		return nil, errInvalidBookingID
	}

	if err = h.svc.DeleteBookingByID(ctx, bookingId); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &bookingv1.DeleteBookingResponse{
		Message: "success",
	}, nil
}
