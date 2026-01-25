package response

import (
	"time"

	"hotel/internal/http/utils/pagination"
	"hotel/internal/repository/models"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Room struct {
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Description *string           `json:"description"`
	Price       decimal.Decimal   `json:"price"`
	Type        models.RoomType   `json:"type"`
	Status      models.RoomStatus `json:"status"`
	RoomNumber  string            `json:"room_number"`
	Title       string            `json:"title"`
	Amenities   []string          `json:"amenities"`
	Images      []string          `json:"images"`
	Capacity    int               `json:"capacity"`
	AreaSqm     float64           `json:"area_sqm"`
	Floor       int               `json:"floor"`
	ID          uuid.UUID         `json:"id"`
}

type RoomShort struct {
	Title      string            `json:"title"`
	RoomNumber string            `json:"room_number"`
	Type       models.RoomType   `json:"type"`
	Status     models.RoomStatus `json:"status"`
	Price      decimal.Decimal   `json:"price"`
	Amenities  []string          `json:"amenities"`
	Images     []string          `json:"images"`
	Capacity   int               `json:"capacity"`
	AreaSqm    float64           `json:"area_sqm"`
	ID         uuid.UUID         `json:"id"`
}

type RoomUpdate struct {
	Description *string         `json:"description"`
	Title       string          `json:"title"`
	RoomNumber  string          `json:"room_number"`
	Type        models.RoomType `json:"type"`
	Price       decimal.Decimal `json:"price"`
	Amenities   []string        `json:"amenities"`
	Images      []string        `json:"images"`
	Capacity    int             `json:"capacity"`
	AreaSqm     float64         `json:"area_sqm"`
	Floor       int             `json:"floor"`
}

type RoomStatusUpdate struct {
	Status models.RoomStatus `json:"status"`
}

type RoomList struct {
	Links           pagination.Links `json:"links"`
	Rooms           []RoomShort      `json:"rooms"`
	CurrentPage     uint64           `json:"current_page"`
	Limit           uint64           `json:"limit"`
	TotalPageCount  uint64           `json:"total_page_count"`
	TotalRoomsCount uint64           `json:"total_rooms_count"`
}
