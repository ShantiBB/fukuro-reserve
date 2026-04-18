package smoke

import (
	"net/http"
	"testing"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func runHotelsSmoke(t *testing.T, env *fixtures.Env) {
	t.Helper()

	t.Run("20 create hotel", func(t *testing.T) {
		var resp dto.HotelResponse
		status, body := env.RequestJSON(
			http.MethodPost,
			"/api/v1/hotels",
			env.Data.OwnerAccess,
			dto.CreateHotelRequest{
				CountryCode: "jp",
				CitySlug:    "tokyo",
				Title:       env.Data.HotelTitle,
				OwnerId:     1,
				Description: "http client smoke hotel",
				Address:     "1 HTTP Street, Tokyo",
				Location:    &dto.LocationDTO{Latitude: 35.66, Longitude: 139.70},
			},
			&resp,
		)
		env.RequireStatus(status, http.StatusCreated, body)
		env.Data.HotelID = resp.Id
		env.Data.HotelSlug = resp.HotelSlug
	})

	t.Run("21 list hotels", func(t *testing.T) {
		var resp dto.HotelsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/hotels?country_code=jp&city_slug=tokyo&sort_by=title&page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("22 get hotel by slug", func(t *testing.T) {
		var resp dto.HotelResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/hotels/jp/tokyo/"+env.Data.HotelSlug,
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("23 update hotel", func(t *testing.T) {
		var resp dto.HotelResponse
		status, body := env.RequestJSON(
			http.MethodPut,
			"/api/v1/hotels/jp/tokyo/"+env.Data.HotelSlug,
			env.Data.OwnerAccess,
			dto.UpdateHotelRequest{
				Description: "updated by http client",
				Address:     "2 HTTP Street, Tokyo",
				Location:    &dto.LocationDTO{Latitude: 35.67, Longitude: 139.71},
			},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("24 rename hotel", func(t *testing.T) {
		var resp dto.HotelResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/hotels/jp/tokyo/"+env.Data.HotelSlug+"/title",
			env.Data.OwnerAccess,
			dto.UpdateHotelTitleRequest{Title: env.Data.HotelTitleRenamed},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.HotelSlug != "" {
			env.Data.HotelSlug = resp.HotelSlug
		}
	})

	t.Run("25 get renamed hotel", func(t *testing.T) {
		var resp dto.HotelResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/hotels/jp/tokyo/"+env.Data.HotelSlug,
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})
}
