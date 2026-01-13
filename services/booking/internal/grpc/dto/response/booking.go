package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"booking/internal/repository/models"
)

type Booking struct {
	ID               uuid.UUID            `json:"id"`
	UserID           int64                `json:"user_id"`
	HotelID          string               `json:"hotel_id"`
	BookingRooms     []BookingRoom        `json:"booking_rooms"`
	CheckIn          time.Time            `json:"check_in"`
	CheckOut         time.Time            `json:"check_out"`
	Status           models.BookingStatus `json:"status"`
	GuestName        string               `json:"guest_name"`
	GuestEmail       *string              `json:"guest_email"`
	GuestPhone       *string              `json:"guest_phone"`
	Currency         string               `json:"currency"`
	FinalTotalAmount decimal.Decimal      `json:"final_total_amount"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
}
