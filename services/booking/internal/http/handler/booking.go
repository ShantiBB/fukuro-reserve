package handler

import (
	"booking/internal/http/utils/helper"
	"booking/internal/http/utils/validation"
	"context"
	"hotel/pkg/utils/consts"
	"net/http"

	"booking/internal/repository/models"

	"github.com/google/uuid"
)

type BookingService interface {
	BookingCreate(ctx context.Context, b models.CreateBooking) (models.Booking, error)
	BookingGetAll(
		ctx context.Context,
		bookingRef models.BookingRef,
		limit uint64,
		offset uint64,
	) (models.BookingList, error)
	BookingGetByID(ctx context.Context, id uuid.UUID) (models.Booking, error)
	BookingUpdateByID(ctx context.Context, id uuid.UUID, b models.UpdateBooking) error
	BookingStatusUpdateByID(ctx context.Context, id uuid.UUID, b models.BookingStatusInfo) error
	BookingDeleteByID(ctx context.Context, id uuid.UUID) error
}

// BookingCreate   godoc
// @Summary      Create booking
// @Description  Create a new booking from all users
// @Tags         bookings
// @Accept       json
// @Produce      json
// @Param		 country_code	path		string	true	"Country Code"
// @Param	  	 city_slug    	path		string	true	"City BookingSlug"
// @Param        request        body        request.BookingCreate  true  "Bookings data"
// @Success      201            {object}    response.Booking
// @Failure      400            {object}    response.ErrorSchema
// @Failure      401            {object}    response.ErrorSchema
// @Failure      409            {object}    response.ErrorSchema
// @Failure      500            {object}    response.ErrorSchema
// @Security     Bearer
// @Router       /bookings/  [post]
func (h *Handler) BookingCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	hotelRef := middleware.GetHotelRef(ctx)

	var req request.HotelCreate
	if err := helper.ParseJSON(w, r, &req, validation.CustomValidationError); err != nil {
		return
	}

	newHotel := mapper.HotelCreateRequestToEntity(req)
	createdHotel, err := h.svc.HotelCreate(ctx, hotelRef, newHotel)
	errHandler := &helper.ErrorHandler{Conflict: consts.UniqueHotelField}
	if err = errHandler.Handle(w, r, err); err != nil {
		return
	}

	hotelResponse := mapper.HotelCreateEntityToResponse(createdHotel)
	helper.SendSuccess(w, r, http.StatusCreated, hotelResponse)
}
