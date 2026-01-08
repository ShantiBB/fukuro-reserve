package request

import (
	"hotel/internal/repository/models"

	"github.com/shopspring/decimal"
)

type RoomCreate struct {
	Title       string          `json:"title" validate:"required"`
	Description string          `json:"description"`
	RoomNumber  string          `json:"room_number" validate:"required"`
	Type        models.RoomType `json:"type" validate:"required"`
	Price       decimal.Decimal `json:"price" validate:"required"`
	Capacity    int             `json:"capacity" validate:"required"`
	AreaSqm     float64         `json:"area_sqm"`
	Floor       int             `json:"floor"`
	Amenities   []string        `json:"amenities"`
	Images      []string        `json:"images"`
}

type RoomUpdate struct {
	Title       string          `json:"title" validate:"required"`
	Description string          `json:"description" validate:"required"`
	RoomNumber  string          `json:"room_number" validate:"required"`
	Type        models.RoomType `json:"type" validate:"required"`
	Price       decimal.Decimal `json:"price" validate:"required"`
	Capacity    int             `json:"capacity" validate:"required"`
	AreaSqm     float64         `json:"area_sqm" validate:"required"`
	Floor       int             `json:"floor" validate:"required"`
	Amenities   []string        `json:"amenities" validate:"required"`
	Images      []string        `json:"images" validate:"required"`
}
