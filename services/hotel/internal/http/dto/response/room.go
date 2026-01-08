package response

import (
	"fukuro-reserve/pkg/utils/helper"
	"hotel/internal/repository/models"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Room struct {
	ID          uuid.UUID         `json:"id"`
	Title       string            `json:"title"`
	Description *string           `json:"description"`
	RoomNumber  string            `json:"room_number"`
	Type        models.RoomType   `json:"type"`
	Status      models.RoomStatus `json:"status"`
	Price       decimal.Decimal   `json:"price"`
	Capacity    int               `json:"capacity"`
	AreaSqm     float64           `json:"area_sqm"`
	Floor       int               `json:"floor"`
	Amenities   []string          `json:"amenities"`
	Images      []string          `json:"images"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type RoomShort struct {
	ID         uuid.UUID         `json:"id"`
	Title      string            `json:"title"`
	RoomNumber string            `json:"room_number"`
	Type       models.RoomType   `json:"type"`
	Status     models.RoomStatus `json:"status"`
	Price      decimal.Decimal   `json:"price"`
	Capacity   int               `json:"capacity"`
	AreaSqm    float64           `json:"area_sqm"`
	Amenities  []string          `json:"amenities"`
	Images     []string          `json:"images"`
}

type RoomUpdate struct {
	Title       string          `json:"title"`
	Description *string         `json:"description"`
	RoomNumber  string          `json:"room_number"`
	Type        models.RoomType `json:"type"`
	Price       decimal.Decimal `json:"price"`
	Capacity    int             `json:"capacity"`
	AreaSqm     float64         `json:"area_sqm"`
	Floor       int             `json:"floor"`
	Amenities   []string        `json:"amenities"`
	Images      []string        `json:"images"`
}
type RoomList struct {
	Rooms           []RoomShort            `json:"rooms"`
	CurrentPage     uint64                 `json:"current_page"`
	Limit           uint64                 `json:"limit"`
	Links           helper.PaginationLinks `json:"links"`
	TotalPageCount  uint64                 `json:"total_page_count"`
	TotalRoomsCount uint64                 `json:"total_rooms_count"`
}
