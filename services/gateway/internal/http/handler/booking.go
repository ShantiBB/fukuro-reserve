package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/query"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/validation"
	bookingservice "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/service/booking"
)

type BookingHandler struct {
	service *bookingservice.Service
}

func NewBookingHandler(service *bookingservice.Service) *BookingHandler {
	return &BookingHandler{service: service}
}

// CreateBooking godoc
// @Summary Create a new booking
// @Tags bookings
// @Accept json
// @Produce json
// @Param request body dto.CreateBookingRequest true "Create booking request"
// @Success 201 {object} dto.BookingResponse
// @Router /api/v1/bookings [post]
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.GuestName == "" {
		responder.Error(w, http.StatusBadRequest, "guest_name is required")
		return
	}
	if err := validation.ValidateUUID(req.HotelId); err != nil {
		responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validation.ValidateCurrency(req.Currency); err != nil {
		responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validation.ValidateAmount(req.ExpectedTotalAmount); err != nil {
		responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	if req.CheckIn.IsZero() {
		responder.Error(w, http.StatusBadRequest, "check_in is required")
		return
	}
	if !req.CheckIn.After(time.Now()) {
		responder.Error(w, http.StatusBadRequest, "check_in must be in the future")
		return
	}
	if req.CheckOut.IsZero() {
		responder.Error(w, http.StatusBadRequest, "check_out is required")
		return
	}
	if !req.CheckOut.After(req.CheckIn) {
		responder.Error(w, http.StatusBadRequest, "check_out must be after check_in")
		return
	}
	if len(req.Rooms) == 0 {
		responder.Error(w, http.StatusBadRequest, "at least one room is required")
		return
	}
	for i, room := range req.Rooms {
		roomCopy := &validation.CreateBookingRoomRequest{RoomId: room.RoomId, Adults: room.Adults, Children: room.Children, PricePerNight: room.PricePerNight}
		if err := validation.ValidateRoom(roomCopy); err != nil {
			responder.Error(w, http.StatusBadRequest, "room "+string(rune('A'+i))+": "+err.Error())
			return
		}
	}

	resp, err := h.service.CreateBooking(r.Context(), req)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusCreated, resp)
}

// GetBookings godoc
// @Summary Get all bookings
// @Tags bookings
// @Produce json
// @Param userId query int false "User ID"
// @Param hotelId query string false "Hotel ID"
// @Param status query string false "Status"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {object} dto.BookingsResponse
// @Router /api/v1/bookings [get]
func (h *BookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	userID := int64(0)
	if user := r.URL.Query().Get("userId"); user != "" {
		userID = query.ParseInt64(user)
	}
	if hotelID := r.URL.Query().Get("hotelId"); hotelID != "" {
		if err := validation.ValidateUUID(hotelID); err != nil {
			responder.Error(w, http.StatusBadRequest, "invalid hotel_id format")
			return
		}
	}

	resp, err := h.service.GetBookings(
		r.Context(),
		userID,
		r.URL.Query().Get("hotelId"),
		r.URL.Query().Get("status"),
		query.ParseUint64(r.URL.Query().Get("page")),
		query.ParseUint64(r.URL.Query().Get("limit")),
	)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// GetBooking godoc
// @Summary Get booking by ID
// @Tags bookings
// @Produce json
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} dto.BookingResponse
// @Router /api/v1/bookings/{bookingId} [get]
func (h *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := chi.URLParam(r, "bookingId")
	if bookingID == "" {
		responder.Error(w, http.StatusBadRequest, "booking id is required")
		return
	}
	if err := validation.ValidateUUID(bookingID); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid booking id format")
		return
	}

	resp, err := h.service.GetBooking(r.Context(), bookingID)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// ConfirmBooking godoc
// @Summary Confirm booking
// @Tags bookings
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} dto.StatusResponse
// @Router /api/v1/bookings/{bookingId}/confirm [patch]
func (h *BookingHandler) ConfirmBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := chi.URLParam(r, "bookingId")
	if bookingID == "" {
		responder.Error(w, http.StatusBadRequest, "booking id is required")
		return
	}
	if err := validation.ValidateUUID(bookingID); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid booking id format")
		return
	}

	resp, err := h.service.ConfirmBooking(r.Context(), bookingID)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// CancelBooking godoc
// @Summary Cancel booking
// @Tags bookings
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} dto.StatusResponse
// @Router /api/v1/bookings/{bookingId}/cancel [patch]
func (h *BookingHandler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := chi.URLParam(r, "bookingId")
	if bookingID == "" {
		responder.Error(w, http.StatusBadRequest, "booking id is required")
		return
	}
	if err := validation.ValidateUUID(bookingID); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid booking id format")
		return
	}

	resp, err := h.service.CancelBooking(r.Context(), bookingID)
	if err != nil {
		responder.GRPCError(w, err)
		return
	}

	responder.JSON(w, http.StatusOK, resp)
}

// DeleteBooking godoc
// @Summary Delete booking
// @Tags bookings
// @Param bookingId path string true "Booking ID"
// @Success 204
// @Router /api/v1/bookings/{bookingId} [delete]
func (h *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := chi.URLParam(r, "bookingId")
	if bookingID == "" {
		responder.Error(w, http.StatusBadRequest, "booking id is required")
		return
	}
	if err := validation.ValidateUUID(bookingID); err != nil {
		responder.Error(w, http.StatusBadRequest, "invalid booking id format")
		return
	}

	if err := h.service.DeleteBooking(r.Context(), bookingID); err != nil {
		responder.GRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
