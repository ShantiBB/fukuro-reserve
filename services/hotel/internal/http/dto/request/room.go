package request

import (
	"hotel/internal/repository/models"

	"github.com/shopspring/decimal"
)

type RoomCreate struct {
	Title       *string          `json:"title" validate:"required,min=3,max=100"`
	Description *string          `json:"description" validate:"omitempty,max=1000"`
	RoomNumber  *string          `json:"room_number" validate:"required,min=1,max=10"`
	Type        *models.RoomType `json:"type" validate:"required,room_type"`
	Price       *decimal.Decimal `json:"price" validate:"required,decimal_gt=0,decimal_lt=100000000"`
	Capacity    *int             `json:"capacity" validate:"required,gte=1,lte=10"`
	AreaSqm     *float64         `json:"area_sqm" validate:"required,gt=0,lte=9999.99"`
	Floor       *int             `json:"floor" validate:"required,gte=0,lte=2147483647"`
	Amenities   []string         `json:"amenities" validate:"omitempty,max=50,dive,min=1,max=100"`
	Images      []string         `json:"images" validate:"omitempty,max=50,dive,min=1,max=500"`
}

type RoomUpdate struct {
	Title       *string          `json:"title" validate:"required,min=3,max=100"`
	Description *string          `json:"description" validate:"omitempty,max=1000"`
	RoomNumber  *string          `json:"room_number" validate:"required,min=1,max=10"`
	Type        *models.RoomType `json:"type" validate:"required,room_type"`
	Price       *decimal.Decimal `json:"price" validate:"required,decimal_gt=0,decimal_lt=100000000"`
	Capacity    *int             `json:"capacity" validate:"required,gte=1,lte=10"`
	AreaSqm     *float64         `json:"area_sqm" validate:"required,gt=0,lte=9999.99"`
	Floor       *int             `json:"floor" validate:"required,gte=0,lte=2147483647"`
	Amenities   []string         `json:"amenities" validate:"omitempty,max=50,dive,min=1,max=100"`
	Images      []string         `json:"images" validate:"omitempty,max=50,dive,min=1,max=500"`
}

type RoomStatusUpdate struct {
	Status models.RoomStatus `json:"status" validate:"required,room_status"`
}
