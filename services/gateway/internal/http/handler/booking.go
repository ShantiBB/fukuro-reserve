package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/request"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
)

// CreateBooking godoc
// @Summary       Create a new booking
// @Description   Creates a booking with guest and room allocation details. Requires JWT auth.
// @Tags          bookings
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         request body dto.CreateBookingRequest true "Create booking request"
// @Success       201 {object} dto.BookingResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/bookings [post]
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	countryCode, citySlug, hotelID, _, ok := bookingScopeFromPath(c, false)
	if !ok {
		return
	}

	var req dto.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}
	resp, err := h.service.CreateBooking(c.Request.Context(), countryCode, citySlug, hotelID, req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetBookings godoc
// @Summary       Get all bookings
// @Description   Returns bookings with optional filters and pagination. Requires JWT auth.
// @Tags          bookings
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         userId query int false "Filter by user ID" minimum(1) example(1)
// @Param         status query string false "Filter by booking status" example(BOOKING_STATUS_CONFIRMED)
// @Param         page query int false "Page number (starts from 1)" default(1) minimum(1)
// @Param         limit query int false "Page size" default(10) minimum(1) maximum(100)
// @Success       200 {object} dto.BookingsResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/bookings [get]
func (h *BookingHandler) GetBookings(c *gin.Context) {
	countryCode, citySlug, hotelID, _, ok := bookingScopeFromPath(c, false)
	if !ok {
		return
	}

	userID, err := request.OptionalPositiveInt64Query(c, "userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrUserIDMustBePositiveInteger})
		return
	}

	page, err := request.OptionalUint64Query(c, "page")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrPageMustBePositiveInteger})
		return
	}
	limit, err := request.OptionalUint64Query(c, "limit")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrLimitMustBePositiveInteger})
		return
	}

	resp, err := h.service.GetBookings(
		c.Request.Context(),
		countryCode,
		citySlug,
		userID,
		hotelID,
		c.Query("status"),
		page,
		limit,
	)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRoomBookings godoc
// @Summary       Get room bookings
// @Description   Returns bookings for a specific room. Requires JWT auth.
// @Tags          bookings
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         roomId path string true "Room ID" example(1f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         status query string false "Filter by booking status" example(BOOKING_STATUS_CONFIRMED)
// @Param         page query int false "Page number (starts from 1)" default(1) minimum(1)
// @Param         limit query int false "Page size" default(10) minimum(1) maximum(100)
// @Success       200 {object} dto.BookingsResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/rooms/{roomId}/bookings [get]
func (h *BookingHandler) GetRoomBookings(c *gin.Context) {
	countryCode, citySlug, hotelID, _, ok := bookingScopeFromPath(c, false)
	if !ok {
		return
	}

	roomID := c.Param("roomId")
	if roomID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrRoomIDRequired})
		return
	}

	page, err := request.OptionalUint64Query(c, "page")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrPageMustBePositiveInteger})
		return
	}
	limit, err := request.OptionalUint64Query(c, "limit")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrLimitMustBePositiveInteger})
		return
	}

	resp, err := h.service.GetRoomBookings(
		c.Request.Context(),
		countryCode,
		citySlug,
		hotelID,
		roomID,
		c.Query("status"),
		page,
		limit,
	)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetBooking godoc
// @Summary       Get booking by ID
// @Description   Returns a booking by ID. Requires JWT auth.
// @Tags          bookings
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         bookingId path string true "Booking ID" example(2f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       200 {object} dto.BookingResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/bookings/{bookingId} [get]
func (h *BookingHandler) GetBooking(c *gin.Context) {
	countryCode, citySlug, _, bookingID, ok := bookingScopeFromPath(c, true)
	if !ok {
		return
	}

	resp, err := h.service.GetBooking(c.Request.Context(), countryCode, citySlug, bookingID)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ConfirmBooking godoc
// @Summary       Confirm booking
// @Description   Confirms a booking by ID. Requires JWT auth.
// @Tags          bookings
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         bookingId path string true "Booking ID" example(2f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       200 {object} dto.StatusResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/bookings/{bookingId}/confirm [patch]
func (h *BookingHandler) ConfirmBooking(c *gin.Context) {
	countryCode, citySlug, _, bookingID, ok := bookingScopeFromPath(c, true)
	if !ok {
		return
	}

	resp, err := h.service.ConfirmBooking(c.Request.Context(), countryCode, citySlug, bookingID)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CancelBooking godoc
// @Summary       Cancel booking
// @Description   Cancels a booking by ID. Requires JWT auth.
// @Tags          bookings
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         bookingId path string true "Booking ID" example(2f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       200 {object} dto.StatusResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/bookings/{bookingId}/cancel [patch]
func (h *BookingHandler) CancelBooking(c *gin.Context) {
	countryCode, citySlug, _, bookingID, ok := bookingScopeFromPath(c, true)
	if !ok {
		return
	}

	resp, err := h.service.CancelBooking(c.Request.Context(), countryCode, citySlug, bookingID)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteBooking godoc
// @Summary       Delete booking
// @Description   Deletes a booking by ID. Requires JWT auth.
// @Tags          bookings
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         bookingId path string true "Booking ID" example(2f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       204
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/bookings/{bookingId} [delete]
func (h *BookingHandler) DeleteBooking(c *gin.Context) {
	countryCode, citySlug, _, bookingID, ok := bookingScopeFromPath(c, true)
	if !ok {
		return
	}

	if err := h.service.DeleteBooking(c.Request.Context(), countryCode, citySlug, bookingID); err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func bookingScopeFromPath(c *gin.Context, requireBookingID bool) (countryCode, citySlug, hotelID, bookingID string, ok bool) {
	countryCode = c.Param("countryCode")
	if countryCode == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCountryCodeRequired})
		return "", "", "", "", false
	}

	citySlug = c.Param("citySlug")
	if citySlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCitySlugRequired})
		return "", "", "", "", false
	}

	hotelID = c.Param("hotelId")
	if hotelID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrHotelIDRequired})
		return "", "", "", "", false
	}

	if requireBookingID {
		bookingID = c.Param("bookingId")
		if bookingID == "" {
			c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrBookingIDRequired})
			return "", "", "", "", false
		}
	}

	return countryCode, citySlug, hotelID, bookingID, true
}
