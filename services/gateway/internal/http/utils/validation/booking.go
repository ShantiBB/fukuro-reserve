package validation

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	valconst "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/validation/constants"
)

func ValidateCreateBookingRequest(req dto.CreateBookingRequest) error {
	if strings.TrimSpace(req.GuestName) == "" {
		return errors.New(valconst.ErrBookingGuestNameRequired)
	}
	if err := ValidateUUID(req.HotelId); err != nil {
		return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldBookingHotelID, err)
	}
	if err := ValidateCurrency(req.Currency); err != nil {
		return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldBookingCurrency, err)
	}
	if err := ValidateAmount(req.ExpectedTotalAmount); err != nil {
		return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldExpectedTotalAmount, err)
	}
	if !req.CheckIn.After(time.Now()) {
		return errors.New(valconst.ErrBookingCheckInFuture)
	}
	if !req.CheckOut.After(req.CheckIn) {
		return errors.New(valconst.ErrBookingCheckOutAfterIn)
	}
	if len(req.Rooms) == 0 {
		return errors.New(valconst.ErrBookingAtLeastOneRoom)
	}

	for i, room := range req.Rooms {
		if room == nil {
			return fmt.Errorf(valconst.ErrRoomsIndexedRequired, i)
		}
		if err := ValidateRoom(room); err != nil {
			return fmt.Errorf(valconst.ErrRoomsIndexedWrap, i, err)
		}
	}

	if req.GuestEmail != "" {
		if err := validateEmail(req.GuestEmail); err != nil {
			return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldBookingGuestEmail, err)
		}
	}

	return nil
}
