package validation

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
)

func ValidateCreateBookingRequest(req dto.CreateBookingRequest) error {
	if strings.TrimSpace(req.GuestName) == "" {
		return errors.New(consts.ErrBookingGuestNameRequired)
	}
	if err := ValidateUUID(req.HotelId); err != nil {
		return fmt.Errorf(consts.ErrWrapWithField, consts.FieldBookingHotelID, err)
	}
	if err := ValidateCurrency(req.Currency); err != nil {
		return fmt.Errorf(consts.ErrWrapWithField, consts.FieldBookingCurrency, err)
	}
	if err := ValidateAmount(req.ExpectedTotalAmount); err != nil {
		return fmt.Errorf(consts.ErrWrapWithField, consts.FieldExpectedTotalAmount, err)
	}
	if !req.CheckIn.After(time.Now()) {
		return errors.New(consts.ErrBookingCheckInFuture)
	}
	if !req.CheckOut.After(req.CheckIn) {
		return errors.New(consts.ErrBookingCheckOutAfterIn)
	}
	if len(req.Rooms) == 0 {
		return errors.New(consts.ErrBookingAtLeastOneRoom)
	}

	for i, room := range req.Rooms {
		if room == nil {
			return fmt.Errorf(consts.ErrRoomsIndexedRequired, i)
		}
		if err := ValidateRoom(room); err != nil {
			return fmt.Errorf(consts.ErrRoomsIndexedWrap, i, err)
		}
	}

	if req.GuestEmail != "" {
		if err := validateEmail(req.GuestEmail); err != nil {
			return fmt.Errorf(consts.ErrWrapWithField, consts.FieldBookingGuestEmail, err)
		}
	}

	return nil
}
