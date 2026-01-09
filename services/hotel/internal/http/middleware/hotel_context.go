package middleware

import (
	"context"
	"net/http"

	"hotel/internal/http/utils/helper"
	"hotel/internal/repository/models"
)

type contextKey string

const HotelRefKey contextKey = "hotelRef"

func HotelPathValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hotelRef, err := helper.ParseHotelPathParams(r)
		if err != nil {
			helper.SendError(w, r, http.StatusBadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), HotelRefKey, hotelRef)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetHotelRef(ctx context.Context) models.HotelRef {
	return ctx.Value(HotelRefKey).(models.HotelRef)
}
