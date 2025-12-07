package handler

import (
	"context"
	"errors"
	"net/http"

	"fukuro-reserve/pkg/utils/consts"
	"fukuro-reserve/pkg/utils/helper"
	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
	"hotel/internal/repository/postgres/models"
)

type HotelService interface {
	HotelCreate(ctx context.Context, h models.HotelCreate) (models.Hotel, error)
	HotelGetByIDOrName(ctx context.Context, field any) (models.Hotel, error)
	HotelGetAll(ctx context.Context, limit, offset uint64) ([]models.HotelShort, error)
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

	if ok := helper.ParseJSON(w, r, &req, nil); !ok {
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
