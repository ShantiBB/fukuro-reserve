package smoke

import (
	"fmt"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func bookingsBasePath(env *fixtures.Env) string {
	return fmt.Sprintf(
		"/api/v1/jp/tokyo/hotels/%s/rooms/%s/bookings",
		env.Data.HotelID,
		env.Data.RoomID,
	)
}

func bookingsByIDPath(env *fixtures.Env, bookingID string) string {
	return bookingsBasePath(env) + "/" + bookingID
}
