package helper

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"hotel/internal/http/dto/request"
	"hotel/internal/http/utils/mapper"
	"hotel/internal/http/utils/validation"
	"hotel/internal/repository/models"
	"hotel/pkg/utils/consts"
)

func ParseJSON(
	w http.ResponseWriter, r *http.Request,
	v any,
	customErr func(validator.FieldError) string,
) error {
	if err := render.DecodeJSON(r.Body, v); err != nil {
		errMsg := validation.ErrorResp(consts.InvalidJSON)
		SendError(w, r, http.StatusBadRequest, errMsg)
		return err
	}

	if errMsg := validation.CheckErrors(v, customErr); errMsg != nil {
		SendError(w, r, http.StatusBadRequest, errMsg)
		return consts.InvalidJSON
	}

	return nil
}

func ParseHotelPathParams(r *http.Request) (models.HotelRef, *validation.ValidateError) {
	pathParams := request.HotelPathParams{
		CountryCode: chi.URLParam(r, "countryCode"),
		CitySlug:    chi.URLParam(r, "citySlug"),
		HotelSlug:   chi.URLParam(r, "hotelSlug"),
	}

	if errMsg := validation.CheckErrors(pathParams, validation.CustomValidationError); errMsg != nil {
		return models.HotelRef{}, errMsg
	}

	return mapper.HotelPathParamsToEntity(pathParams), nil
}

func ParseUUIDParam(r *http.Request, paramName string) (uuid.UUID, error) {
	paramID := chi.URLParam(r, paramName)
	id, err := uuid.Parse(paramID)
	if err != nil {
		return uuid.Nil, consts.InvalidID
	}
	return id, nil
}
