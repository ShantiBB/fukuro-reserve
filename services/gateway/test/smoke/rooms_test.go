package smoke

import (
	"net/http"
	"testing"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func runRoomsSmoke(t *testing.T, env *fixtures.Env) {
	t.Helper()

	t.Run("30 create room", func(t *testing.T) {
		var resp dto.RoomResponse
		status, body := env.RequestJSON(
			http.MethodPost,
			"/api/v1/rooms",
			env.Data.OwnerAccess,
			dto.CreateRoomRequest{
				CountryCode: "jp",
				CitySlug:    "tokyo",
				HotelSlug:   env.Data.HotelSlug,
				Title:       "HTTP Room 101",
				Description: "http client room",
				RoomNumber:  "101",
				Type:        "ROOM_TYPE_SINGLE",
				Price:       "150.00",
				Capacity:    2,
				AreaSqm:     18.5,
				Floor:       1,
				Amenities:   []string{"wifi"},
				Images:      []string{"https://example.com/http-room.jpg"},
			},
			&resp,
		)
		env.RequireStatus(status, http.StatusCreated, body)
		env.Data.RoomID = resp.Id
	})

	t.Run("31 list rooms", func(t *testing.T) {
		var resp dto.RoomsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/rooms?country_code=jp&city_slug=tokyo&hotel_slug="+env.Data.HotelSlug+"&page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("32 get room", func(t *testing.T) {
		var resp dto.RoomResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/rooms/"+env.Data.RoomID,
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("33 update room", func(t *testing.T) {
		var resp dto.RoomResponse
		status, body := env.RequestJSON(
			http.MethodPut,
			"/api/v1/rooms/"+env.Data.RoomID,
			env.Data.OwnerAccess,
			dto.UpdateRoomRequest{
				Title:       "HTTP Room 101 Updated",
				RoomNumber:  "101",
				Type:        "ROOM_TYPE_DOUBLE",
				Description: "updated by http client",
				Price:       "175.00",
				Capacity:    3,
				AreaSqm:     20.0,
				Floor:       1,
				Amenities:   []string{"wifi", "tv"},
				Images:      []string{"https://example.com/http-room-updated.jpg"},
			},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("34 update room status", func(t *testing.T) {
		var resp dto.StatusResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/rooms/"+env.Data.RoomID+"/status",
			env.Data.OwnerAccess,
			dto.UpdateRoomStatusRequest{Status: "ROOM_STATUS_MAINTENANCE"},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})
}
