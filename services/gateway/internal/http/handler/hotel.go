package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/mapper"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/query"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/responder"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

type HotelHandler struct {
	clients    *clients.Clients
	pagination config.PaginationConfig
}

func NewHotelHandler(clients *clients.Clients, pagination config.PaginationConfig) *HotelHandler {
	return &HotelHandler{clients: clients, pagination: pagination}
}

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

	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	location := &hotelv1.CreateHotelLocationRequest{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}

	resp, err := h.clients.Hotel.CreateHotel(
		r.Context(), &hotelv1.CreateHotelRequest{
			CountryCode: req.CountryCode,
			CitySlug:    req.CitySlug,
			Title:       req.Title,
			OwnerId:     req.OwnerId,
			Description: description,
			Address:     req.Address,
			Location:    location,
		},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusCreated, mapper.HotelResponseFromProto(resp.Hotel))
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

	page := query.ParseUint64(r.URL.Query().Get("page"))
	if page == 0 {
		page = h.pagination.DefaultPage
	}

	limit := query.ParseUint64(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = h.pagination.DefaultPageSize
	}

	resp, err := h.clients.Hotel.GetHotels(
		r.Context(), &hotelv1.GetHotelsRequest{
			CountryCode: countryCode,
			CitySlug:    citySlug,
			SortBy:      sortBy,
			Page:        page,
			Limit:       limit,
		},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, mapper.HotelsShortResponseFromProto(resp))
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
	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	hotelSlug := chi.URLParam(r, "hotelSlug")

	resp, err := h.clients.Hotel.GetHotel(
		r.Context(), &hotelv1.GetHotelRequest{
			CountryCode: countryCode,
			CitySlug:    citySlug,
			HotelSlug:   hotelSlug,
		},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, mapper.HotelDetailResponseFromProto(resp))
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
	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	hotelSlug := chi.URLParam(r, "hotelSlug")

	var req dto.UpdateHotelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	location := &hotelv1.UpdateHotelLocationRequest{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}

	resp, err := h.clients.Hotel.UpdateHotel(
		r.Context(), &hotelv1.UpdateHotelRequest{
			CountryCode: countryCode,
			CitySlug:    citySlug,
			HotelSlug:   hotelSlug,
			Description: description,
			Address:     req.Address,
			Location:    location,
		},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, mapper.UpdateHotelResponseFromProto(resp.Hotel))
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
	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	hotelSlug := chi.URLParam(r, "hotelSlug")

	var req dto.UpdateHotelTitleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.clients.Hotel.UpdateHotelTitle(
		r.Context(), &hotelv1.UpdateHotelTitleRequest{
			CountryCode: countryCode,
			CitySlug:    citySlug,
			HotelSlug:   hotelSlug,
			Title:       req.Title,
		},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, mapper.UpdateHotelTitleResponseFromProto(resp.Hotel))
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
	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	hotelSlug := chi.URLParam(r, "hotelSlug")

	_, err := h.clients.Hotel.DeleteHotel(
		r.Context(), &hotelv1.DeleteHotelRequest{
			CountryCode: countryCode,
			CitySlug:    citySlug,
			HotelSlug:   hotelSlug,
		},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

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

	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	var price float32
	if req.Price != "" {
		var p float64
		if _, err := fmt.Sscanf(req.Price, "%f", &p); err != nil {
			responder.Error(w, http.StatusBadRequest, "invalid price")
			return
		}
		price = float32(p)
	}

	if req.CountryCode == "" || req.CitySlug == "" || req.HotelSlug == "" {
		responder.Error(
			w,
			http.StatusBadRequest,
			"country_code, city_slug, and hotel_slug are required",
		)
		return
	}

	resp, err := h.clients.Room.CreateRoom(
		r.Context(), &hotelv1.CreateRoomRequest{
			CountryCode: req.CountryCode,
			CitySlug:    req.CitySlug,
			HotelSlug:   req.HotelSlug,
			Title:       req.Title,
			Description: description,
			RoomNumber:  req.RoomNumber,
			Type:        hotelv1.RoomType(hotelv1.RoomType_value[req.Type]),
			Price:       price,
			Capacity:    req.Capacity,
			AreaSqm:     req.AreaSqm,
			Floor:       req.Floor,
			Amenities:   req.Amenities,
			Images:      req.Images,
		},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusCreated, mapper.RoomResponseFromProto(resp.Room))
}

// GetRooms godoc
// @Summary Get rooms
// @Tags rooms
// @Produce json
// @Param country_code query string true "Country code"
// @Param city_slug query string true "City slug"
// @Param hotel_slug query string true "Hotel slug"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
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

	page := query.ParseUint64(r.URL.Query().Get("page"))
	if page == 0 {
		page = h.pagination.DefaultPage
	}

	limit := query.ParseUint64(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = h.pagination.DefaultPageSize
	}

	resp, err := h.clients.Room.GetRooms(
		r.Context(), &hotelv1.GetRoomsRequest{
			CountryCode: countryCode,
			CitySlug:    citySlug,
			HotelSlug:   hotelSlug,
			Page:        page,
			Limit:       limit,
		},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, mapper.RoomsShortResponseFromProto(resp))
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

	resp, err := h.clients.Room.GetRoom(
		r.Context(), &hotelv1.GetRoomRequest{Id: roomID},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, mapper.RoomResponseFromProto(resp.Room))
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

	var description string
	if req.Description != "" {
		description = req.Description
	}

	resp, err := h.clients.Room.UpdateRoom(
		r.Context(), &hotelv1.UpdateRoomRequest{
			Id:          roomID,
			Title:       req.Title,
			RoomNumber:  req.RoomNumber,
			Type:        hotelv1.RoomType(hotelv1.RoomType_value[req.Type]),
			Description: description,
			Price:       req.Price,
			Capacity:    req.Capacity,
			AreaSqm:     req.AreaSqm,
			Floor:       req.Floor,
			Amenities:   req.Amenities,
			Images:      req.Images,
		},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, mapper.UpdateRoomResponseFromProto(resp.Room))
}

// UpdateRoomStatus godoc
// @Summary Update room status
// @Tags rooms
// @Accept json
// @Produce json
// @Param roomId path string true "Room ID"
// @Param request body dto.UpdateRoomStatusRequest true "Update status request"
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

	resp, err := h.clients.Room.UpdateRoomStatus(
		r.Context(), &hotelv1.UpdateRoomStatusRequest{
			Id:     roomID,
			Status: hotelv1.RoomStatus(hotelv1.RoomStatus_value[req.Status]),
		},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, dto.StatusResponse{Status: resp.Status.String()})
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

	_, err := h.clients.Room.DeleteRoom(
		r.Context(), &hotelv1.DeleteRoomRequest{Id: roomID},
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
