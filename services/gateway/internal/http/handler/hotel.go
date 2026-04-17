package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

type HotelHandler struct {
	clients *clients.Clients
}

func NewHotelHandler(clients *clients.Clients) *HotelHandler {
	return &HotelHandler{clients: clients}
}

func (h *HotelHandler) HotelRoutes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware)

	r.Post("/", h.CreateHotel)
	r.Get("/", h.GetHotels)
	r.Get("/{countryCode}/{citySlug}/{hotelSlug}", h.GetHotel)
	r.Put("/{countryCode}/{citySlug}/{hotelSlug}", h.UpdateHotel)
	r.Patch("/{countryCode}/{citySlug}/{hotelSlug}/title", h.UpdateHotelTitle)
	r.Delete("/{countryCode}/{citySlug}/{hotelSlug}", h.DeleteHotel)

	return r
}

func (h *HotelHandler) RoomRoutes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware)

	r.Post("/", h.CreateRoom)
	r.Get("/", h.GetRooms)
	r.Get("/{roomId}", h.GetRoom)
	r.Put("/{roomId}", h.UpdateRoom)
	r.Patch("/{roomId}/status", h.UpdateRoomStatus)
	r.Delete("/{roomId}", h.DeleteRoom)

	return r
}

// CreateHotel godoc
// @Summary Create a new hotel
// @Tags hotels
// @Accept json
// @Produce json
// @Param request body CreateHotelRequest true "Create hotel request"
// @Success 201 {object} HotelResponse
// @Router /api/v1/hotels [post]
func (h *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	var req CreateHotelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
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
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, hotelResponseFromProto(resp.Hotel))
}

// GetHotels godoc
// @Summary Get hotels
// @Tags hotels
// @Produce json
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {array} HotelShortResponse
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

	page := utils.ParseUint64(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	limit := utils.ParseUint64(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 100
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
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, hotelsShortResponseFromProto(resp))
}

// GetHotel godoc
// @Summary Get hotel by slug
// @Tags hotels
// @Produce json
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Success 200 {object} HotelResponse
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
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, hotelDetailResponseFromProto(resp))
}

// UpdateHotel godoc
// @Summary Update hotel
// @Tags hotels
// @Accept json
// @Produce json
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Param request body UpdateHotelRequest true "Update hotel request"
// @Success 200 {object} HotelResponse
// @Router /api/v1/hotels/{countryCode}/{citySlug}/{hotelSlug} [put]
func (h *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	hotelSlug := chi.URLParam(r, "hotelSlug")

	var req UpdateHotelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
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
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, updateHotelResponseFromProto(resp.Hotel))
}

// UpdateHotelTitle godoc
// @Summary Update hotel title
// @Tags hotels
// @Accept json
// @Produce json
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Param request body UpdateHotelTitleRequest true "Update hotel title request"
// @Success 200 {object} HotelResponse
// @Router /api/v1/hotels/{countryCode}/{citySlug}/{hotelSlug}/title [patch]
func (h *HotelHandler) UpdateHotelTitle(w http.ResponseWriter, r *http.Request) {
	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	hotelSlug := chi.URLParam(r, "hotelSlug")

	var req UpdateHotelTitleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
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
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, updateHotelTitleResponseFromProto(resp.Hotel))
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
		utils.RespondGRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreateRoom godoc
// @Summary Create a new room
// @Tags rooms
// @Accept json
// @Produce json
// @Param request body CreateRoomRequest true "Create room request"
// @Success 201 {object} RoomResponse
// @Router /api/v1/rooms [post]
func (h *HotelHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req CreateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
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
			utils.RespondError(w, http.StatusBadRequest, "invalid price")
			return
		}
		price = float32(p)
	}

	if req.CountryCode == "" || req.CitySlug == "" || req.HotelSlug == "" {
		utils.RespondError(
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
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, roomResponseFromProto(resp.Room))
}

// GetRooms godoc
// @Summary Get rooms
// @Tags rooms
// @Produce json
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {array} RoomShortResponse
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

	page := utils.ParseUint64(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	limit := utils.ParseUint64(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 100
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
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, roomsShortResponseFromProto(resp))
}

// GetRoom godoc
// @Summary Get room by ID
// @Tags rooms
// @Produce json
// @Param roomId path string true "Room ID"
// @Success 200 {object} RoomResponse
// @Router /api/v1/rooms/{roomId} [get]
func (h *HotelHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomId")
	if roomID == "" {
		utils.RespondError(w, http.StatusBadRequest, "room id is required")
		return
	}

	resp, err := h.clients.Room.GetRoom(
		r.Context(), &hotelv1.GetRoomRequest{Id: roomID},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, roomResponseFromProto(resp.Room))
}

// UpdateRoom godoc
// @Summary Update room
// @Tags rooms
// @Accept json
// @Produce json
// @Param roomId path string true "Room ID"
// @Param request body UpdateRoomRequest true "Update room request"
// @Success 200 {object} RoomResponse
// @Router /api/v1/rooms/{roomId} [put]
func (h *HotelHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomId")
	if roomID == "" {
		utils.RespondError(w, http.StatusBadRequest, "room id is required")
		return
	}

	var req UpdateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
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
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, updateRoomResponseFromProto(resp.Room))
}

// UpdateRoomStatus godoc
// @Summary Update room status
// @Tags rooms
// @Accept json
// @Produce json
// @Param roomId path string true "Room ID"
// @Param request body UpdateRoomStatusRequest true "Update status request"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/rooms/{roomId}/status [patch]
func (h *HotelHandler) UpdateRoomStatus(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomId")
	if roomID == "" {
		utils.RespondError(w, http.StatusBadRequest, "room id is required")
		return
	}

	var req UpdateRoomStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	resp, err := h.clients.Room.UpdateRoomStatus(
		r.Context(), &hotelv1.UpdateRoomStatusRequest{
			Id:     roomID,
			Status: hotelv1.RoomStatus(hotelv1.RoomStatus_value[req.Status]),
		},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{"status": resp.Status.String()})
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
		utils.RespondError(w, http.StatusBadRequest, "room id is required")
		return
	}

	_, err := h.clients.Room.DeleteRoom(
		r.Context(), &hotelv1.DeleteRoomRequest{Id: roomID},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
