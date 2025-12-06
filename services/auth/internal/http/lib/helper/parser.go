package helper

import (
	"fmt"
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
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	if limit == "" {
		limit = "100"
	}
	if offset == "" {
		offset = "0"
	}

	limitUInt, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid limit: %w", err)
	}
	if limitUInt > MaxLimit {
		limitUInt = MaxLimit
	}

	offsetInt, err := strconv.ParseUint(offset, 10, 32)
	if err != nil {
		return nil, err
	}

	return &request.PaginationQuery{
		Limit:  limitUInt,
		Offset: offsetInt,
	}, nil
}
