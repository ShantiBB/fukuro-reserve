package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/request"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
)

// CreateRoom godoc
// @Summary Create a new room
// @Tags rooms
// @Accept json
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelId path string true "Hotel ID"
// @Param request body dto.CreateRoomRequest true "Create room request"
// @Success 201 {object} dto.RoomResponse
// @Router /{countryCode}/{citySlug}/hotels/{hotelId}/rooms [post]
func (h *HotelHandler) CreateRoom(c *gin.Context) {
	countryCode := c.Param("countryCode")
	if countryCode == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCountryCodeRequired})
		return
	}

	citySlug := c.Param("citySlug")
	if citySlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCitySlugRequired})
		return
	}

	hotelID := c.Param("hotelId")
	if hotelID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrHotelIDRequired})
		return
	}

	var req dto.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.CreateRoomByHotelID(c.Request.Context(), hotelID, req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetRoomsByHotelSlug godoc
// @Summary Get rooms by hotel slug
// @Tags rooms
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Success 200 {object} dto.RoomsResponse
// @Router /{countryCode}/{citySlug}/hotels/slug/{hotelSlug}/rooms [get]
func (h *HotelHandler) GetRoomsByHotelSlug(c *gin.Context) {
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

	countryCode := c.Param("countryCode")
	if countryCode == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCountryCodeRequired})
		return
	}
	citySlug := c.Param("citySlug")
	if citySlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCitySlugRequired})
		return
	}
	hotelSlug := c.Param("hotelSlug")
	if hotelSlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrRoomHotelSlugReq})
		return
	}

	resp, err := h.service.GetRooms(
		c.Request.Context(),
		countryCode,
		citySlug,
		hotelSlug,
		page,
		limit,
	)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRoomsByHotelID godoc
// @Summary Get rooms by hotel ID
// @Tags rooms
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelId path string true "Hotel ID"
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Success 200 {object} dto.RoomsResponse
// @Router /{countryCode}/{citySlug}/hotels/{hotelId}/rooms [get]
func (h *HotelHandler) GetRoomsByHotelID(c *gin.Context) {
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

	countryCode := c.Param("countryCode")
	if countryCode == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCountryCodeRequired})
		return
	}

	citySlug := c.Param("citySlug")
	if citySlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCitySlugRequired})
		return
	}

	hotelID := c.Param("hotelId")
	if hotelID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrHotelIDRequired})
		return
	}

	resp, err := h.service.GetRoomsByHotelID(c.Request.Context(), hotelID, page, limit)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRoomByHotelSlug godoc
// @Summary Get room by hotel slug and room ID
// @Tags rooms
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Param roomId path string true "Room ID"
// @Success 200 {object} dto.RoomResponse
// @Router /{countryCode}/{citySlug}/hotels/slug/{hotelSlug}/rooms/{roomId} [get]
func (h *HotelHandler) GetRoomByHotelSlug(c *gin.Context) {
	countryCode := c.Param("countryCode")
	if countryCode == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCountryCodeRequired})
		return
	}

	citySlug := c.Param("citySlug")
	if citySlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCitySlugRequired})
		return
	}

	hotelSlug := c.Param("hotelSlug")
	if hotelSlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrRoomHotelSlugReq})
		return
	}

	roomID := c.Param("roomId")
	if roomID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrRoomIDRequired})
		return
	}

	resp, err := h.service.GetRoom(c.Request.Context(), roomID)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRoomByHotelID godoc
// @Summary Get room by hotel ID and room ID
// @Tags rooms
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelId path string true "Hotel ID"
// @Param roomId path string true "Room ID"
// @Success 200 {object} dto.RoomResponse
// @Router /{countryCode}/{citySlug}/hotels/{hotelId}/rooms/{roomId} [get]
func (h *HotelHandler) GetRoomByHotelID(c *gin.Context) {
	countryCode := c.Param("countryCode")
	if countryCode == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCountryCodeRequired})
		return
	}

	citySlug := c.Param("citySlug")
	if citySlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCitySlugRequired})
		return
	}

	hotelID := c.Param("hotelId")
	if hotelID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrHotelIDRequired})
		return
	}

	roomID := c.Param("roomId")
	if roomID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrRoomIDRequired})
		return
	}

	resp, err := h.service.GetRoom(c.Request.Context(), roomID)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateRoom godoc
// @Summary Update room
// @Tags rooms
// @Accept json
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelId path string true "Hotel ID"
// @Param roomId path string true "Room ID"
// @Param request body dto.UpdateRoomRequest true "Update room request"
// @Success 200 {object} dto.RoomResponse
// @Router /{countryCode}/{citySlug}/hotels/{hotelId}/rooms/{roomId} [put]
func (h *HotelHandler) UpdateRoom(c *gin.Context) {
	countryCode := c.Param("countryCode")
	if countryCode == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCountryCodeRequired})
		return
	}

	citySlug := c.Param("citySlug")
	if citySlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCitySlugRequired})
		return
	}

	hotelID := c.Param("hotelId")
	if hotelID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrHotelIDRequired})
		return
	}

	roomID := c.Param("roomId")
	if roomID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrRoomIDRequired})
		return
	}

	var req dto.UpdateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.UpdateRoom(c.Request.Context(), roomID, req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateRoomStatus godoc
// @Summary Update room status
// @Tags rooms
// @Accept json
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelId path string true "Hotel ID"
// @Param roomId path string true "Room ID"
// @Param request body dto.UpdateRoomStatusRequest true "Update room status request"
// @Success 200 {object} dto.StatusResponse
// @Router /{countryCode}/{citySlug}/hotels/{hotelId}/rooms/{roomId}/status [patch]
func (h *HotelHandler) UpdateRoomStatus(c *gin.Context) {
	countryCode := c.Param("countryCode")
	if countryCode == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCountryCodeRequired})
		return
	}

	citySlug := c.Param("citySlug")
	if citySlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCitySlugRequired})
		return
	}

	hotelID := c.Param("hotelId")
	if hotelID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrHotelIDRequired})
		return
	}

	roomID := c.Param("roomId")
	if roomID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrRoomIDRequired})
		return
	}

	var req dto.UpdateRoomStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.UpdateRoomStatus(c.Request.Context(), roomID, req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteRoom godoc
// @Summary Delete room
// @Tags rooms
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelId path string true "Hotel ID"
// @Param roomId path string true "Room ID"
// @Success 204
// @Router /{countryCode}/{citySlug}/hotels/{hotelId}/rooms/{roomId} [delete]
func (h *HotelHandler) DeleteRoom(c *gin.Context) {
	countryCode := c.Param("countryCode")
	if countryCode == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCountryCodeRequired})
		return
	}

	citySlug := c.Param("citySlug")
	if citySlug == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrCitySlugRequired})
		return
	}

	hotelID := c.Param("hotelId")
	if hotelID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrHotelIDRequired})
		return
	}

	roomID := c.Param("roomId")
	if roomID == "" {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrRoomIDRequired})
		return
	}

	if err := h.service.DeleteRoom(c.Request.Context(), roomID); err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
