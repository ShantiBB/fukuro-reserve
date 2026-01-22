package helper

import (
	"time"

	"booking/pkg/lib/utils/consts"
)

func Nights(checkIn, checkOut time.Time) (int, error) {
	in := time.Date(
		checkIn.Year(), checkIn.Month(), checkIn.Day(),
		0, 0, 0, 0, time.UTC,
	)
	out := time.Date(
		checkOut.Year(), checkOut.Month(), checkOut.Day(),
		0, 0, 0, 0, time.UTC,
	)

	nights := int(out.Sub(in).Hours() / 24)
	if nights <= 0 {
		return 0, consts.ErrInvalidDates
	}
	return nights, nil
}
