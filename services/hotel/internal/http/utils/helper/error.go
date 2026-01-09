package helper

import (
	"errors"
	"net/http"

	"hotel/internal/http/dto/response"
	"hotel/pkg/utils/consts"
)

type ErrorHandler struct {
	NotFoundError error
	ConflictError error
}

func (h *ErrorHandler) Handle(w http.ResponseWriter, r *http.Request, err error) error {
	if err == nil {
		return nil
	}

	if h.NotFoundError != nil && errors.Is(err, h.NotFoundError) {
		errMsg := response.ErrorResp(h.NotFoundError)
		SendError(w, r, http.StatusNotFound, errMsg)
		return err
	}

	if h.ConflictError != nil && errors.Is(err, h.ConflictError) {
		errMsg := response.ErrorResp(h.ConflictError)
		SendError(w, r, http.StatusConflict, errMsg)
		return err
	}

	errMsg := response.ErrorResp(consts.InternalServer)
	SendError(w, r, http.StatusInternalServerError, errMsg)
	return err
}
