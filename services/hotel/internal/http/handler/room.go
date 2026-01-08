package handler

import (
	"context"
	"errors"
	"hotel/internal/http/mapper"
	"hotel/internal/repository/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"fukuro-reserve/pkg/utils/consts"
	"fukuro-reserve/pkg/utils/helper"
	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
)

type RoomService interface {
	RoomCreate(ctx context.Context, hotel models.HotelRef, room models.RoomCreate) (models.Room, error)
	RoomGetByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID) (models.Room, error)
	RoomGetAll(ctx context.Context, hotel models.HotelRef, limit, offset uint64) (models.RoomList, error)
	RoomUpdateByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID, room models.RoomUpdate) error
	RoomDeleteByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID) error
}

// RoomCreate   godoc
// @Summary      Create room
// @Description  Create a new room from admin or owner provider
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param		 country_code    path		string	true	"Country Code"
// @Param		 city_slug    	 path		string	true	"City Slug"
// @Param		 hotel_slug      path		string	true	"Hotel slug"
// @Param        request         body       request.RoomCreate  true  "Room data"
// @Success      201             {object}   response.Room
// @Failure      400             {object}   response.ErrorSchema
// @Failure      401             {object}   response.ErrorSchema
// @Failure      409             {object}   response.ErrorSchema
// @Failure      500             {object}   response.ErrorSchema
// @Security     Bearer
// @Router       /{country_code}/{city_slug}/hotels/{hotel_slug}/rooms/  [post]
func (h *Handler) RoomCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hotelRef := models.HotelRef{
		CountryCode: chi.URLParam(r, "countryCode"),
		CitySlug:    chi.URLParam(r, "citySlug"),
		HotelSlug:   chi.URLParam(r, "hotelSlug"),
	}

	var req request.RoomCreate
	if err := helper.ParseJSON(w, r, &req, nil); err != nil {
		return
	}

	newRoom := mapper.RoomCreateRequestToEntity(req)
	createdRoom, err := h.svc.RoomCreate(ctx, hotelRef, newRoom)
	if err != nil {
		if errors.Is(err, consts.UniqueRoomField) {
			errMsg := response.ErrorResp(consts.UniqueRoomField)
			helper.SendError(w, r, http.StatusConflict, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	roomResponse := mapper.RoomEntityToResponse(createdRoom)
	helper.SendSuccess(w, r, http.StatusCreated, roomResponse)
}

// RoomGetAll    godoc
//
// @Summary		 Get rooms
// @Description	 Get rooms from all users
// @Tags		 rooms
// @Accept		 json
// @Produce		 json
// @Param		 country_code       path		string	true	"Country Code"
// @Param		 city_slug    	    path		string	true	"City Slug"
// @Param		 hotel_slug         path		string	true	"Hotel slug"
// @Param	     page	            query		uint64	false	"Page"	default(1)
// @Param	     limit	            query		uint64	false	"Limit"	default(20)
// @Success		 200		        {object}	response.RoomList
// @Failure		 401		        {object}	response.ErrorSchema
// @Failure		 500		        {object}	response.ErrorSchema
// @Security	 Bearer
// @Router		 /{country_code}/{city_slug}/hotels/{hotel_slug}/rooms/ [get]
func (h *Handler) RoomGetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hotelRef := models.HotelRef{
		CountryCode: chi.URLParam(r, "countryCode"),
		CitySlug:    chi.URLParam(r, "citySlug"),
		HotelSlug:   chi.URLParam(r, "hotelSlug"),
	}

	pagination, err := helper.ParsePaginationQuery(r)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidQueryParam)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	roomList, err := h.svc.RoomGetAll(ctx, hotelRef, pagination.Page, pagination.Limit)
	if err != nil {
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	rooms := make([]response.RoomShort, 0, len(roomList.Rooms))
	for _, room := range roomList.Rooms {
		roomResponse := mapper.RoomShortEntityToShortResponse(room)
		rooms = append(rooms, roomResponse)
	}

	totalPageCount := (roomList.TotalCount + pagination.Limit - 1) / pagination.Limit
	pageLinks := helper.BuildPaginationLinks(r, pagination, totalPageCount)
	roomListResp := response.RoomList{
		Rooms:           rooms,
		CurrentPage:     pagination.Page,
		Limit:           pagination.Limit,
		Links:           pageLinks,
		TotalPageCount:  totalPageCount,
		TotalRoomsCount: roomList.TotalCount,
	}

	helper.SendSuccess(w, r, http.StatusOK, roomListResp)
}

// RoomGetByID    godoc
//
//	@Summary		Get room by ID
//	@Description	Get room by ID from all users
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Param		    country_code   path		string	true	"Country Code"
//	@Param		    city_slug      path		string	true	"City Slug"
//	@Param		    hotel_slug     path		string	true	"Hotel slug"
//	@Param			id	           path		string	true	"Room ID"
//	@Success		200	           {object}	response.Room
//	@Failure		400	           {object}	response.ErrorSchema
//	@Failure		401	           {object}	response.ErrorSchema
//	@Failure		404	           {object}	response.ErrorSchema
//	@Failure		500	           {object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/{country_code}/{city_slug}/hotels/{hotel_slug}/rooms/{id} [get]
func (h *Handler) RoomGetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hotelRef := models.HotelRef{
		CountryCode: chi.URLParam(r, "countryCode"),
		CitySlug:    chi.URLParam(r, "citySlug"),
		HotelSlug:   chi.URLParam(r, "hotelSlug"),
	}

	paramID := chi.URLParam(r, "id")
	id, err := uuid.Parse(paramID)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	room, err := h.svc.RoomGetByID(ctx, hotelRef, id)
	if err != nil {
		if errors.Is(err, consts.RoomNotFound) {
			errMsg := response.ErrorResp(consts.RoomNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	roomResponse := mapper.RoomEntityToResponse(room)
	helper.SendSuccess(w, r, http.StatusOK, roomResponse)
}

// RoomUpdateByID    godoc
//
//	@Summary		Update room by ID
//	@Description	Update room by ID from admin or owner provider
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Param		    country_code    path		string	true	"Country Code"
//	@Param		    city_slug       path		string	true	"City Slug"
//	@Param		    hotel_slug      path		string	true	"Hotel slug"
//	@Param			id	path		string	true	"Room ID"
//	@Param          request  body   request.RoomUpdate  true  "Room data"
//	@Success		200	{object}	response.RoomUpdate
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/{country_code}/{city_slug}/hotels/{hotel_slug}/rooms/{id} [put]
func (h *Handler) RoomUpdateByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hotelRef := models.HotelRef{
		CountryCode: chi.URLParam(r, "countryCode"),
		CitySlug:    chi.URLParam(r, "citySlug"),
		HotelSlug:   chi.URLParam(r, "hotelSlug"),
	}

	paramID := chi.URLParam(r, "id")
	id, err := uuid.Parse(paramID)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	var req request.RoomUpdate
	if err = helper.ParseJSON(w, r, &req, nil); err != nil {
		return
	}

	roomUpdate := mapper.RoomUpdateRequestToEntity(req)
	if err = h.svc.RoomUpdateByID(ctx, hotelRef, id, roomUpdate); err != nil {
		if errors.Is(err, consts.RoomNotFound) {
			errMsg := response.ErrorResp(consts.RoomNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	roomResponse := mapper.RoomUpdateEntityToResponse(roomUpdate)
	helper.SendSuccess(w, r, http.StatusOK, roomResponse)
}

// RoomDeleteByID    godoc
//
//	@Summary		Delete room by ID
//	@Description	Delete room by ID from admin or owner provider
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Param		    country_code    path		string	true	"Country Code"
//	@Param		    city_slug       path		string	true	"City Slug"
//	@Param		    hotel_slug      path		string	true	"Hotel slug"
//	@Param			id	path		string	true	"Room ID"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/{country_code}/{city_slug}/hotels/{hotel_slug}/rooms/{id} [delete]
func (h *Handler) RoomDeleteByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hotelRef := models.HotelRef{
		CountryCode: chi.URLParam(r, "countryCode"),
		CitySlug:    chi.URLParam(r, "citySlug"),
		HotelSlug:   chi.URLParam(r, "hotelSlug"),
	}

	paramID := chi.URLParam(r, "id")
	id, err := uuid.Parse(paramID)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	if err = h.svc.RoomDeleteByID(ctx, hotelRef, id); err != nil {
		if errors.Is(err, consts.RoomNotFound) {
			errMsg := response.ErrorResp(consts.RoomNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	helper.SendSuccess(w, r, http.StatusNoContent, nil)
}
