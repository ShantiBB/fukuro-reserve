package helper

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"fukuro-reserve/pkg/utils/consts"
	"fukuro-reserve/pkg/utils/validation"
)

func ParseJSON(
	w http.ResponseWriter, r *http.Request,
	v any,
	customErr func(validator.FieldError) string,
) bool {
	if err := render.DecodeJSON(r.Body, v); err != nil {
		errMsg := validation.ErrorResp(consts.InvalidJSON)
		SendError(w, r, http.StatusBadRequest, errMsg)
		return false
	}

	if errMsg := validation.CheckErrors(v, customErr); errMsg != nil {
		SendError(w, r, http.StatusBadRequest, errMsg)
		return false
	}

	return true
}

func ParseID(w http.ResponseWriter, r *http.Request) int64 {
	paramID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		errMsg := validation.ErrorResp(consts.InvalidID)
		SendError(w, r, http.StatusBadRequest, errMsg)
		return 0
	}

	return id
}
