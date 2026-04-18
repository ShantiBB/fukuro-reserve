package smoke

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func runCleanupSmoke(t *testing.T, env *fixtures.Env) {
	t.Helper()

	t.Run("90 delete booking", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodDelete,
			"/api/v1/bookings/"+env.Data.BookingID,
			env.Data.OwnerAccess,
			nil,
			nil,
		)
		env.RequireStatus(status, http.StatusNoContent, body)
	})

	t.Run("91 delete room", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodDelete,
			"/api/v1/rooms/"+env.Data.RoomID,
			env.Data.OwnerAccess,
			nil,
			nil,
		)
		env.RequireStatus(status, http.StatusNoContent, body)
	})

	t.Run("92 delete hotel", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodDelete,
			"/api/v1/hotels/jp/tokyo/"+env.Data.HotelSlug,
			env.Data.OwnerAccess,
			nil,
			nil,
		)
		env.RequireStatus(status, http.StatusNoContent, body)
	})

	t.Run("93 delete managed user as admin", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodDelete,
			"/api/v1/auth/users/"+strconv.FormatInt(env.Data.ManagedID, 10),
			env.Data.AdminAccess,
			nil,
			nil,
		)
		env.RequireStatus(status, http.StatusNoContent, body)
	})

	t.Run("94 delete owner self", func(t *testing.T) {
		status, body := env.RequestJSON(
			http.MethodDelete,
			"/api/v1/auth/users/"+strconv.FormatInt(env.Data.OwnerID, 10),
			env.Data.OwnerAccess,
			nil,
			nil,
		)
		env.RequireStatus(status, http.StatusNoContent, body)
	})
}
