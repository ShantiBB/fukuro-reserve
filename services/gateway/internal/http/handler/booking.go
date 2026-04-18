package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/request"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/validation"
)

// CreateBooking godoc
// @Summary Create a new booking
// @Tags bookings
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body dto.CreateBookingRequest true "Create booking request"
// @Success 201 {object} dto.BookingResponse
// @Router /bookings [post]
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req dto.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}
	if err := validation.ValidateCreateBookingRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
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
// @Summary Get all bookings
// @Tags bookings
// @Produce json
// @Security Bearer
// @Param userId query int false "User ID"
// @Param hotelId query string false "Hotel ID"
// @Param status query string false "Status"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {object} dto.BookingsResponse
// @Router /bookings [get]
func (h *BookingHandler) GetBookings(c *gin.Context) {
	userID, err := request.OptionalPositiveInt64Query(c, "userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrUserIDMustBePositiveInteger})
		return
	}

	hotelID := c.Query("hotelId")
	if hotelID != "" {
		if err = validation.ValidateUUID(hotelID); err != nil {
			c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidHotelIDFormat})
			return
		}
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
// @Summary Get booking by ID
// @Tags bookings
// @Produce json
// @Security Bearer
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} dto.BookingResponse
// @Router /bookings/{bookingId} [get]
func (h *BookingHandler) GetBooking(c *gin.Context) {
	bookingID := c.Param("bookingId")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrBookingIDRequired})
		return
	}
	if err := validation.ValidateUUID(bookingID); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidBookingIDFormat})
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
// @Summary Confirm booking
// @Tags bookings
// @Security Bearer
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} dto.StatusResponse
// @Router /bookings/{bookingId}/confirm [patch]
func (h *BookingHandler) ConfirmBooking(c *gin.Context) {
	bookingID := c.Param("bookingId")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrBookingIDRequired})
		return
	}
	if err := validation.ValidateUUID(bookingID); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidBookingIDFormat})
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
// @Summary Cancel booking
// @Tags bookings
// @Security Bearer
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} dto.StatusResponse
// @Router /bookings/{bookingId}/cancel [patch]
func (h *BookingHandler) CancelBooking(c *gin.Context) {
	bookingID := c.Param("bookingId")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrBookingIDRequired})
		return
	}
	if err := validation.ValidateUUID(bookingID); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidBookingIDFormat})
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
// @Summary Delete booking
// @Tags bookings
// @Security Bearer
// @Param bookingId path string true "Booking ID"
// @Success 204
// @Router /bookings/{bookingId} [delete]
func (h *BookingHandler) DeleteBooking(c *gin.Context) {
	bookingID := c.Param("bookingId")
	if bookingID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrBookingIDRequired})
		return
	}
	if err := validation.ValidateUUID(bookingID); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidBookingIDFormat})
		return
	}

	if err := h.service.DeleteBooking(c.Request.Context(), bookingID); err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
