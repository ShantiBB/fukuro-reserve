package helper

import (
	"net/http"

	"github.com/go-chi/render"

	"auth/internal/http/dto/response"
	"auth/internal/http/lib/validation"
	"fukuro-reserve/pkg/utils/errs"
)

func ParseJSON(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := render.DecodeJSON(r.Body, v); err != nil {
		errMsg := response.ErrorResp(errs.InvalidJSON)
		SendError(w, r, http.StatusBadRequest, errMsg)
		return false
	}

	if errResp := validation.CheckErrors(v); errResp != nil {
		SendError(w, r, http.StatusBadRequest, errResp)
		return false
	}

	return true
}
