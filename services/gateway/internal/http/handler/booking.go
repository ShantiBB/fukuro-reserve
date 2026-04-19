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
// @Param         request body dto.CreateBookingRequest true "Create booking request"
// @Success       201 {object} dto.BookingResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Router        /bookings [post]
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req dto.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.CreateBooking(c.Request.Context(), req)
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
// @Param         userId query int false "Filter by user ID" minimum(1) example(1)
// @Param         hotelId query string false "Filter by hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         status query string false "Filter by booking status" example(BOOKING_STATUS_CONFIRMED)
// @Param         page query int false "Page number (starts from 1)" default(1) minimum(1)
// @Param         limit query int false "Page size" default(10) minimum(1) maximum(100)
// @Success       200 {object} dto.BookingsResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Router        /bookings [get]
func (h *BookingHandler) GetBookings(c *gin.Context) {
	userID, err := request.OptionalPositiveInt64Query(c, "userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrUserIDMustBePositiveInteger})
		return
	}

	hotelID := c.Query("hotelId")

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

// GetBooking godoc
// @Summary       Get booking by ID
// @Description   Returns a booking by ID. Requires JWT auth.
// @Tags          bookings
// @Produce       json
// @Security      Bearer
// @Param         bookingId path string true "Booking ID" example(2f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       200 {object} dto.BookingResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /bookings/{bookingId} [get]
func (h *BookingHandler) GetBooking(c *gin.Context) {
	bookingID := c.Param("bookingId")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrBookingIDRequired})
		return
	}

	resp, err := h.service.GetBooking(c.Request.Context(), bookingID)
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
// @Param         bookingId path string true "Booking ID" example(2f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       200 {object} dto.StatusResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /bookings/{bookingId}/confirm [patch]
func (h *BookingHandler) ConfirmBooking(c *gin.Context) {
	bookingID := c.Param("bookingId")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrBookingIDRequired})
		return
	}

	resp, err := h.service.ConfirmBooking(c.Request.Context(), bookingID)
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
// @Param         bookingId path string true "Booking ID" example(2f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       200 {object} dto.StatusResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /bookings/{bookingId}/cancel [patch]
func (h *BookingHandler) CancelBooking(c *gin.Context) {
	bookingID := c.Param("bookingId")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrBookingIDRequired})
		return
	}

	resp, err := h.service.CancelBooking(c.Request.Context(), bookingID)
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
// @Param         bookingId path string true "Booking ID" example(2f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       204
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /bookings/{bookingId} [delete]
func (h *BookingHandler) DeleteBooking(c *gin.Context) {
	bookingID := c.Param("bookingId")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrBookingIDRequired})
		return
	}

	if err := h.service.DeleteBooking(c.Request.Context(), bookingID); err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
