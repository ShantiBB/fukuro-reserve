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
// @Summary Create a new hotel
// @Tags hotels
// @Accept json
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param request body dto.CreateHotelBody true "Create hotel request"
// @Success 201 {object} dto.HotelResponse
// @Router /{countryCode}/{citySlug}/hotels [post]
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
// @Summary Get hotels
// @Tags hotels
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param sort_by query string false "Sort field"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {object} dto.HotelsResponse
// @Router /{countryCode}/{citySlug}/hotels [get]
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

	sortBy := request.FirstNonEmptyQuery(c, "sortBy", "sort_by")
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

// GetHotelByID GetHotel godoc
// @Summary Get hotel by ID
// @Tags hotels
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelId path string true "Hotel ID"
// @Success 200 {object} dto.HotelResponse
// @Router /{countryCode}/{citySlug}/hotels/{hotelId} [get]
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
// @Summary Get hotel by slug
// @Tags hotels
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelSlug path string true "Hotel slug"
// @Success 200 {object} dto.HotelResponse
// @Router /{countryCode}/{citySlug}/hotels/slug/{hotelSlug} [get]
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
// @Summary Update hotel
// @Tags hotels
// @Accept json
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelId path string true "Hotel ID"
// @Param request body dto.UpdateHotelRequest true "Update hotel request"
// @Success 200 {object} dto.HotelResponse
// @Router /{countryCode}/{citySlug}/hotels/{hotelId} [put]
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
// @Summary Update hotel title
// @Tags hotels
// @Accept json
// @Produce json
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelId path string true "Hotel ID"
// @Param request body dto.UpdateHotelTitleRequest true "Update hotel title request"
// @Success 200 {object} dto.HotelResponse
// @Router /{countryCode}/{citySlug}/hotels/{hotelId}/title [patch]
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
// @Summary Delete hotel
// @Tags hotels
// @Security Bearer
// @Param countryCode path string true "Country code"
// @Param citySlug path string true "City slug"
// @Param hotelId path string true "Hotel ID"
// @Success 204
// @Router /{countryCode}/{citySlug}/hotels/{hotelId} [delete]
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
