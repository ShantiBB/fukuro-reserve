package handler

import (
	"context"
	"log/slog"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/grpc/utils/helper"
	"hotel/internal/grpc/utils/mapper"
)

func (h *Handler) CreateHotel(
	ctx context.Context,
	req *hotelv1.CreateHotelRequest,
) (*hotelv1.CreateHotelResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	hotel := mapper.CreateHotelRequestToDomain(req)
	created, err := h.svc.CreateHotel(ctx, hotel)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.CreateHotelResponse{
		Hotel: mapper.CreateHotelResponseToProto(created),
	}, nil
}

func (h *Handler) GetHotels(
	ctx context.Context,
	req *hotelv1.GetHotelsRequest,
) (*hotelv1.GetHotelsResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	page, limit, ref := mapper.GetHotelsRequestToDomain(req)
	//TODO Реализовать парс сортировки
	hotelList, err := h.svc.GetHotels(ctx, ref, "title", page, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.GetHotelsResponse{
		Hotels:     mapper.HotelsResponseToProto(hotelList.Hotels),
		TotalCount: hotelList.TotalCount,
		Page:       req.Page,
		Limit:      req.Limit,
	}, nil
}

func (h *Handler) GetHotel(
	ctx context.Context,
	req *hotelv1.GetHotelRequest,
) (*hotelv1.GetHotelResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	ref := mapper.GetHotelRefRequestToDomain(req)
	hotel, err := h.svc.GetHotelBySlug(ctx, ref)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.GetHotelResponse{
		Hotel: mapper.HotelResponseToProto(hotel),
	}, nil
}

func (h *Handler) UpdateHotel(
	ctx context.Context,
	req *hotelv1.UpdateHotelRequest,
) (*hotelv1.UpdateHotelResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	ref := mapper.GetHotelRefRequestToDomain(req)
	hotel := mapper.UpdateHotelRequestToDomain(req)
	if err := h.svc.UpdateHotelBySlug(ctx, ref, hotel); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.UpdateHotelResponse{
		Hotel: mapper.UpdateHotelResponseToProto(hotel),
	}, nil
}

func (h *Handler) UpdateHotelTitle(
	ctx context.Context,
	req *hotelv1.UpdateHotelTitleRequest,
) (*hotelv1.UpdateHotelTitleResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	ref := mapper.GetHotelRefRequestToDomain(req)
	hotel := mapper.UpdateHotelTitleRequestToDomain(req)

	updated, err := h.svc.UpdateHotelTitleBySlug(ctx, ref, hotel)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.UpdateHotelTitleResponse{
		Hotel: mapper.UpdateHotelTitleResponseToProto(updated),
	}, nil
}

func (h *Handler) DeleteHotel(
	ctx context.Context,
	req *hotelv1.DeleteHotelRequest,
) (*hotelv1.DeleteHotelResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	ref := mapper.GetHotelRefRequestToDomain(req)
	if err := h.svc.DeleteHotelBySlug(ctx, ref); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.DeleteHotelResponse{
		Message: "success",
	}, nil
}
