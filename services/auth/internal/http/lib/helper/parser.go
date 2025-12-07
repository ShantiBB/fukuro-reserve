package helper

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"auth/internal/http/dto/response"
	"auth/internal/http/lib/validation"
	"fukuro-reserve/pkg/utils/consts"
)

const MaxLimit = 100

func ParseJSON(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := render.DecodeJSON(r.Body, v); err != nil {
		errMsg := response.ErrorResp(consts.InvalidJSON)
		SendError(w, r, http.StatusBadRequest, errMsg)
		return false
	}

	if errMsg := validation.CheckErrors(v); errMsg != nil {
		SendError(w, r, http.StatusBadRequest, errMsg)
		return false
	}

	return true
}

func ParseID(w http.ResponseWriter, r *http.Request) int64 {
	paramID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		SendError(w, r, http.StatusBadRequest, errMsg)
		return 0
	}

	return id
}
