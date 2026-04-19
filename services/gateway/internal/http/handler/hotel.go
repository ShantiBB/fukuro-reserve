package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/request"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/responder"
)

// CreateHotel godoc
// @Summary       Create a new hotel
// @Description   Creates a hotel in the selected location. Requires JWT auth.
// @Tags          hotels
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         request body dto.CreateHotelBody true "Create hotel request"
// @Success       201 {object} dto.HotelResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels [post]
func (h *HotelHandler) CreateHotel(c *gin.Context) {
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

	var body dto.CreateHotelBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	req := dto.CreateHotelRequest{
		CountryCode: countryCode,
		CitySlug:    citySlug,
		Title:       body.Title,
		Description: body.Description,
		Address:     body.Address,
		OwnerId:     body.OwnerId,
		Location:    body.Location,
	}

	resp, err := h.service.CreateHotel(c.Request.Context(), req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetHotels godoc
// @Summary       Get hotels
// @Description   Public endpoint. Returns hotels in location with optional sorting and pagination.
// @Tags          hotels
// @Produce       json
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         sort_by query string false "Sort field" Enums(title,rating,created_at,updated_at) default(title)
// @Param         page query int false "Page number (starts from 1)" default(1) minimum(1)
// @Param         limit query int false "Page size" default(10) minimum(1) maximum(100)
// @Success       200 {object} dto.HotelsResponse
// @Failure       400 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels [get]
func (h *HotelHandler) GetHotels(c *gin.Context) {
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

	sortBy := request.FirstNonEmptyQuery(c, "sort_by")
	if sortBy == "" {
		sortBy = "title"
	}

	resp, err := h.service.GetHotels(
		c.Request.Context(),
		countryCode,
		citySlug,
		sortBy,
		page,
		limit,
	)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHotelByID godoc
// @Summary       Get hotel by ID
// @Description   Public endpoint. Returns hotel details by hotel ID.
// @Tags          hotels
// @Produce       json
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       200 {object} dto.HotelResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId} [get]
func (h *HotelHandler) GetHotelByID(c *gin.Context) {
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

	resp, err := h.service.GetHotelByID(c.Request.Context(), hotelID)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHotelBySlug godoc
// @Summary       Get hotel by slug
// @Description   Public endpoint. Returns hotel details by SEO-friendly hotel slug.
// @Tags          hotels
// @Produce       json
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelSlug path string true "Hotel slug" example(imperial-hotel-tokyo)
// @Success       200 {object} dto.HotelResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/slug/{hotelSlug} [get]
func (h *HotelHandler) GetHotelBySlug(c *gin.Context) {
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

	resp, err := h.service.GetHotelBySlug(c.Request.Context(), countryCode, citySlug, hotelSlug)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateHotel godoc
// @Summary       Update hotel
// @Description   Fully updates mutable hotel fields. Requires JWT auth.
// @Tags          hotels
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         request body dto.UpdateHotelRequest true "Update hotel request"
// @Success       200 {object} dto.HotelResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId} [put]
func (h *HotelHandler) UpdateHotel(c *gin.Context) {
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

	var req dto.UpdateHotelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.UpdateHotelByID(
		c.Request.Context(),
		hotelID,
		req,
	)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateHotelTitle godoc
// @Summary       Update hotel title
// @Description   Partially updates hotel title and slug. Requires JWT auth.
// @Tags          hotels
// @Accept        json
// @Produce       json
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Param         request body dto.UpdateHotelTitleRequest true "Update hotel title request"
// @Success       200 {object} dto.HotelResponse
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId}/title [patch]
func (h *HotelHandler) UpdateHotelTitle(c *gin.Context) {
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

	var req dto.UpdateHotelTitleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}

	resp, err := h.service.UpdateHotelTitleByID(
		c.Request.Context(),
		hotelID,
		req,
	)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteHotel godoc
// @Summary       Delete hotel
// @Description   Deletes hotel by ID. Requires JWT auth.
// @Tags          hotels
// @Security      Bearer
// @Param         countryCode path string true "Country code (ISO 3166-1 alpha-2)" example(jp)
// @Param         citySlug path string true "City slug" example(tokyo)
// @Param         hotelId path string true "Hotel ID" example(0f8fad5b-d9cb-469f-a165-70867728950e)
// @Success       204
// @Failure       400 {object} responder.ErrorResponse
// @Failure       401 {object} responder.ErrorResponse
// @Failure       404 {object} responder.ErrorResponse
// @Router        /{countryCode}/{citySlug}/hotels/{hotelId} [delete]
func (h *HotelHandler) DeleteHotel(c *gin.Context) {
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

	if err := h.service.DeleteHotelByID(c.Request.Context(), hotelID); err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
