package helper

import (
	"time"

	"github.com/shopspring/decimal"

	"booking/internal/repository/models"
	"booking/pkg/utils/consts"
)

func CalculateFinalTotalAmount(
	checkIn time.Time,
	checkOut time.Time,
	rooms []*models.CreateBookingRoom,
	expected decimal.Decimal,
) (decimal.Decimal, error) {

	nights, err := Nights(checkIn, checkOut)
	if err != nil {
		return decimal.Zero, err
	}

	total := decimal.Zero
	nightsDec := decimal.NewFromInt(int64(nights))

	for _, room := range rooms {
		roomTotal := room.PricePerNight.Mul(nightsDec)
		total = total.Add(roomTotal)
	}

	if !expected.IsZero() {
		diff := total.Sub(expected).Abs()
		if diff.GreaterThan(decimal.RequireFromString("0.01")) {
			return decimal.Zero, consts.ErrPriceChanged
		}
	}

	return total, nil
}
