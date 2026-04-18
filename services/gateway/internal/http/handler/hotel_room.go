package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/query"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
)

// CreateRoom godoc
// @Summary Create a new room
// @Tags rooms
// @Accept json
// @Produce json
// @Param request body dto.CreateRoomRequest true "Create room request"
// @Success 201 {object} dto.RoomResponse
// @Router /api/v1/rooms [post]
func (h *HotelHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.CreateRoom(r.Context(), req)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusCreated, resp)
}

// GetRooms godoc
// @Summary Get rooms
// @Tags rooms
// @Produce json
// @Param countryCode query string true "Country code"
// @Param citySlug query string true "City slug"
// @Param hotelSlug query string true "Hotel slug"
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Success 200 {object} dto.RoomsResponse
// @Router /api/v1/rooms [get]
func (h *HotelHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Query().Get("countryCode")
	if countryCode == "" {
		countryCode = r.URL.Query().Get("country_code")
	}

	citySlug := r.URL.Query().Get("citySlug")
	if citySlug == "" {
		citySlug = r.URL.Query().Get("city_slug")
	}

	hotelSlug := r.URL.Query().Get("hotelSlug")
	if hotelSlug == "" {
		hotelSlug = r.URL.Query().Get("hotel_slug")
	}

	resp, err := h.service.GetRooms(
		r.Context(),
		countryCode,
		citySlug,
		hotelSlug,
		query.ParseUint64(r.URL.Query().Get("page")),
		query.ParseUint64(r.URL.Query().Get("limit")),
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// GetRoom godoc
// @Summary Get room by ID
// @Tags rooms
// @Produce json
// @Param roomId path string true "Room ID"
// @Success 200 {object} dto.RoomResponse
// @Router /api/v1/rooms/{roomId} [get]
func (h *HotelHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomId")
	if roomID == "" {
		responder.Error(w, http.StatusBadRequest, "room id is required")
		return
	}

	resp, err := h.service.GetRoom(r.Context(), roomID)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// UpdateRoom godoc
// @Summary Update room
// @Tags rooms
// @Accept json
// @Produce json
// @Param roomId path string true "Room ID"
// @Param request body dto.UpdateRoomRequest true "Update room request"
// @Success 200 {object} dto.RoomResponse
// @Router /api/v1/rooms/{roomId} [put]
func (h *HotelHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomId")
	if roomID == "" {
		responder.Error(w, http.StatusBadRequest, "room id is required")
		return
	}

	var req dto.UpdateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.UpdateRoom(r.Context(), roomID, req)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// UpdateRoomStatus godoc
// @Summary Update room status
// @Tags rooms
// @Accept json
// @Produce json
// @Param roomId path string true "Room ID"
// @Param request body dto.UpdateRoomStatusRequest true "Update room status request"
// @Success 200 {object} dto.StatusResponse
// @Router /api/v1/rooms/{roomId}/status [patch]
func (h *HotelHandler) UpdateRoomStatus(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomId")
	if roomID == "" {
		responder.Error(w, http.StatusBadRequest, "room id is required")
		return
	}

	var req dto.UpdateRoomStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.service.UpdateRoomStatus(r.Context(), roomID, req)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// DeleteRoom godoc
// @Summary Delete room
// @Tags rooms
// @Param roomId path string true "Room ID"
// @Success 204
// @Router /api/v1/rooms/{roomId} [delete]
func (h *HotelHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomId")
	if roomID == "" {
		responder.Error(w, http.StatusBadRequest, "room id is required")
		return
	}

	if err := h.service.DeleteRoom(r.Context(), roomID); err != nil {
		responder.GRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
