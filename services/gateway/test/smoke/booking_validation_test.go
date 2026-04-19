package smoke

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/test/smoke/fixtures"
)

func runBookingValidationSmoke(t *testing.T, env *fixtures.Env) {
	t.Helper()

	t.Run("booking grpc validation matrix", func(t *testing.T) {
		createPath := bookingsBasePath(env)

		checkIn := time.Now().UTC().Add(72 * time.Hour).Truncate(time.Second)
		checkOut := checkIn.Add(48 * time.Hour)
		validBody := map[string]any{
			"user_id":               env.Data.OwnerID,
			"check_in":              checkIn.Format(time.RFC3339),
			"check_out":             checkOut.Format(time.RFC3339),
			"guest_name":            "Booking Validation",
			"guest_email":           env.Data.OwnerEmail,
			"guest_phone":           "+79990000001",
			"currency":              "USD",
			"expected_total_amount": "300.00",
			"rooms": []map[string]any{
				{
					"room_id":         env.Data.RoomID,
					"adults":          2,
					"children":        1,
					"price_per_night": "150.00",
				},
			},
		}
		validQuoteBody := map[string]any{
			"check_in":  checkIn.Format(time.RFC3339),
			"check_out": checkOut.Format(time.RFC3339),
			"currency":  "USD",
			"rooms": []map[string]any{
				{
					"room_id":         env.Data.RoomID,
					"adults":          2,
					"children":        1,
					"price_per_night": "150.00",
				},
			},
		}

		t.Run("create booking user_id gt 0", func(t *testing.T) {
			body := cloneMap(validBody)
			body["user_id"] = int64(0)
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "user_id")
		})

		t.Run("create booking hotel_id uuid from path", func(t *testing.T) {
			assertValidationFields(t, env,
				http.MethodPost,
				"/api/v1/jp/tokyo/hotels/not-a-uuid/bookings",
				env.Data.OwnerAccess,
				validBody,
				"hotel_id",
			)
		})

		t.Run("create booking check_in gt_now", func(t *testing.T) {
			body := cloneMap(validBody)
			body["check_in"] = time.Now().UTC().Add(-48 * time.Hour).Format(time.RFC3339)
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "check_in")
		})

		t.Run("create booking check_out after check_in cel", func(t *testing.T) {
			body := cloneMap(validBody)
			body["check_out"] = checkIn.Add(-time.Hour).Format(time.RFC3339)
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body)
		})

		t.Run("create booking guest_name min len", func(t *testing.T) {
			body := cloneMap(validBody)
			body["guest_name"] = ""
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "guest_name")
		})

		t.Run("create booking guest_email format", func(t *testing.T) {
			body := cloneMap(validBody)
			body["guest_email"] = "invalid-email"
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "guest_email")
		})

		t.Run("create booking guest_phone min len", func(t *testing.T) {
			body := cloneMap(validBody)
			body["guest_phone"] = "1234"
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "guest_phone")
		})

		t.Run("create booking guest_phone max len", func(t *testing.T) {
			body := cloneMap(validBody)
			body["guest_phone"] = "+799900000011223344556677889900001"
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "guest_phone")
		})

		t.Run("create booking currency pattern", func(t *testing.T) {
			body := cloneMap(validBody)
			body["currency"] = "usd"
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "currency")
		})

		t.Run("create booking expected_total_amount required", func(t *testing.T) {
			body := cloneMap(validBody)
			body["expected_total_amount"] = ""
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "expected_total_amount")
		})

		t.Run("create booking expected_total_amount pattern", func(t *testing.T) {
			body := cloneMap(validBody)
			body["expected_total_amount"] = "10,50"
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "expected_total_amount")
		})

		t.Run("create booking rooms min_items", func(t *testing.T) {
			body := cloneMap(validBody)
			body["rooms"] = []map[string]any{}
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "rooms")
		})

		t.Run("create booking room_id uuid", func(t *testing.T) {
			body := cloneMap(validBody)
			body["rooms"] = []map[string]any{
				{
					"room_id":         "not-a-uuid",
					"adults":          2,
					"children":        1,
					"price_per_night": "150.00",
				},
			}
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "room_id")
		})

		t.Run("create booking adults gte 1", func(t *testing.T) {
			body := cloneMap(validBody)
			body["rooms"] = []map[string]any{
				{
					"room_id":         env.Data.RoomID,
					"adults":          0,
					"children":        1,
					"price_per_night": "150.00",
				},
			}
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "adults")
		})

		t.Run("create booking adults lte 20", func(t *testing.T) {
			body := cloneMap(validBody)
			body["rooms"] = []map[string]any{
				{
					"room_id":         env.Data.RoomID,
					"adults":          21,
					"children":        1,
					"price_per_night": "150.00",
				},
			}
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "adults")
		})

		t.Run("create booking children lte 20", func(t *testing.T) {
			body := cloneMap(validBody)
			body["rooms"] = []map[string]any{
				{
					"room_id":         env.Data.RoomID,
					"adults":          2,
					"children":        21,
					"price_per_night": "150.00",
				},
			}
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "children")
		})

		t.Run("create booking price_per_night required", func(t *testing.T) {
			body := cloneMap(validBody)
			body["rooms"] = []map[string]any{
				{
					"room_id":         env.Data.RoomID,
					"adults":          2,
					"children":        1,
					"price_per_night": "",
				},
			}
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "price_per_night")
		})

		t.Run("create booking price_per_night pattern", func(t *testing.T) {
			body := cloneMap(validBody)
			body["rooms"] = []map[string]any{
				{
					"room_id":         env.Data.RoomID,
					"adults":          2,
					"children":        1,
					"price_per_night": "150,00",
				},
			}
			assertValidationFields(t, env, http.MethodPost, createPath, env.Data.OwnerAccess, body, "price_per_night")
		})

		t.Run("get bookings userId positive may be normalized by gateway", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, createPath+"?userId=0&page=1&limit=10", env.Data.OwnerAccess, nil, nil)
			if status == http.StatusBadRequest {
				env.RequireError(status, http.StatusBadRequest, body, "userId must be a positive integer")
				return
			}
			if status != http.StatusOK {
				t.Fatalf("unexpected status for userId=0 normalization: got=%d body=%s", status, string(body))
			}
		})

		t.Run("get bookings page numeric in gateway parser", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, createPath+"?page=bad&limit=10", env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusBadRequest, body, "page must be a positive integer")
		})

		t.Run("get bookings limit numeric in gateway parser", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, createPath+"?page=1&limit=bad", env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusBadRequest, body, "limit must be a positive integer")
		})

		t.Run("get bookings hotel_id uuid", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, "/api/v1/jp/tokyo/hotels/not-a-uuid/bookings?page=1&limit=10", env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusBadRequest, body, "invalid hotel ID")
		})

		t.Run("get room bookings room_id uuid", func(t *testing.T) {
			status, body := env.RequestJSON(
				http.MethodGet,
				"/api/v1/jp/tokyo/hotels/"+env.Data.HotelID+"/rooms/not-a-uuid/bookings?page=1&limit=10",
				env.Data.OwnerAccess,
				nil,
				nil,
			)
			env.RequireError(status, http.StatusBadRequest, body, "invalid booking room ID")
		})

		t.Run("get bookings limit lte 100", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, createPath+"?page=1&limit=101", env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusBadRequest, body, "invalid pagination")
		})

		t.Run("get my bookings page numeric in gateway parser", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, myBookingsBasePath()+"?page=bad&limit=10", env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusBadRequest, body, "page must be a positive integer")
		})

		t.Run("get my bookings limit numeric in gateway parser", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, myBookingsBasePath()+"?page=1&limit=bad", env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusBadRequest, body, "limit must be a positive integer")
		})

		t.Run("get my bookings limit lte 100", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, myBookingsBasePath()+"?page=1&limit=101", env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusBadRequest, body, "invalid pagination")
		})

		t.Run("quote booking hotel_id uuid from path", func(t *testing.T) {
			assertValidationFields(t, env,
				http.MethodPost,
				"/api/v1/jp/tokyo/hotels/not-a-uuid/bookings/quote",
				"",
				validQuoteBody,
				"hotel_id",
			)
		})

		t.Run("quote booking check_in gt_now", func(t *testing.T) {
			body := cloneMap(validQuoteBody)
			body["check_in"] = time.Now().UTC().Add(-48 * time.Hour).Format(time.RFC3339)
			assertValidationFields(t, env, http.MethodPost, quoteBookingPath(env), "", body, "check_in")
		})

		t.Run("quote booking check_out after check_in cel", func(t *testing.T) {
			body := cloneMap(validQuoteBody)
			body["check_out"] = checkIn.Add(-time.Hour).Format(time.RFC3339)
			assertValidationFields(t, env, http.MethodPost, quoteBookingPath(env), "", body)
		})

		t.Run("quote booking currency pattern", func(t *testing.T) {
			body := cloneMap(validQuoteBody)
			body["currency"] = "usd"
			assertValidationFields(t, env, http.MethodPost, quoteBookingPath(env), "", body, "currency")
		})

		t.Run("quote booking rooms min_items", func(t *testing.T) {
			body := cloneMap(validQuoteBody)
			body["rooms"] = []map[string]any{}
			assertValidationFields(t, env, http.MethodPost, quoteBookingPath(env), "", body, "rooms")
		})

		t.Run("quote booking room_id uuid", func(t *testing.T) {
			body := cloneMap(validQuoteBody)
			body["rooms"] = []map[string]any{
				{
					"room_id":         "not-a-uuid",
					"adults":          2,
					"children":        1,
					"price_per_night": "150.00",
				},
			}
			assertValidationFields(t, env, http.MethodPost, quoteBookingPath(env), "", body, "room_id")
		})

		t.Run("quote booking price_per_night pattern", func(t *testing.T) {
			body := cloneMap(validQuoteBody)
			body["rooms"] = []map[string]any{
				{
					"room_id":         env.Data.RoomID,
					"adults":          2,
					"children":        1,
					"price_per_night": "150,00",
				},
			}
			assertValidationFields(t, env, http.MethodPost, quoteBookingPath(env), "", body, "price_per_night")
		})

		t.Run("availability check_in required", func(t *testing.T) {
			status, body := env.RequestJSON(
				http.MethodGet,
				availabilityBasePath(env)+"?check_out="+checkOut.Format("2006-01-02"),
				"",
				nil,
				nil,
			)
			env.RequireError(status, http.StatusBadRequest, body, "check_in is required")
		})

		t.Run("availability check_out required", func(t *testing.T) {
			status, body := env.RequestJSON(
				http.MethodGet,
				availabilityBasePath(env)+"?check_in="+checkIn.Format("2006-01-02"),
				"",
				nil,
				nil,
			)
			env.RequireError(status, http.StatusBadRequest, body, "check_out is required")
		})

		t.Run("availability check_in date format", func(t *testing.T) {
			status, body := env.RequestJSON(
				http.MethodGet,
				availabilityBasePath(env)+"?check_in=bad-date&check_out="+checkOut.Format("2006-01-02"),
				"",
				nil,
				nil,
			)
			env.RequireError(status, http.StatusBadRequest, body, "date must be in YYYY-MM-DD format")
		})

		t.Run("availability check_out after check_in", func(t *testing.T) {
			status, body := env.RequestJSON(
				http.MethodGet,
				availabilityBasePath(env)+"?check_in="+checkIn.Format("2006-01-02")+"&check_out="+checkIn.Format("2006-01-02"),
				"",
				nil,
				nil,
			)
			env.RequireError(status, http.StatusBadRequest, body, "check_out must be after check_in")
		})

		t.Run("availability page numeric in gateway parser", func(t *testing.T) {
			status, body := env.RequestJSON(
				http.MethodGet,
				availabilityBasePath(env)+"?check_in="+checkIn.Format("2006-01-02")+"&check_out="+checkOut.Format("2006-01-02")+"&page=bad&limit=10",
				"",
				nil,
				nil,
			)
			env.RequireError(status, http.StatusBadRequest, body, "page must be a positive integer")
		})

		t.Run("availability limit numeric in gateway parser", func(t *testing.T) {
			status, body := env.RequestJSON(
				http.MethodGet,
				availabilityBasePath(env)+"?check_in="+checkIn.Format("2006-01-02")+"&check_out="+checkOut.Format("2006-01-02")+"&page=1&limit=bad",
				"",
				nil,
				nil,
			)
			env.RequireError(status, http.StatusBadRequest, body, "limit must be a positive integer")
		})

		t.Run("availability hotel_id uuid", func(t *testing.T) {
			assertValidationFields(t,
				env,
				http.MethodGet,
				"/api/v1/jp/tokyo/hotels/not-a-uuid/rooms/availability?check_in="+checkIn.Format("2006-01-02")+"&check_out="+checkOut.Format("2006-01-02")+"&page=1&limit=10",
				"",
				nil,
				"hotel_id",
			)
		})

		t.Run("get booking id uuid", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodGet, createPath+"/not-a-uuid", env.Data.OwnerAccess, nil, "id")
		})

		t.Run("confirm booking id uuid", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPatch, createPath+"/not-a-uuid/confirm", env.Data.OwnerAccess, nil, "id")
		})

		t.Run("cancel booking id uuid", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodPatch, createPath+"/not-a-uuid/cancel", env.Data.OwnerAccess, nil, "id")
		})

		t.Run("delete booking id uuid", func(t *testing.T) {
			assertValidationFields(t, env, http.MethodDelete, createPath+"/not-a-uuid", env.Data.OwnerAccess, nil, "id")
		})

		t.Run("gateway normalization for get_bookings page and limit zero", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, createPath+"?page=0&limit=0", env.Data.OwnerAccess, nil, nil)
			if status == http.StatusBadRequest {
				t.Fatalf("not reachable via gateway normalization: get_bookings.page >= 1, get_bookings.limit >= 1; body=%s", string(body))
			}
		})
	})

	t.Run("booking grpc domain errors and statuses", func(t *testing.T) {
		createPath := bookingsBasePath(env)
		checkIn := time.Now().UTC().Add(30 * 24 * time.Hour).Truncate(time.Second)
		checkOut := checkIn.Add(48 * time.Hour)

		createReq := dto.CreateBookingRequest{
			UserId:              env.Data.OwnerID,
			CheckIn:             checkIn,
			CheckOut:            checkOut,
			GuestName:           "Booking Error Case",
			GuestEmail:          env.Data.OwnerEmail,
			GuestPhone:          "+79990000056",
			Currency:            "USD",
			ExpectedTotalAmount: "350.00",
			Rooms: []*dto.CreateBookingRoomRequest{
				{RoomId: env.Data.RoomID, Adults: 2, Children: 1, PricePerNight: "175.00"},
			},
		}

		t.Run("401 missing authorization header", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, createPath+"?page=1&limit=10", "", nil, nil)
			env.RequireError(status, http.StatusUnauthorized, body, "authorization header is required")
		})

		t.Run("401 missing authorization header on create booking", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPost, createPath, "", createReq, nil)
			env.RequireError(status, http.StatusUnauthorized, body, "authorization header is required")
		})

		t.Run("401 invalid jwt token", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, createPath+"?page=1&limit=10", "invalid.token.value", nil, nil)
			env.RequireError(status, http.StatusUnauthorized, body, "invalid token")
		})

		t.Run("401 missing authorization header on my bookings", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, myBookingsBasePath()+"?page=1&limit=10", "", nil, nil)
			env.RequireError(status, http.StatusUnauthorized, body, "authorization header is required")
		})

		t.Run("400 failed precondition expected total mismatch", func(t *testing.T) {
			req := createReq
			req.ExpectedTotalAmount = "1.00"

			status, body := env.RequestJSON(http.MethodPost, createPath, env.Data.OwnerAccess, req, nil)
			env.RequireError(status, http.StatusBadRequest, body, "expected total amount does not match calculated total")
		})

		var conflictBookingID string
		t.Run("create booking for conflict setup", func(t *testing.T) {
			var resp dto.BookingResponse
			status, body := env.RequestJSON(http.MethodPost, createPath, env.Data.OwnerAccess, createReq, &resp)
			env.RequireStatus(status, http.StatusCreated, body)
			conflictBookingID = resp.Id
			if resp.Status != "BOOKING_STATUS_PENDING" {
				t.Fatalf("unexpected booking status: got=%q want=%q", resp.Status, "BOOKING_STATUS_PENDING")
			}
		})

		t.Run("409 room lock already exists", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPost, createPath, env.Data.OwnerAccess, createReq, nil)
			env.RequireError(status, http.StatusConflict, body, "room lock already exists")
		})

		t.Run("cleanup conflict setup booking", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodDelete, bookingsByIDPath(env, conflictBookingID), env.Data.OwnerAccess, nil, nil)
			env.RequireStatus(status, http.StatusNoContent, body)
		})

		unknownBookingID := uuid.NewString()
		t.Run("404 get unknown booking", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodGet, bookingsByIDPath(env, unknownBookingID), env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusNotFound, body, "booking not found")
		})

		t.Run("404 confirm unknown booking", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPatch, bookingsByIDPath(env, unknownBookingID)+"/confirm", env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusNotFound, body, "booking not found")
		})

		t.Run("404 cancel unknown booking", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodPatch, bookingsByIDPath(env, unknownBookingID)+"/cancel", env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusNotFound, body, "booking not found")
		})

		t.Run("404 delete unknown booking", func(t *testing.T) {
			status, body := env.RequestJSON(http.MethodDelete, bookingsByIDPath(env, unknownBookingID), env.Data.OwnerAccess, nil, nil)
			env.RequireError(status, http.StatusNotFound, body, "booking not found")
		})
	})
}

func cloneMap(src map[string]any) map[string]any {
	dst := make(map[string]any, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func assertBookingResponseFields(t *testing.T, got *dto.BookingResponse, env *fixtures.Env) {
	t.Helper()

	if got == nil {
		t.Fatalf("booking response is nil")
	}
	if got.Id == "" {
		t.Fatalf("booking id is empty")
	}
	if got.UserId != env.Data.OwnerID {
		t.Fatalf("unexpected booking user_id: got=%d want=%d", got.UserId, env.Data.OwnerID)
	}
	if got.HotelId != env.Data.HotelID {
		t.Fatalf("unexpected booking hotel_id: got=%q want=%q", got.HotelId, env.Data.HotelID)
	}
	if got.GuestName == "" || got.Currency == "" || got.ExpectedTotalAmount == "" {
		t.Fatalf("booking has empty required fields: %+v", got)
	}
	if got.CheckIn.IsZero() || got.CheckOut.IsZero() || got.CreatedAt.IsZero() || got.UpdatedAt.IsZero() {
		t.Fatalf("booking has zero timestamps: %+v", got)
	}
	if len(got.BookingRooms) == 0 {
		t.Fatalf("booking_rooms is empty: %+v", got)
	}

	for _, room := range got.BookingRooms {
		if room.Id == "" || room.RoomId == "" || room.PricePerNight == "" {
			t.Fatalf("booking room has empty required fields: %+v", room)
		}
		if room.Adults == 0 {
			t.Fatalf("booking room adults must be > 0: %+v", room)
		}
	}
}

func assertBookingShortFields(t *testing.T, got *dto.BookingShortResponse) {
	t.Helper()

	if got == nil {
		t.Fatalf("booking short response is nil")
	}
	if got.Id == "" || got.HotelId == "" || got.GuestName == "" || got.Currency == "" {
		t.Fatalf("booking short has empty required fields: %+v", got)
	}
	if got.UserId <= 0 {
		t.Fatalf("booking short user_id must be > 0: %+v", got)
	}
	if got.CheckIn.IsZero() || got.CheckOut.IsZero() {
		t.Fatalf("booking short has zero dates: %+v", got)
	}
	for _, room := range got.BookingRooms {
		if room == nil {
			t.Fatalf("booking short room is nil")
		}
		if room.Id == "" || room.RoomId == "" || room.PricePerNight == "" {
			t.Fatalf("booking short room has empty required fields: %+v", room)
		}
	}
}

func assertBookingStatusEnum(t *testing.T, status string) {
	t.Helper()

	switch status {
	case "BOOKING_STATUS_PENDING", "BOOKING_STATUS_CONFIRMED", "BOOKING_STATUS_CANCELLED":
		return
	default:
		t.Fatalf("unexpected booking status enum: %q", status)
	}
}

func findBookingByID(bookings []*dto.BookingShortResponse, id string) *dto.BookingShortResponse {
	for _, booking := range bookings {
		if booking != nil && booking.Id == id {
			return booking
		}
	}
	return nil
}

func findBookingByIDInResponse(resp *dto.BookingsResponse, id string) *dto.BookingShortResponse {
	if resp == nil {
		return nil
	}
	return findBookingByID(resp.Bookings, id)
}

func mustQueryUserID(id int64) string {
	return strconv.FormatInt(id, 10)
}
