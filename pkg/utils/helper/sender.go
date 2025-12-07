package helper

import (
	"net/http"

	"github.com/go-chi/render"
)

func SendJSON(w http.ResponseWriter, r *http.Request, code int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	render.JSON(w, r, v)
}

func SendError(w http.ResponseWriter, r *http.Request, code int, v any) {
	SendJSON(w, r, code, v)
}

func SendSuccess(w http.ResponseWriter, r *http.Request, code int, v any) {
	SendJSON(w, r, code, v)
}
