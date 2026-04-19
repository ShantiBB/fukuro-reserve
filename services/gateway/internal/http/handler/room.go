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
// @Summary       Create a new room
// @Description   Creates a room in the selected hotel. Requires JWT auth.
// @Tags          rooms
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         request body dto.CreateRoomRequest true "Create room request"
// @Success       201 {object} dto.RoomResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/rooms [post]
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
// @Summary       Get rooms by hotel slug
// @Description   Public endpoint. Returns rooms for hotel resolved by slug.
// @Tags          rooms
// @Produce       json
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelSlug path string true "Hotel slug" example(imperial-hotel-tokyo)
// @Param         page query int false "Page number (starts from 1)" default(1) minimum(1)
// @Param         limit query int false "Page size" default(10) minimum(1) maximum(100)
// @Success       200 {object} dto.RoomsResponse
// @Failure       400 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/slug/{hotelSlug}/rooms [get]
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
// @Summary       Get rooms by hotel ID
// @Description   Public endpoint. Returns rooms for hotel resolved by hotel ID.
// @Tags          rooms
// @Produce       json
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         page query int false "Page number (starts from 1)" default(1) minimum(1)
// @Param         limit query int false "Page size" default(10) minimum(1) maximum(100)
// @Success       200 {object} dto.RoomsResponse
// @Failure       400 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/rooms [get]
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
// @Summary       Get room by hotel slug and room ID
// @Description   Public endpoint. Returns room details by hotel slug and room ID.
// @Tags          rooms
// @Produce       json
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelSlug path string true "Hotel slug" example(imperial-hotel-tokyo)
// @Param         roomId path string true "Room ID" example(1f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       200 {object} dto.RoomResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/slug/{hotelSlug}/rooms/{roomId} [get]
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
// @Summary       Get room by hotel ID and room ID
// @Description   Public endpoint. Returns room details by hotel ID and room ID.
// @Tags          rooms
// @Produce       json
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         roomId path string true "Room ID" example(1f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       200 {object} dto.RoomResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/rooms/{roomId} [get]
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
// @Summary       Update room
// @Description   Fully updates room fields. Requires JWT auth.
// @Tags          rooms
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         roomId path string true "Room ID" example(1f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         request body dto.UpdateRoomRequest true "Update room request"
// @Success       200 {object} dto.RoomResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/rooms/{roomId} [put]
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
// @Summary       Update room status
// @Description   Updates room lifecycle status. Requires JWT auth.
// @Tags          rooms
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         roomId path string true "Room ID" example(1f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         request body dto.UpdateRoomStatusRequest true "Update room status request"
// @Success       200 {object} dto.StatusResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/rooms/{roomId}/status [patch]
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
// @Summary       Delete room
// @Description   Deletes room by ID. Requires JWT auth.
// @Tags          rooms
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         roomId path string true "Room ID" example(1f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       204
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/rooms/{roomId} [delete]
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
