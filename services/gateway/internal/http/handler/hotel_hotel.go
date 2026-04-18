package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/query"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
)

// CreateHotel godoc
// @Summary Create a new hotel
// @Tags hotels
// @Accept json
// @Produce json
// @Param request body dto.CreateHotelRequest true "Create hotel request"
// @Success 201 {object} dto.HotelResponse
// @Router /api/v1/hotels [post]
func (h *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateHotelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.CreateHotel(r.Context(), req)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusCreated, resp)
}

// GetHotels godoc
// @Summary Get hotels
// @Tags hotels
// @Produce json
// @Param country_code query string true "Country code"
// @Param city_slug query string true "City slug"
// @Param sort_by query string false "Sort field"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {object} dto.HotelsResponse
// @Router /api/v1/hotels [get]
func (h *HotelHandler) GetHotels(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	if countryCode == "" {
		countryCode = r.URL.Query().Get("country_code")
	}

	citySlug := r.URL.Query().Get("citySlug")
	if citySlug == "" {
		citySlug = r.URL.Query().Get("city_slug")
	}

	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		sortBy = r.URL.Query().Get("sort_by")
	}
	if sortBy == "" {
		sortBy = "title"
	}

	resp, err := h.service.GetHotels(
		r.Context(),
		countryCode,
		citySlug,
		sortBy,
		query.ParseUint64(r.URL.Query().Get("page")),
		query.ParseUint64(r.URL.Query().Get("limit")),
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// GetHotel godoc
// @Summary Get hotel by slug
// @Tags hotels
// @Produce json
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Success 200 {object} dto.HotelResponse
// @Router /api/v1/hotels/{countryCode}/{citySlug}/{hotelSlug} [get]
func (h *HotelHandler) GetHotel(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetHotel(
		r.Context(),
		chi.URLParam(r, "countryCode"),
		chi.URLParam(r, "citySlug"),
		chi.URLParam(r, "hotelSlug"),
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// UpdateHotel godoc
// @Summary Update hotel
// @Tags hotels
// @Accept json
// @Produce json
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Param request body dto.UpdateHotelRequest true "Update hotel request"
// @Success 200 {object} dto.HotelResponse
// @Router /api/v1/hotels/{countryCode}/{citySlug}/{hotelSlug} [put]
func (h *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateHotelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.UpdateHotel(
		r.Context(),
		chi.URLParam(r, "countryCode"),
		chi.URLParam(r, "citySlug"),
		chi.URLParam(r, "hotelSlug"),
		req,
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// UpdateHotelTitle godoc
// @Summary Update hotel title
// @Tags hotels
// @Accept json
// @Produce json
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Param request body dto.UpdateHotelTitleRequest true "Update hotel title request"
// @Success 200 {object} dto.HotelResponse
// @Router /api/v1/hotels/{countryCode}/{citySlug}/{hotelSlug}/title [patch]
func (h *HotelHandler) UpdateHotelTitle(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateHotelTitleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.UpdateHotelTitle(
		r.Context(),
		chi.URLParam(r, "countryCode"),
		chi.URLParam(r, "citySlug"),
		chi.URLParam(r, "hotelSlug"),
		req,
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// DeleteHotel godoc
// @Summary Delete hotel
// @Tags hotels
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Success 204
// @Router /api/v1/hotels/{countryCode}/{citySlug}/{hotelSlug} [delete]
func (h *HotelHandler) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	if err := h.service.DeleteHotel(
		r.Context(),
		chi.URLParam(r, "countryCode"),
		chi.URLParam(r, "citySlug"),
		chi.URLParam(r, "hotelSlug"),
	); err != nil {
		responder.GRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
