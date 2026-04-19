package smoke

import (
	"net/http"
	"testing"
	"time"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func runBookingsSmoke(t *testing.T, env *fixtures.Env) {
	t.Helper()

	checkIn := time.Now().UTC().Add(14 * 24 * time.Hour).Truncate(time.Second)
	checkOut := checkIn.Add(48 * time.Hour)
	basePath := bookingsBasePath(env)

	t.Run("40 create booking", func(t *testing.T) {
		var resp dto.BookingResponse
		status, body := env.RequestJSON(
			http.MethodPost,
			basePath,
			env.Data.OwnerAccess,
			dto.CreateBookingRequest{
				UserId:              env.Data.OwnerID,
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
		assertBookingResponseFields(t, &resp, env)
		assertBookingStatusEnum(t, resp.Status)
		if resp.Status != "BOOKING_STATUS_PENDING" {
			t.Fatalf("unexpected booking status: got=%q want=%q", resp.Status, "BOOKING_STATUS_PENDING")
		}
		env.Data.BookingID = resp.Id
	})

	t.Run("41 list bookings all", func(t *testing.T) {
		var resp dto.BookingsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			basePath+"?page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if len(resp.Bookings) == 0 {
			t.Fatalf("bookings list is empty")
		}
		for _, booking := range resp.Bookings {
			assertBookingShortFields(t, booking)
			assertBookingStatusEnum(t, booking.Status)
		}
	})

	t.Run("42 list bookings by user", func(t *testing.T) {
		var resp dto.BookingsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			basePath+"?userId="+mustQueryUserID(env.Data.OwnerID)+"&page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if len(resp.Bookings) == 0 {
			t.Fatalf("bookings list by user is empty")
		}
		for _, booking := range resp.Bookings {
			assertBookingShortFields(t, booking)
			assertBookingStatusEnum(t, booking.Status)
			if booking.UserId != env.Data.OwnerID {
				t.Fatalf("unexpected booking.user_id: got=%d want=%d", booking.UserId, env.Data.OwnerID)
			}
		}
	})

	t.Run("43 list bookings by user and hotel", func(t *testing.T) {
		var resp dto.BookingsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			basePath+"?userId="+mustQueryUserID(env.Data.OwnerID)+"&page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if len(resp.Bookings) == 0 {
			t.Fatalf("bookings list by user+hotel is empty")
		}
		for _, booking := range resp.Bookings {
			assertBookingShortFields(t, booking)
			assertBookingStatusEnum(t, booking.Status)
			if booking.UserId != env.Data.OwnerID {
				t.Fatalf("unexpected booking.user_id: got=%d want=%d", booking.UserId, env.Data.OwnerID)
			}
			if booking.HotelId != env.Data.HotelID {
				t.Fatalf("unexpected booking.hotel_id: got=%q want=%q", booking.HotelId, env.Data.HotelID)
			}
		}
	})

	t.Run("44 get booking", func(t *testing.T) {
		var resp dto.BookingResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			bookingsByIDPath(env, env.Data.BookingID),
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		assertBookingResponseFields(t, &resp, env)
		assertBookingStatusEnum(t, resp.Status)
		if resp.Id != env.Data.BookingID {
			t.Fatalf("unexpected booking id: got=%q want=%q", resp.Id, env.Data.BookingID)
		}
		if resp.Status != "BOOKING_STATUS_PENDING" {
			t.Fatalf("unexpected booking status before confirm: got=%q want=%q", resp.Status, "BOOKING_STATUS_PENDING")
		}
	})

	t.Run("44.1 list bookings filtered by pending status", func(t *testing.T) {
		var resp dto.BookingsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			basePath+"?status=BOOKING_STATUS_PENDING&page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		booking := findBookingByIDInResponse(&resp, env.Data.BookingID)
		if booking == nil {
			t.Fatalf("pending booking not found in filtered response")
		}
		assertBookingShortFields(t, booking)
		if booking.Status != "BOOKING_STATUS_PENDING" {
			t.Fatalf("unexpected filtered booking status: got=%q want=%q", booking.Status, "BOOKING_STATUS_PENDING")
		}
	})

	t.Run("45 confirm booking", func(t *testing.T) {
		var resp dto.StatusResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			bookingsByIDPath(env, env.Data.BookingID)+"/confirm",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.Status != "BOOKING_STATUS_CONFIRMED" {
			t.Fatalf("unexpected confirm status: got=%q want=%q", resp.Status, "BOOKING_STATUS_CONFIRMED")
		}
	})

	t.Run("45.1 list bookings filtered by confirmed status", func(t *testing.T) {
		var resp dto.BookingsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			basePath+"?status=BOOKING_STATUS_CONFIRMED&page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		booking := findBookingByIDInResponse(&resp, env.Data.BookingID)
		if booking == nil {
			t.Fatalf("confirmed booking not found in filtered response")
		}
		assertBookingShortFields(t, booking)
		if booking.Status != "BOOKING_STATUS_CONFIRMED" {
			t.Fatalf("unexpected filtered booking status: got=%q want=%q", booking.Status, "BOOKING_STATUS_CONFIRMED")
		}
	})

	t.Run("46 cancel booking", func(t *testing.T) {
		var resp dto.StatusResponse
		status, body := env.RequestJSON(
			http.MethodPatch,
			bookingsByIDPath(env, env.Data.BookingID)+"/cancel",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		if resp.Status != "BOOKING_STATUS_CANCELLED" {
			t.Fatalf("unexpected cancel status: got=%q want=%q", resp.Status, "BOOKING_STATUS_CANCELLED")
		}
	})

	t.Run("47 get booking after cancel", func(t *testing.T) {
		var resp dto.BookingResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			bookingsByIDPath(env, env.Data.BookingID),
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		assertBookingResponseFields(t, &resp, env)
		assertBookingStatusEnum(t, resp.Status)
		if resp.Status != "BOOKING_STATUS_CANCELLED" {
			t.Fatalf("unexpected booking status after cancel: got=%q want=%q", resp.Status, "BOOKING_STATUS_CANCELLED")
		}
	})

	t.Run("48 list bookings filtered by cancelled status", func(t *testing.T) {
		var resp dto.BookingsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			basePath+"?status=BOOKING_STATUS_CANCELLED&page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		booking := findBookingByIDInResponse(&resp, env.Data.BookingID)
		if booking == nil {
			t.Fatalf("cancelled booking not found in filtered response")
		}
		assertBookingShortFields(t, booking)
		if booking.Status != "BOOKING_STATUS_CANCELLED" {
			t.Fatalf("unexpected filtered booking status: got=%q want=%q", booking.Status, "BOOKING_STATUS_CANCELLED")
		}
	})

	t.Run("49 list bookings with unknown status defaults to unspecified filter", func(t *testing.T) {
		var resp dto.BookingsResponse
		status, body := env.RequestJSON(
			http.MethodGet,
			basePath+"?status=BOOKING_STATUS_UNKNOWN&page=1&limit=10",
			env.Data.OwnerAccess,
			nil,
			&resp,
		)
		env.RequireStatus(status, http.StatusOK, body)
		booking := findBookingByIDInResponse(&resp, env.Data.BookingID)
		if booking == nil {
			t.Fatalf("booking not found when unknown status query is used")
		}
		assertBookingShortFields(t, booking)
		assertBookingStatusEnum(t, booking.Status)
	})
}
