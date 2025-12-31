package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"fukuro-reserve/pkg/utils/consts"
	"fukuro-reserve/pkg/utils/helper"
	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
	"hotel/internal/repository/postgres/models"
)

type HotelService interface {
	HotelCreate(ctx context.Context, h models.HotelCreate) (models.Hotel, error)
	HotelGetByIDOrName(ctx context.Context, field any) (models.Hotel, error)
	HotelGetAll(ctx context.Context, limit, offset uint64) (models.HotelList, error)
	HotelUpdateByID(ctx context.Context, id int64, h models.HotelUpdate) error
	HotelDeleteByID(ctx context.Context, id int64) error
}

// HotelCreate   godoc
// @Summary      Create hotel
// @Description  Create a new hotel from admin provider
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param        request  body      request.HotelCreate  true  "Hotel data"
// @Success      201      {object}  response.Hotel
// @Failure      400      {object}  response.ErrorSchema
// @Failure      401      {object}  response.ErrorSchema
// @Failure      409      {object}  response.ErrorSchema
// @Failure      500      {object}  response.ErrorSchema
// @Security     Bearer
// @Router       /hotels/  [post]
func (h *Handler) HotelCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.HotelCreate

	if err := helper.ParseJSON(w, r, &req, nil); err != nil {
		return
	}

	newHotel := h.HotelCreateRequestToEntity(req)
	createdHotel, err := h.svc.HotelCreate(ctx, newHotel)
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

	userResponse := h.HotelResponseToEntity(createdHotel)
	helper.SendSuccess(w, r, http.StatusCreated, userResponse)
}

// HotelGetAll    godoc
//
//	@Summary		Get hotels
//	@Description	Get hotels from admin or moderator provider
//	@Tags			hotels
//	@Accept			json
//	@Produce		json
//	@Param			page	query		uint64	false	"Page"	default(1)
//	@Param			limit	query		uint64	false	"Limit"	default(100)
//	@Success		200		{object}	response.HotelList
//	@Failure		401		{object}	response.ErrorSchema
//	@Failure		500		{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/hotels/ [get]
func (h *Handler) HotelGetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pagination, err := helper.ParsePaginationQuery(r)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidQueryParam)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotelList, err := h.svc.HotelGetAll(ctx, pagination.Page, pagination.Limit)
	if err != nil {
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotels := make([]response.HotelShort, 0, len(hotelList.Hotels))
	for _, hotel := range hotelList.Hotels {
		hotelResponse := h.HotelShortResponseToEntity(hotel)
		hotels = append(hotels, hotelResponse)
	}

	totalPageCount := (hotelList.TotalCount + pagination.Limit - 1) / pagination.Limit
	pageLinks := helper.BuildPaginationLinks(r, pagination, totalPageCount)
	hotelListResp := response.HotelList{
		Hotels:          hotels,
		CurrentPage:     pagination.Page,
		Limit:           pagination.Limit,
		Links:           pageLinks,
		TotalPageCount:  totalPageCount,
		TotalUsersCount: hotelList.TotalCount,
	}

	helper.SendSuccess(w, r, http.StatusOK, hotelListResp)
}

// HotelGetByID    godoc
//
//	@Summary		Get hotel by ID
//	@Description	Get hotel by ID from admin, moderator or owner provider
//	@Tags			hotels
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Hotel ID"
//	@Success		200	{object}	response.Hotel
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/hotels/{id} [get]
func (h *Handler) HotelGetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := uuid.Parse(paramID)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	user, err := h.svc.HotelGetByIDOrName(ctx, id)
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

	userResponse := h.HotelResponseToEntity(user)
	helper.SendSuccess(w, r, http.StatusOK, userResponse)
}
