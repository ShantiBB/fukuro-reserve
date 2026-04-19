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

func wrongLocationHotelPath(env *fixtures.Env) string {
	return fmt.Sprintf("/api/v1/us/osaka/hotels/%s", env.Data.HotelID)
}

func wrongLocationRoomPath(env *fixtures.Env) string {
	return fmt.Sprintf("/api/v1/us/osaka/hotels/%s/rooms/%s", env.Data.HotelID, env.Data.RoomID)
}

func wrongLocationBookingsBasePath(env *fixtures.Env) string {
	return fmt.Sprintf(
		"/api/v1/us/osaka/hotels/%s/rooms/%s/bookings",
		env.Data.HotelID,
		env.Data.RoomID,
	)
}

func wrongLocationBookingsByIDPath(env *fixtures.Env, bookingID string) string {
	return wrongLocationBookingsBasePath(env) + "/" + bookingID
}
