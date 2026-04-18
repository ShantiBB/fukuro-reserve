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

// CreateHotel godoc
// @Summary Create a new hotel
// @Tags hotels
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body dto.CreateHotelRequest true "Create hotel request"
// @Success 201 {object} dto.HotelResponse
// @Router /hotels [post]
func (h *HotelHandler) CreateHotel(c *gin.Context) {
	var req dto.CreateHotelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}
	if err := validation.ValidateCreateHotelRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.CreateHotel(c.Request.Context(), req)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetHotels godoc
// @Summary Get hotels
// @Tags hotels
// @Produce json
// @Security Bearer
// @Param country_code query string true "Country code"
// @Param city_slug query string true "City slug"
// @Param sort_by query string false "Sort field"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {object} dto.HotelsResponse
// @Router /hotels [get]
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

	sortBy := request.FirstNonEmptyQuery(c, "sortBy", "sort_by")
	if sortBy == "" {
		sortBy = "title"
	}

	resp, err := h.service.GetHotels(
		c.Request.Context(),
		request.FirstNonEmptyQuery(c, "countryCode", "country_code"),
		request.FirstNonEmptyQuery(c, "citySlug", "city_slug"),
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

// GetHotel godoc
// @Summary Get hotel by slug
// @Tags hotels
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Success 200 {object} dto.HotelResponse
// @Router /hotels/{countryCode}/{citySlug}/{hotelSlug} [get]
func (h *HotelHandler) GetHotel(c *gin.Context) {
	resp, err := h.service.GetHotel(
		c.Request.Context(),
		c.Param("countryCode"),
		c.Param("citySlug"),
		c.Param("hotelSlug"),
	)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateHotel godoc
// @Summary Update hotel
// @Tags hotels
// @Accept json
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Param request body dto.UpdateHotelRequest true "Update hotel request"
// @Success 200 {object} dto.HotelResponse
// @Router /hotels/{countryCode}/{citySlug}/{hotelSlug} [put]
func (h *HotelHandler) UpdateHotel(c *gin.Context) {
	var req dto.UpdateHotelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}
	if err := validation.ValidateUpdateHotelRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.UpdateHotel(
		c.Request.Context(),
		c.Param("countryCode"),
		c.Param("citySlug"),
		c.Param("hotelSlug"),
		req,
	)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateHotelTitle godoc
// @Summary Update hotel title
// @Tags hotels
// @Accept json
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Param request body dto.UpdateHotelTitleRequest true "Update hotel title request"
// @Success 200 {object} dto.HotelResponse
// @Router /hotels/{countryCode}/{citySlug}/{hotelSlug}/title [patch]
func (h *HotelHandler) UpdateHotelTitle(c *gin.Context) {
	var req dto.UpdateHotelTitleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: consts.ErrInvalidRequestBody})
		return
	}
	if err := validation.ValidateUpdateHotelTitleRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, &responder.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.UpdateHotelTitle(
		c.Request.Context(),
		c.Param("countryCode"),
		c.Param("citySlug"),
		c.Param("hotelSlug"),
		req,
	)
	if err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteHotel godoc
// @Summary Delete hotel
// @Tags hotels
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Success 204
// @Router /hotels/{countryCode}/{citySlug}/{hotelSlug} [delete]
func (h *HotelHandler) DeleteHotel(c *gin.Context) {
	if err := h.service.DeleteHotel(
		c.Request.Context(),
		c.Param("countryCode"),
		c.Param("citySlug"),
		c.Param("hotelSlug"),
	); err != nil {
		responder.GinGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
