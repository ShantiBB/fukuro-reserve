package smoke

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func runBookingsSmoke(t *testing.T, env *fixtures.Env) {
	t.Helper()

	checkIn, err := time.Parse(time.RFC3339, "2026-06-01T12:00:00Z")
	if err != nil {
		t.Fatalf("parse check-in: %v", err)
	}
	checkOut, err := time.Parse(time.RFC3339, "2026-06-03T12:00:00Z")
	if err != nil {
		t.Fatalf("parse check-out: %v", err)
	}

	t.Run("40 create booking", func(t *testing.T) {
		var resp dto.BookingResponse
		status, body := env.RequestJSON(
			http.MethodPost,
			"/api/v1/bookings",
			env.Data.OwnerAccess,
			dto.CreateBookingRequest{
				UserId:              env.Data.OwnerID,
				HotelId:             env.Data.HotelID,
				CheckIn:             checkIn,
				CheckOut:            checkOut,
				GuestName:           "HTTP Booker",
				GuestEmail:          env.Data.OwnerEmail,
				GuestPhone:          "+79990000055",
				Currency:            "USD",
				ExpectedTotalAmount: "350.00",
				Rooms: []*dto.CreateBookingRoomRequest{
					{RoomId: env.Data.RoomID, Adults: 2, Children: 1, PricePerNight: "175.00"},
				},
			},
			&resp,
		)
		env.RequireStatus(status, http.StatusCreated, body)
		env.Data.BookingID = resp.Id
	})

	t.Run("41 list bookings all", func(t *testing.T) {
		var resp dto.BookingsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/bookings?page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("42 list bookings by user", func(t *testing.T) {
		var resp dto.BookingsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/bookings?userId="+strconv.FormatInt(env.Data.OwnerID, 10)+"&page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("43 list bookings by user and hotel", func(t *testing.T) {
		var resp dto.BookingsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/bookings?userId="+strconv.FormatInt(env.Data.OwnerID, 10)+"&hotelId="+env.Data.HotelID+"&page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("44 get booking", func(t *testing.T) {
		var resp dto.BookingResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			"/api/v1/bookings/"+env.Data.BookingID,
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("45 confirm booking", func(t *testing.T) {
		var resp dto.StatusResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/bookings/"+env.Data.BookingID+"/confirm",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})

	t.Run("46 cancel booking", func(t *testing.T) {
		var resp dto.StatusResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			"/api/v1/bookings/"+env.Data.BookingID+"/cancel",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
	})
}
