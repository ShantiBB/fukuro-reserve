package handler

import (
	"context"
	"net/http"

	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
	"hotel/internal/http/middleware"
	"hotel/internal/http/utils/helper"
	"hotel/internal/http/utils/mapper"
	"hotel/internal/http/utils/pagination"
	"hotel/internal/http/utils/validation"
	"hotel/internal/repository/models"
	"hotel/pkg/utils/consts"

	"github.com/google/uuid"
)

type RoomService interface {
	RoomCreate(ctx context.Context, hotel models.HotelRef, room models.RoomCreate) (models.Room, error)
	RoomGetByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID) (models.Room, error)
	RoomGetAll(ctx context.Context, hotel models.HotelRef, limit, offset uint64) (models.RoomList, error)
	RoomUpdateByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID, room models.RoomUpdate) error
	RoomStatusUpdateByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID, room models.RoomStatusUpdate) error
	RoomDeleteByID(ctx context.Context, hotel models.HotelRef, roomID uuid.UUID) error
}

// RoomCreate   godoc
// @Summary      Create room
// @Description  Create a new room from admin or owner provider
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param		 country_code    path		string	true	"Country Code"
// @Param		 city_slug    	 path		string	true	"City HotelSlug"
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
	hotelRef := middleware.GetHotelRef(ctx)

	var req request.RoomCreate
	if err := helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	createdRoom, err := h.svc.RoomCreate(ctx, hotelRef, mapper.RoomCreateRequestToEntity(req))
	errHandler := &helper.ErrorHandler{
		ConflictError: consts.UniqueRoomField,
	}
	if err = errHandler.Handle(w, r, err); err != nil {
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
// @Param		 city_slug    	    path		string	true	"City HotelSlug"
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
	hotelRef := middleware.GetHotelRef(ctx)

	paginationParams, err := pagination.ParsePaginationQuery(r)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidQueryParam)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	roomList, err := h.svc.RoomGetAll(ctx, hotelRef, paginationParams.Page, paginationParams.Limit)
	errHandler := &helper.ErrorHandler{}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	rooms := make([]response.RoomShort, 0, len(roomList.Rooms))
	for _, room := range roomList.Rooms {
		rooms = append(rooms, mapper.RoomShortEntityToShortResponse(room))
	}

	totalPageCount := (roomList.TotalCount + paginationParams.Limit - 1) / paginationParams.Limit
	pageLinks := pagination.BuildPaginationLinks(r, paginationParams, totalPageCount)
	roomListResp := response.RoomList{
		Rooms:           rooms,
		CurrentPage:     paginationParams.Page,
		Limit:           paginationParams.Limit,
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
//	@Param		    city_slug      path		string	true	"City HotelSlug"
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
	hotelRef := middleware.GetHotelRef(ctx)

	id, err := helper.ParseUUIDParam(r, "id")
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	room, err := h.svc.RoomGetByID(ctx, hotelRef, id)
	errHandler := &helper.ErrorHandler{NotFoundError: consts.RoomNotFound}
	if err = errHandler.Handle(w, r, err); err != nil {
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
//	@Param		    city_slug       path		string	true	"City HotelSlug"
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
	hotelRef := middleware.GetHotelRef(ctx)

	id, err := helper.ParseUUIDParam(r, "id")
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	var req request.RoomUpdate
	if err = helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	roomUpdate := mapper.RoomUpdateRequestToEntity(req)
	err = h.svc.RoomUpdateByID(ctx, hotelRef, id, roomUpdate)
	errHandler := &helper.ErrorHandler{NotFoundError: consts.RoomNotFound}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	roomResponse := mapper.RoomUpdateEntityToResponse(roomUpdate)
	helper.SendSuccess(w, r, http.StatusOK, roomResponse)
}

// RoomStatusUpdateByID    godoc
//
//	@Summary		Update room status by ID
//	@Description	Update room status by ID from admin or owner provider
//	@Tags			rooms
//	@Accept			json
//	@Produce		json
//	@Param		    country_code    path		string	true	"Country Code"
//	@Param		    city_slug       path		string	true	"City HotelSlug"
//	@Param		    hotel_slug      path		string	true	"Hotel slug"
//	@Param			id	path		string	true	"Room ID"
//	@Param          request  body   request.RoomStatusUpdate  true  "Room data"
//	@Success		200	{object}	response.RoomStatusUpdate
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/{country_code}/{city_slug}/hotels/{hotel_slug}/rooms/{id}/update_status [put]
func (h *Handler) RoomStatusUpdateByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	hotelRef := middleware.GetHotelRef(ctx)

	id, err := helper.ParseUUIDParam(r, "id")
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	var req request.RoomStatusUpdate
	if err = helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	roomUpdate := mapper.RoomStatusUpdateRequestToEntity(req)
	err = h.svc.RoomStatusUpdateByID(ctx, hotelRef, id, roomUpdate)
	errHandler := &helper.ErrorHandler{NotFoundError: consts.RoomNotFound}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	roomResponse := mapper.RoomStatusUpdateEntityToResponse(roomUpdate)
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
//	@Param		    city_slug       path		string	true	"City HotelSlug"
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
	hotelRef := middleware.GetHotelRef(ctx)

	id, err := helper.ParseUUIDParam(r, "id")
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	err = h.svc.RoomDeleteByID(ctx, hotelRef, id)
	errHandler := &helper.ErrorHandler{NotFoundError: consts.RoomNotFound}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	helper.SendSuccess(w, r, http.StatusNoContent, nil)
}
