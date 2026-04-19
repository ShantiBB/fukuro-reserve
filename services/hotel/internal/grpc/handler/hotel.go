package handler

import (
	"context"
	"log/slog"

	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
	"github.com/ShantiBB/fukuro-reserve/services/hotel/internal/grpc/utils/helper"
	"github.com/ShantiBB/fukuro-reserve/services/hotel/internal/grpc/utils/mapper"
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

func (h *Handler) GetHotels(ctx context.Context, req *hotelv1.GetHotelsRequest) (*hotelv1.GetHotelsResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	ref, sort, page, limit := mapper.GetHotelsRequestToDomain(req)
	hotelList, err := h.svc.GetHotels(ctx, ref, sort, page, limit)
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

func (h *Handler) GetHotelByID(
	ctx context.Context,
	req *hotelv1.GetHotelByIDRequest,
) (*hotelv1.GetHotelByIDResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	hotelID, err := mapper.GetHotelByIDRequestToDomain(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	hotel, err := h.svc.GetHotelByID(ctx, hotelID)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.GetHotelByIDResponse{
		Hotel: mapper.HotelResponseToProto(hotel),
	}, nil
}

func (h *Handler) GetHotel(ctx context.Context, req *hotelv1.GetHotelRequest) (*hotelv1.GetHotelResponse, error) {
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

func (h *Handler) UpdateHotelByID(
	ctx context.Context,
	req *hotelv1.UpdateHotelByIDRequest,
) (*hotelv1.UpdateHotelByIDResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	hotelID, err := helper.ParseHotelID(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	hotel := mapper.UpdateHotelByIDRequestToDomain(req)
	if err := h.svc.UpdateHotelByID(ctx, hotelID, hotel); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.UpdateHotelByIDResponse{
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

func (h *Handler) UpdateHotelTitleByID(
	ctx context.Context,
	req *hotelv1.UpdateHotelTitleByIDRequest,
) (*hotelv1.UpdateHotelTitleByIDResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	hotelID, err := helper.ParseHotelID(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	hotel := mapper.UpdateHotelTitleByIDRequestToDomain(req)
	updated, err := h.svc.UpdateHotelTitleByID(ctx, hotelID, hotel)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.UpdateHotelTitleByIDResponse{
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

func (h *Handler) DeleteHotelByID(
	ctx context.Context,
	req *hotelv1.DeleteHotelByIDRequest,
) (*hotelv1.DeleteHotelByIDResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	hotelID, err := helper.ParseHotelID(req.Id)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	if err := h.svc.DeleteHotelByID(ctx, hotelID); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &hotelv1.DeleteHotelByIDResponse{
		Message: "success",
	}, nil
}
