package handler

import (
	"context"
	"errors"
	"hotel/internal/http/mapper"
	"hotel/internal/repository/models"
	"net/http"

	"fukuro-reserve/pkg/utils/consts"
	"fukuro-reserve/pkg/utils/helper"
	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"

	"github.com/go-chi/chi/v5"
)

type HotelService interface {
	HotelCreate(ctx context.Context, countryCode, citySlug string, h models.HotelCreate) (models.Hotel, error)
	HotelGetBySlug(ctx context.Context, countryCode, citySlug, slug string) (models.Hotel, error)
	HotelGetAll(ctx context.Context, countryCode, citySlug, sortField string, page, limit uint64) (models.HotelList, error)
	HotelUpdateBySlug(ctx context.Context, countryCode, citySlug, hotelSlug string, h models.HotelUpdate) error
	HotelDeleteBySlug(ctx context.Context, countryCode, citySlug, hotelSlug string) error
}

// HotelCreate   godoc
// @Summary      Create hotel
// @Description  Create a new hotel from admin provider
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param		 country_code	path		string	true	"Country Code"
// @Param	  	 city_slug    	path		string	true	"City Slug"
// @Param        request        body        request.HotelCreate  true  "Hotel data"
// @Success      201            {object}    response.Hotel
// @Failure      400            {object}    response.ErrorSchema
// @Failure      401            {object}    response.ErrorSchema
// @Failure      409            {object}    response.ErrorSchema
// @Failure      500            {object}    response.ErrorSchema
// @Security     Bearer
// @Router       /{country_code}/{city_slug}/hotels/  [post]
func (h *Handler) HotelCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.HotelCreate

	if err := helper.ParseJSON(w, r, &req, nil); err != nil {
		return
	}

	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	newHotel := mapper.HotelCreateRequestToEntity(req)

	createdHotel, err := h.svc.HotelCreate(ctx, countryCode, citySlug, newHotel)
	if err != nil {
		if errors.Is(err, consts.UniqueHotelField) {
			errMsg := response.ErrorResp(consts.UniqueHotelField)
			helper.SendError(w, r, http.StatusConflict, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotelResponse := mapper.HotelCreateEntityToResponse(createdHotel)
	helper.SendSuccess(w, r, http.StatusCreated, hotelResponse)
}

// HotelGetAll    godoc
//
//	@Summary		Get hotels
//	@Description	Get hotels from admin or moderator provider
//	@Tags			hotels
//	@Accept			json
//	@Produce		json
//	@Param			country_code	path		string	true	"Country Code"
//	@Param			city_slug    	path		string	true	"City Slug"
//	@Param			page	        query		uint64	false	"Page"	default(1)
//	@Param			limit	        query		uint64	false	"Limit"	default(20)
//	@Success		200		        {object}	response.HotelList
//	@Failure		401		        {object}	response.ErrorSchema
//	@Failure		500		        {object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/{country_code}/{city_slug}/hotels/ [get]
func (h *Handler) HotelGetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	sortField := chi.URLParam(r, "sort")

	pagination, err := helper.ParsePaginationQuery(r)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidQueryParam)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotelList, err := h.svc.HotelGetAll(ctx, countryCode, citySlug, sortField, pagination.Page, pagination.Limit)
	if err != nil {
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotels := make([]response.HotelShort, 0, len(hotelList.Hotels))
	for _, hotel := range hotelList.Hotels {
		hotelResponse := mapper.HotelShortEntityToShortResponse(hotel)
		hotels = append(hotels, hotelResponse)
	}

	totalPageCount := (hotelList.TotalCount + pagination.Limit - 1) / pagination.Limit
	pageLinks := helper.BuildPaginationLinks(r, pagination, totalPageCount)
	hotelListResp := response.HotelList{
		Hotels:           hotels,
		CurrentPage:      pagination.Page,
		Limit:            pagination.Limit,
		Links:            pageLinks,
		TotalPageCount:   totalPageCount,
		TotalHotelsCount: hotelList.TotalCount,
	}

	helper.SendSuccess(w, r, http.StatusOK, hotelListResp)
}

// HotelGetBySlug    godoc
//
//	@Summary		Get hotel by slug
//	@Description	Get hotel by slug from admin, moderator or owner provider
//	@Tags			hotels
//	@Accept			json
//	@Produce		json
//	@Param			country_code	path		         string	true	"Country Code"
//	@Param			city_slug    	path		         string	true	"City Slug"
//	@Param			hotel_slug      path		         string	true	"Hotel slug"
//	@Success		200	{object}	response.Hotel
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/{country_code}/{city_slug}/hotels/{hotel_slug} [get]
func (h *Handler) HotelGetBySlug(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	slug := chi.URLParam(r, "hotelSlug")

	hotel, err := h.svc.HotelGetBySlug(ctx, countryCode, citySlug, slug)
	if err != nil {
		if errors.Is(err, consts.HotelNotFound) {
			errMsg := response.ErrorResp(consts.HotelNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotelResponse := mapper.HotelGetEntityToResponse(hotel)
	helper.SendSuccess(w, r, http.StatusOK, hotelResponse)
}

// HotelUpdateBySlug    godoc
//
//	@Summary		Update hotel by slug
//	@Description	Update hotel by slug from admin, moderator or owner provider
//	@Tags			hotels
//	@Accept			json
//	@Produce		json
//	@Param			country_code	path		string	true	"Country Code"
//	@Param			city_slug    	path		string	true	"City Slug"
//	@Param			hotel_slug	    path		string	true	"Hotel slug"
//	@Param          request         body        request.HotelUpdate  true  "Hotel data"
//	@Success		200	{object}	            response.HotelUpdate
//	@Failure		400	{object}	            response.ErrorSchema
//	@Failure		401	{object}	            response.ErrorSchema
//	@Failure		404	{object}	            response.ErrorSchema
//	@Failure		500	{object}	            response.ErrorSchema
//	@Security		Bearer
//	@Router			/{country_code}/{city_slug}/hotels/{hotel_slug} [put]
func (h *Handler) HotelUpdateBySlug(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	slug := chi.URLParam(r, "hotelSlug")

	var req request.HotelUpdate
	if err := helper.ParseJSON(w, r, &req, nil); err != nil {
		return
	}

	hotelUpdate := mapper.HotelUpdateRequestToEntity(req)
	if err := h.svc.HotelUpdateBySlug(ctx, countryCode, citySlug, slug, hotelUpdate); err != nil {
		if errors.Is(err, consts.HotelNotFound) {
			errMsg := response.ErrorResp(consts.HotelNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotelResponse := mapper.HotelUpdateEntityToResponse(hotelUpdate)
	helper.SendSuccess(w, r, http.StatusOK, hotelResponse)
}

// HotelDeleteBySlug    godoc
//
//	@Summary		Delete hotel by slug
//	@Description	Delete hotel by slug from admin or owner provider
//	@Tags			hotels
//	@Accept			json
//	@Produce		json
//	@Param			country_code	path		string	true	"Country Code"
//	@Param			city_slug    	path		string	true	"City Slug"
//	@Param			hotel_slug      path		string	true	"Hotel slug"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/{country_code}/{city_slug}/hotels/{hotel_slug} [delete]
func (h *Handler) HotelDeleteBySlug(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	countryCode := chi.URLParam(r, "countryCode")
	citySlug := chi.URLParam(r, "citySlug")
	slug := chi.URLParam(r, "hotelSlug")

	if err := h.svc.HotelDeleteBySlug(ctx, countryCode, citySlug, slug); err != nil {
		if errors.Is(err, consts.HotelNotFound) {
			errMsg := response.ErrorResp(consts.HotelNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	helper.SendSuccess(w, r, http.StatusNoContent, nil)
}
