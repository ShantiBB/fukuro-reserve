package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/types/known/timestamppb"

	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/validation"
)

type BookingHandler struct {
	clients *clients.Clients
}

func NewBookingHandler(clients *clients.Clients) *BookingHandler {
	return &BookingHandler{clients: clients}
}

// CreateBooking godoc
// @Summary Create a new booking
// @Tags bookings
// @Accept json
// @Produce json
// @Param request body CreateBookingRequest true "Create booking request"
// @Success 201 {object} BookingResponse
// @Router /api/v1/bookings [post]
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Validate required fields
	if req.GuestName == "" {
		utils.RespondError(w, http.StatusBadRequest, "guest_name is required")
		return
	}

	if err := validation.ValidateUUID(req.HotelId); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validation.ValidateCurrency(req.Currency); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validation.ValidateAmount(req.ExpectedTotalAmount); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate dates
	if req.CheckIn.IsZero() {
		utils.RespondError(w, http.StatusBadRequest, "check_in is required")
		return
	}
	if !req.CheckIn.After(time.Now()) {
		utils.RespondError(w, http.StatusBadRequest, "check_in must be in the future")
		return
	}

	if req.CheckOut.IsZero() {
		utils.RespondError(w, http.StatusBadRequest, "check_out is required")
		return
	}
	if !req.CheckOut.After(req.CheckIn) {
		utils.RespondError(w, http.StatusBadRequest, "check_out must be after check_in")
		return
	}

	// Validate rooms
	if len(req.Rooms) == 0 {
		utils.RespondError(w, http.StatusBadRequest, "at least one room is required")
		return
	}

	// Validate each room
	for i, room := range req.Rooms {
		roomCopy := &validation.CreateBookingRoomRequest{
			RoomId:        room.RoomId,
			Adults:        room.Adults,
			Children:      room.Children,
			PricePerNight: room.PricePerNight,
		}
		if err := validation.ValidateRoom(roomCopy); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "room "+string(rune('A'+i))+": "+err.Error())
			return
		}
	}

	var guestEmail *string
	if req.GuestEmail != "" {
		guestEmail = &req.GuestEmail
	}

	var guestPhone *string
	if req.GuestPhone != "" {
		guestPhone = &req.GuestPhone
	}

	rooms := make([]*bookingv1.CreateBookingRoomRequest, len(req.Rooms))
	for i, room := range req.Rooms {
		rooms[i] = &bookingv1.CreateBookingRoomRequest{
			RoomId:        room.RoomId,
			Adults:        room.Adults,
			Children:      room.Children,
			PricePerNight: room.PricePerNight,
		}
	}

	resp, err := h.clients.Booking.CreateBooking(
		r.Context(), &bookingv1.CreateBookingRequest{
			UserId:              req.UserId,
			HotelId:             req.HotelId,
			CheckIn:             timestamppb.New(req.CheckIn),
			CheckOut:            timestamppb.New(req.CheckOut),
			GuestName:           req.GuestName,
			GuestEmail:          guestEmail,
			GuestPhone:          guestPhone,
			Currency:            req.Currency,
			ExpectedTotalAmount: req.ExpectedTotalAmount,
			Rooms:               rooms,
		},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, bookingResponseFromProto(resp.Booking))
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
// @Success 200 {object} BookingsResponse
// @Router /api/v1/bookings [get]
func (h *BookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	req := &bookingv1.GetBookingsRequest{}

	if userID := r.URL.Query().Get("userId"); userID != "" {
		req.UserId = utils.ParseInt64(userID)
	}
	if hotelID := r.URL.Query().Get("hotelId"); hotelID != "" {
		if err := validation.ValidateUUID(hotelID); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "invalid hotel_id format")
			return
		}
		req.HotelId = hotelID
	}
	if status := r.URL.Query().Get("status"); status != "" {
		req.Status = bookingv1.BookingStatus(bookingv1.BookingStatus_value[status])
	}

	page := utils.ParseUint64(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	limit := utils.ParseUint64(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 100
	}

	req.Page = page
	req.Limit = limit

	resp, err := h.clients.Booking.GetBookings(r.Context(), req)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, bookingsShortResponseFromProto(resp))
}

// GetBooking godoc
// @Summary Get booking by ID
// @Tags bookings
// @Produce json
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} BookingResponse
// @Router /api/v1/bookings/{bookingId} [get]
func (h *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := chi.URLParam(r, "bookingId")
	if bookingID == "" {
		utils.RespondError(w, http.StatusBadRequest, "booking id is required")
		return
	}

	if err := validation.ValidateUUID(bookingID); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid booking id format")
		return
	}

	resp, err := h.clients.Booking.GetBooking(r.Context(), &bookingv1.GetBookingRequest{Id: bookingID})
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, bookingResponseFromProto(resp.Booking))
}

// ConfirmBooking godoc
// @Summary Confirm booking
// @Tags bookings
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} StatusResponse
// @Router /api/v1/bookings/{bookingId}/confirm [patch]
func (h *BookingHandler) ConfirmBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := chi.URLParam(r, "bookingId")
	if bookingID == "" {
		utils.RespondError(w, http.StatusBadRequest, "booking id is required")
		return
	}

	if err := validation.ValidateUUID(bookingID); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid booking id format")
		return
	}

	resp, err := h.clients.Booking.ConfirmBookingStatus(
		r.Context(), &bookingv1.ConfirmBookingStatusRequest{Id: bookingID},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, StatusResponse{Status: resp.Status.String()})
}

// CancelBooking godoc
// @Summary Cancel booking
// @Tags bookings
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} StatusResponse
// @Router /api/v1/bookings/{bookingId}/cancel [patch]
func (h *BookingHandler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := chi.URLParam(r, "bookingId")
	if bookingID == "" {
		utils.RespondError(w, http.StatusBadRequest, "booking id is required")
		return
	}

	if err := validation.ValidateUUID(bookingID); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid booking id format")
		return
	}

	resp, err := h.clients.Booking.CancelBookingStatus(
		r.Context(), &bookingv1.CancelBookingStatusRequest{Id: bookingID},
	)
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, StatusResponse{Status: resp.Status.String()})
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
		utils.RespondError(w, http.StatusBadRequest, "booking id is required")
		return
	}

	if err := validation.ValidateUUID(bookingID); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid booking id format")
		return
	}

	_, err := h.clients.Booking.DeleteBooking(r.Context(), &bookingv1.DeleteBookingRequest{Id: bookingID})
	if err != nil {
		utils.RespondGRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
