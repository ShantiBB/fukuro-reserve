package helper

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"auth/internal/http/dto/request"
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

	if errResp := validation.CheckErrors(v); errResp != nil {
		SendError(w, r, http.StatusBadRequest, errResp)
		return false
	}

	return true
}

func ParsePaginationQuery(r *http.Request) (*request.PaginationQuery, error) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	if limit == "" || limit == "0" {
		return nil, consts.InvalidQueryParam
	}
	if page == "" || page == "0" {
		return nil, consts.InvalidQueryParam
	}

	pageUInt, err := strconv.ParseUint(page, 10, 32)
	if err != nil {
		return nil, consts.InvalidQueryParam
	}

	limitUInt, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		return nil, consts.InvalidQueryParam
	}

	return &request.PaginationQuery{
		Page:  pageUInt,
		Limit: limitUInt,
	}, nil
}
