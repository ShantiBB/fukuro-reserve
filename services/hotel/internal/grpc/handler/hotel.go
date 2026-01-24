package handler

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/grpc/utils/helper"
	"hotel/internal/grpc/utils/mapper"
	"hotel/internal/repository/models"
)

func (h *Handler) CreateHotel(
	ctx context.Context,
	req *hotelv1.CreateHotelRequest,
) (*hotelv1.CreateHotelResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	hotel := mapper.CreateHotelRequestToDomain(req)

	created, err := h.svc.CreateHotel(ctx, hotel)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &hotelv1.CreateHotelResponse{
		Hotel: mapper.HotelResponseToProto(created),
	}, nil
}

func (h *Handler) GetHotels(
	ctx context.Context,
	req *hotelv1.GetHotelsRequest,
) (*hotelv1.GetHotelsResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	page, limit, ref := mapper.GetHotelsRequestToDomain(req)

	bookingList, err := h.svc.GetHotels(ctx, ref, "title", page, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &hotelv1.GetHotelsResponse{
		Hotels:     mapper.HotelsResponseToProto(bookingList.Hotels),
		TotalCount: bookingList.TotalCount,
		Page:       req.Page,
		Limit:      req.Limit,
	}, nil
}

func (h *Handler) GetHotel(
	ctx context.Context,
	req *hotelv1.GetHotelRequest,
) (*hotelv1.GetHotelResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ref := models.HotelRef{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		HotelSlug:   req.Slug,
	}
	hotel, err := h.svc.GetHotelBySlug(ctx, ref)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
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
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ref := models.HotelRef{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		HotelSlug:   req.Slug,
	}
	hotel := mapper.UpdateHotelRequestToDomain(req)

	if err := h.svc.UpdateHotelBySlug(ctx, ref, hotel); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &hotelv1.UpdateHotelResponse{
		Hotel: mapper.HotelUpdateResponseToProto(hotel),
	}, nil
}

func (h *Handler) DeleteHotel(
	ctx context.Context,
	req *hotelv1.DeleteHotelRequest,
) (*hotelv1.DeleteHotelResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ref := models.HotelRef{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		HotelSlug:   req.Slug,
	}
	if err := h.svc.DeleteHotelBySlug(ctx, ref); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &hotelv1.DeleteHotelResponse{
		Message: "success",
	}, nil
}
