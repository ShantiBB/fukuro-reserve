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
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms",
			env.Data.OwnerAccess,
			dto.CreateRoomRequest{
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

	t.Run("30.1 create room as user forbidden", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodPost,
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms",
			env.Data.ManagedAccess,
			dto.CreateRoomRequest{},
			nil,
		)
		env.RequireError(status, http.StatusForbidden, body, "forbidden")
	})

	t.Run("30.2 create room anonymous unauthorized", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodPost,
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms",
			"",
			dto.CreateRoomRequest{},
			nil,
		)
		env.RequireError(status, http.StatusUnauthorized, body, "authorization header is required")
	})

	t.Run("30.3 create room as admin allowed by role", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodPost,
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms",
			env.Data.AdminAccess,
			dto.CreateRoomRequest{},
			nil,
		)
		env.RequireStatus(status, http.StatusBadRequest, body)
	})

	t.Run("31 list rooms by hotel id", func(t *testing.T) {
		var resp dto.RoomsResponse
		status, body := env.RequestJSON(http.MethodGet, "/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms?page=1&limit=10", env.Data.OwnerAccess, nil, &resp)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("31.1 list rooms by hotel id anonymous", func(t *testing.T) {
		var resp dto.RoomsResponse
		status, body := env.RequestJSON(http.MethodGet, "/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms?page=1&limit=10", "", nil, &resp)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("32 list rooms by hotel slug", func(t *testing.T) {
		var resp dto.RoomsResponse
		status, body := env.RequestJSON(http.MethodGet, "/api/v1/jp/tokyo/hotels/slug/"+env.Data.HotelSlug+"/rooms?page=1&limit=10", env.Data.OwnerAccess, nil, &resp)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("32.1 list rooms by hotel slug anonymous", func(t *testing.T) {
		var resp dto.RoomsResponse
		status, body := env.RequestJSON(http.MethodGet, "/api/v1/jp/tokyo/hotels/slug/"+env.Data.HotelSlug+"/rooms?page=1&limit=10", "", nil, &resp)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("33 get room by hotel slug", func(t *testing.T) {
		var resp dto.RoomResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/jp/tokyo/hotels/slug/"+env.Data.HotelSlug+"/rooms/"+env.Data.RoomID,
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("33.1 get room by hotel slug anonymous", func(t *testing.T) {
		var resp dto.RoomResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/jp/tokyo/hotels/slug/"+env.Data.HotelSlug+"/rooms/"+env.Data.RoomID,
			"",
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("34 get room by hotel id", func(t *testing.T) {
		var resp dto.RoomResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms/"+env.Data.RoomID,
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("34.1 get room by hotel id anonymous", func(t *testing.T) {
		var resp dto.RoomResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms/"+env.Data.RoomID,
			"",
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("35 update room", func(t *testing.T) {
		var resp dto.RoomResponse
		status, body := env.RequestJSON(
			http.MethodPut,
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms/"+env.Data.RoomID,
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

	t.Run("35.1 update room as user forbidden", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodPut,
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms/"+env.Data.RoomID,
			env.Data.ManagedAccess,
			dto.UpdateRoomRequest{},
			nil,
		)
		env.RequireError(status, http.StatusForbidden, body, "forbidden")
	})

	t.Run("36 update room status", func(t *testing.T) {
		var resp dto.StatusResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms/"+env.Data.RoomID+"/status",
			env.Data.OwnerAccess,
			dto.UpdateRoomStatusRequest{Status: "ROOM_STATUS_MAINTENANCE"},
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("36.1 update room status as user forbidden", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms/"+env.Data.RoomID+"/status",
			env.Data.ManagedAccess,
			dto.UpdateRoomStatusRequest{Status: "ROOM_STATUS_AVAILABLE"},
			nil,
		)
		env.RequireError(status, http.StatusForbidden, body, "forbidden")
	})

	t.Run("36.2 delete room as user forbidden", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodDelete,
			"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms/"+env.Data.RoomID,
			env.Data.ManagedAccess,
			nil,
			nil,
		)
		env.RequireError(status, http.StatusForbidden, body, "forbidden")
	})
}
