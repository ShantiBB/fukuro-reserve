package smoke

import (
	"fmt"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func bookingsBasePath(env *fixtures.Env) string {
	return fmt.Sprintf(
		"/api/v1/jp/tokyo/hotels/%s/bookings",
		env.Data.HotelID,
	)
}

func myBookingsBasePath() string {
	return "/api/v1/users/me/bookings"
}

func bookingsByIDPath(env *fixtures.Env, bookingID string) string {
	return bookingsBasePath(env) + "/" + bookingID
}

func quoteBookingPath(env *fixtures.Env) string {
	return bookingsBasePath(env) + "/quote"
}

func roomBookingsBasePath(env *fixtures.Env) string {
	return fmt.Sprintf(
		"/api/v1/jp/tokyo/hotels/%s/rooms/%s/bookings",
		env.Data.HotelID,
		env.Data.RoomID,
	)
}

func availabilityBasePath(env *fixtures.Env) string {
	return fmt.Sprintf(
		"/api/v1/jp/tokyo/hotels/%s/rooms/availability",
		env.Data.HotelID,
	)
}

func wrongLocationHotelPath(env *fixtures.Env) string {
	return fmt.Sprintf("/api/v1/us/osaka/hotels/%s", env.Data.HotelID)
}

func wrongLocationRoomPath(env *fixtures.Env) string {
	return fmt.Sprintf("/api/v1/us/osaka/hotels/%s/rooms/%s", env.Data.HotelID, env.Data.RoomID)
}

func wrongLocationBookingsBasePath(env *fixtures.Env) string {
	return fmt.Sprintf(
		"/api/v1/us/osaka/hotels/%s/bookings",
		env.Data.HotelID,
	)
}

func wrongLocationBookingsByIDPath(env *fixtures.Env, bookingID string) string {
	return wrongLocationBookingsBasePath(env) + "/" + bookingID
}

func wrongLocationRoomBookingsBasePath(env *fixtures.Env) string {
	return fmt.Sprintf(
		"/api/v1/us/osaka/hotels/%s/rooms/%s/bookings",
		env.Data.HotelID,
		env.Data.RoomID,
	)
}

func wrongLocationAvailabilityBasePath(env *fixtures.Env) string {
	return fmt.Sprintf(
		"/api/v1/us/osaka/hotels/%s/rooms/availability",
		env.Data.HotelID,
	)
}
