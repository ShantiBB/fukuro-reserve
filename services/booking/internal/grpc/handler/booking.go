package handler

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/grpc/utils/helper"
	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/grpc/utils/mapper"
	"github.com/ShantiBB/fukuro-reserve/services/booking/internal/repository/models"
)

func (h *Handler) CreateBooking(
	ctx context.Context,
	req *bookingv1.CreateBookingRequest,
) (*bookingv1.CreateBookingResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	booking, err := mapper.CreateBookingRequestToDomain(req)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	rooms, err := mapper.CreateBookingRoomsToDomain(req.Rooms)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	created, err := h.svc.BookingCreate(ctx, booking, rooms)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &bookingv1.CreateBookingResponse{
		Booking: mapper.BookingToProto(created),
	}, nil
}

func (h *Handler) GetBookings(
	ctx context.Context,
	req *bookingv1.GetBookingsRequest,
) (*bookingv1.GetBookingsResponse, error) {
	if req.Page == 0 || req.Limit == 0 || req.Limit > 100 {
		return nil, status.Error(codes.InvalidArgument, "invalid pagination")
	}

	bookingRef, err := mapper.GetBookingsRequestToDomain(req)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	bookingList, err := h.svc.GetBookings(ctx, bookingRef, req.Page, req.Limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &bookingv1.GetBookingsResponse{
		Bookings:   mapper.BookingListToProto(bookingList.Bookings),
		TotalCount: bookingList.TotalCount,
		Page:       req.Page,
		Limit:      req.Limit,
	}, nil
}

func (h *Handler) GetUnavailableRooms(
	ctx context.Context,
	req *bookingv1.GetUnavailableRoomsRequest,
) (*bookingv1.GetUnavailableRoomsResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	bookingRef, checkIn, checkOut, err := mapper.GetUnavailableRoomsRequestToDomain(req)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	roomIDs, err := h.svc.GetUnavailableRoomIDs(ctx, bookingRef, checkIn, checkOut)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	resp := &bookingv1.GetUnavailableRoomsResponse{
		RoomIds: make([]string, len(roomIDs)),
	}
	for i, roomID := range roomIDs {
		resp.RoomIds[i] = roomID.String()
	}

	return resp, nil
}

func (h *Handler) GetBooking(
	ctx context.Context,
	req *bookingv1.GetBookingRequest,
) (*bookingv1.GetBookingResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	bookingId, err := mapper.GetBookingRequestToDomain(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	bookingRef := mapper.BookingLocationRefToDomain(req)
	booking, err := h.svc.GetBookingById(ctx, bookingRef, bookingId)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
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
		return nil, helper.HandleValidationErr(err)
	}

	bookingId, err := mapper.GetBookingRequestToDomain(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	bookingRef := mapper.BookingLocationRefToDomain(req)
	if err = h.svc.UpdateBookingStatus(ctx, bookingRef, bookingId, models.BookingStatusConfirmed); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
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
		return nil, helper.HandleValidationErr(err)
	}

	bookingId, err := mapper.GetBookingRequestToDomain(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	bookingRef := mapper.BookingLocationRefToDomain(req)
	if err = h.svc.UpdateBookingStatus(ctx, bookingRef, bookingId, models.BookingStatusCancelled); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
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
		return nil, helper.HandleValidationErr(err)
	}

	bookingId, err := mapper.GetBookingRequestToDomain(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	bookingRef := mapper.BookingLocationRefToDomain(req)
	if err = h.svc.DeleteBookingByID(ctx, bookingRef, bookingId); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &bookingv1.DeleteBookingResponse{
		Message: "success",
	}, nil
}
