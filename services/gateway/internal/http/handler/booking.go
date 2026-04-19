package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	httpmiddleware "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
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

// QuoteBooking godoc
// @Summary       Quote booking price
// @Description   Calculates booking price for requested rooms and dates without creating a booking.
// @Tags          bookings
// @Accept        json
// @Produce       json
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         request body dto.QuoteBookingRequest true "Quote booking request"
// @Success       200 {object} dto.QuoteBookingResponse
// @Failure       400 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/bookings/quote [post]
func (h *BookingHandler) QuoteBooking(c *gin.Context) {
	countryCode, citySlug, hotelID, _, ok := bookingScopeFromPath(c, false)
	if !ok {
		return
	}

	var req dto.QuoteBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.QuoteBooking(c.Request.Context(), countryCode, citySlug, hotelID, req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
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

// GetCurrentUserBookings godoc
// @Summary       Get current user bookings
// @Description   Returns bookings for the current user. Requires JWT auth.
// @Tags          bookings
// @Produce       json
// @Security      Bearer
// @Param         countryCode query string false "Filter by country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug query string false "Filter by city slug" example(tokyo)
// @Param         status query string false "Filter by booking status" example(BOOKING_STATUS_CONFIRMED)
// @Param         page query int false "Page number (starts from 1)" default(1) minimum(1)
// @Param         limit query int false "Page size" default(10) minimum(1) maximum(100)
// @Success       200 {object} dto.BookingsResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Router        /users/me/bookings [get]
func (h *BookingHandler) GetCurrentUserBookings(c *gin.Context) {
	userID, ok := userIDFromContext(c)
	if !ok {
		return
	}

	countryCode := strings.TrimSpace(c.Query("countryCode"))
	citySlug := strings.TrimSpace(c.Query("citySlug"))

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
		"",
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
// @Tags          rooms
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

// GetAvailability godoc
// @Summary       Get available hotel rooms
// @Description   Returns rooms available for the requested date range.
// @Tags          rooms
// @Produce       json
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         check_in query string true "Check-in date in YYYY-MM-DD format" example(2026-05-10)
// @Param         check_out query string true "Check-out date in YYYY-MM-DD format" example(2026-05-13)
// @Param         page query int false "Page number (starts from 1)" default(1) minimum(1)
// @Param         limit query int false "Page size" default(10) minimum(1) maximum(100)
// @Success       200 {object} dto.AvailabilityResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/rooms/availability [get]
func (h *BookingHandler) GetAvailability(c *gin.Context) {
	countryCode, citySlug, hotelID, _, ok := bookingScopeFromPath(c, false)
	if !ok {
		return
	}

	checkIn, ok := requiredDateQuery(c, "check_in", consts.ErrCheckInRequired)
	if !ok {
		return
	}
	checkOut, ok := requiredDateQuery(c, "check_out", consts.ErrCheckOutRequired)
	if !ok {
		return
	}
	if !checkOut.After(checkIn) {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCheckOutMustBeAfterCheckIn})
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

	resp, err := h.service.GetAvailability(c.Request.Context(), countryCode, citySlug, hotelID, checkIn, checkOut, page, limit)
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

// UpdateBookingGuestInfo godoc
// @Summary       Update booking guest info
// @Description   Updates guest contact fields for a booking. Requires JWT auth.
// @Tags          bookings
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         bookingId path string true "Booking ID" example(2f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         request body dto.UpdateBookingGuestInfoRequest true "Update booking guest info request"
// @Success       200 {object} dto.BookingResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/bookings/{bookingId}/guest-info [patch]
func (h *BookingHandler) UpdateBookingGuestInfo(c *gin.Context) {
	countryCode, citySlug, hotelID, bookingID, ok := bookingScopeFromPath(c, true)
	if !ok {
		return
	}

	var req dto.UpdateBookingGuestInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.UpdateBookingGuestInfo(c.Request.Context(), countryCode, citySlug, hotelID, bookingID, req)
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
	countryCode, citySlug, ok = bookingLocationFromPath(c)
	if !ok {
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

func bookingLocationFromPath(c *gin.Context) (countryCode, citySlug string, ok bool) {
	countryCode = c.Param("countryCode")
	if countryCode == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCountryCodeRequired})
		return "", "", false
	}

	citySlug = c.Param("citySlug")
	if citySlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCitySlugRequired})
		return "", "", false
	}

	return countryCode, citySlug, true
}

func userIDFromContext(c *gin.Context) (int64, bool) {
	rawUserID, ok := c.Get(string(httpmiddleware.UserIDKey))
	if !ok {
		c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: consts.ErrInvalidJWTUserIDClaim})
		return 0, false
	}

	userID, ok := rawUserID.(int64)
	if !ok || userID <= 0 {
		c.JSON(http.StatusUnauthorized, &responder.ErrorResponse{Error: consts.ErrInvalidJWTUserIDClaim})
		return 0, false
	}

	return userID, true
}

func requiredDateQuery(c *gin.Context, key string, requiredError string) (time.Time, bool) {
	raw := strings.TrimSpace(c.Query(key))
	if raw == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: requiredError})
		return time.Time{}, false
	}

	value, err := time.Parse("2006-01-02", raw)
	if err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrDateMustBeYYYYMMDD})
		return time.Time{}, false
	}

	return value, true
}
