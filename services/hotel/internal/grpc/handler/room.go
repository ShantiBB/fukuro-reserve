package handler

import (
	"context"
	"log/slog"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/grpc/utils/helper"
	"hotel/internal/grpc/utils/mapper"
)

func (h *Handler) CreateRoom(
	ctx context.Context,
	req *hotelv1.CreateRoomRequest,
) (*hotelv1.CreateRoomResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	ref := mapper.GetHotelRefRequestToDomain(req)
	room := mapper.CreateRoomRequestToDomain(req)
	created, err := h.svc.CreateRoom(ctx, ref, room)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.CreateRoomResponse{
		Room: mapper.RoomResponseToProto(created),
	}, nil
}

func (h *Handler) GetRooms(
	ctx context.Context,
	req *hotelv1.GetRoomsRequest,
) (*hotelv1.GetRoomsResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	ref := mapper.GetHotelRefRequestToDomain(req)
	roomList, err := h.svc.GetRooms(ctx, ref, req.Page, req.Limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.GetRoomsResponse{
		Rooms:      mapper.RoomsResponseToProto(roomList.Rooms),
		TotalCount: roomList.TotalCount,
		Page:       req.Page,
		Limit:      req.Limit,
	}, nil
}

func (h *Handler) GetRoom(
	ctx context.Context,
	req *hotelv1.GetRoomRequest,
) (*hotelv1.GetRoomResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	roomID, err := helper.ParseRoomID(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	hotel, err := h.svc.GetRoomByID(ctx, roomID)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.GetRoomResponse{
		Room: mapper.RoomResponseToProto(hotel),
	}, nil
}

func (h *Handler) UpdateRoom(
	ctx context.Context,
	req *hotelv1.UpdateRoomRequest,
) (*hotelv1.UpdateRoomResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	roomID, err := helper.ParseRoomID(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	updated, err := mapper.UpdateRoomRequestToDomain(req)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	if err = h.svc.UpdateRoomByID(ctx, roomID, updated); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.UpdateRoomResponse{
		Room: mapper.UpdateRoomResponseToProto(updated),
	}, nil
}

func (h *Handler) UpdateRoomStatus(
	ctx context.Context,
	req *hotelv1.UpdateRoomStatusRequest,
) (*hotelv1.UpdateRoomStatusResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	roomID, err := helper.ParseRoomID(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	room := mapper.UpdateRoomStatusRequestToDomain(req)

	if err = h.svc.UpdateRoomStatusByID(ctx, roomID, room); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.UpdateRoomStatusResponse{
		Status: mapper.UpdateRoomStatusResponseToProto(room),
	}, nil
}

func (h *Handler) DeleteRoom(
	ctx context.Context,
	req *hotelv1.DeleteRoomRequest,
) (*hotelv1.DeleteRoomResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	roomID, err := helper.ParseRoomID(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	if err = h.svc.DeleteRoomByID(ctx, roomID); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.DeleteRoomResponse{
		Message: "success",
	}, nil
}
