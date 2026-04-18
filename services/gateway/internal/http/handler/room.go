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

// CreateRoom godoc
// @Summary Create a new room
// @Tags rooms
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body dto.CreateRoomRequest true "Create room request"
// @Success 201 {object} dto.RoomResponse
// @Router /rooms [post]
func (h *HotelHandler) CreateRoom(c *gin.Context) {
	var req dto.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}
	if err := validation.ValidateCreateRoomRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.CreateRoom(c.Request.Context(), req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetRooms godoc
// @Summary Get rooms
// @Tags rooms
// @Produce json
// @Security Bearer
// @Param countryCode query string true "Country code"
// @Param citySlug query string true "City slug"
// @Param hotelSlug query string true "Hotel slug"
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Success 200 {object} dto.RoomsResponse
// @Router /rooms [get]
func (h *HotelHandler) GetRooms(c *gin.Context) {
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

	resp, err := h.service.GetRooms(
		c.Request.Context(),
		request.FirstNonEmptyQuery(c, "countryCode", "country_code"),
		request.FirstNonEmptyQuery(c, "citySlug", "city_slug"),
		request.FirstNonEmptyQuery(c, "hotelSlug", "hotel_slug"),
		page,
		limit,
	)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRoom godoc
// @Summary Get room by ID
// @Tags rooms
// @Produce json
// @Security Bearer
// @Param roomId path string true "Room ID"
// @Success 200 {object} dto.RoomResponse
// @Router /rooms/{roomId} [get]
func (h *HotelHandler) GetRoom(c *gin.Context) {
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
// @Param roomId path string true "Room ID"
// @Param request body dto.UpdateRoomRequest true "Update room request"
// @Success 200 {object} dto.RoomResponse
// @Router /rooms/{roomId} [put]
func (h *HotelHandler) UpdateRoom(c *gin.Context) {
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
	if err := validation.ValidateUpdateRoomRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
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
// @Param roomId path string true "Room ID"
// @Param request body dto.UpdateRoomStatusRequest true "Update room status request"
// @Success 200 {object} dto.StatusResponse
// @Router /rooms/{roomId}/status [patch]
func (h *HotelHandler) UpdateRoomStatus(c *gin.Context) {
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
	if err := validation.ValidateUpdateRoomStatusRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
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
// @Param roomId path string true "Room ID"
// @Success 204
// @Router /rooms/{roomId} [delete]
func (h *HotelHandler) DeleteRoom(c *gin.Context) {
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
