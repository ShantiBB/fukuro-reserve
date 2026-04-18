package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	valconst "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/validation/constants"
)

func ValidateCreateRoomRequest(req dto.CreateRoomRequest) error {
	if strings.TrimSpace(req.RoomNumber) == "" {
		return errors.New(valconst.ErrRoomNumberRequired)
	}
	if strings.TrimSpace(req.Type) == "" {
		return errors.New(valconst.ErrRoomTypeRequired)
	}
	if strings.TrimSpace(req.HotelSlug) == "" {
		return errors.New(valconst.ErrRoomHotelSlugReq)
	}
	if strings.TrimSpace(req.Title) == "" {
		return errors.New(valconst.ErrTitleRequired)
	}
	if strings.TrimSpace(req.CountryCode) == "" {
		return errors.New(valconst.ErrCountryCodeRequired)
	}
	if strings.TrimSpace(req.CitySlug) == "" {
		return errors.New(valconst.ErrCitySlugRequired)
	}
	if err := ValidateAmount(req.Price); err != nil {
		return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldRoomPrice, err)
	}
	if req.Capacity < 1 || req.Capacity > 20 {
		return errors.New(valconst.ErrRoomCapacityRange)
	}
	if req.Floor < 0 {
		return errors.New(valconst.ErrRoomFloorNonNeg)
	}
	if req.AreaSqm <= 0 {
		return errors.New(valconst.ErrRoomAreaPositive)
	}
	return nil
}

func ValidateUpdateRoomRequest(req dto.UpdateRoomRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New(valconst.ErrTitleRequired)
	}
	if strings.TrimSpace(req.RoomNumber) == "" {
		return errors.New(valconst.ErrRoomNumberRequired)
	}
	if strings.TrimSpace(req.Type) == "" {
		return errors.New(valconst.ErrRoomTypeRequired)
	}
	if err := ValidateAmount(req.Price); err != nil {
		return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldRoomPrice, err)
	}
	if req.Capacity < 1 || req.Capacity > 20 {
		return errors.New(valconst.ErrRoomCapacityRange)
	}
	if req.Floor < 0 {
		return errors.New(valconst.ErrRoomFloorNonNeg)
	}
	if req.AreaSqm <= 0 {
		return errors.New(valconst.ErrRoomAreaPositive)
	}
	return nil
}

func ValidateUpdateRoomStatusRequest(req dto.UpdateRoomStatusRequest) error {
	if strings.TrimSpace(req.Status) == "" {
		return errors.New(valconst.ErrRoomStatusRequired)
	}
	return nil
}

func ValidateRoom(room *dto.CreateBookingRoomRequest) error {
	if err := ValidateUUID(room.RoomId); err != nil {
		return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldRoomID, err)
	}
	if room.Adults < 1 || room.Adults > 20 {
		return errors.New(valconst.ErrAdultsRange)
	}
	if room.Children > 20 {
		return errors.New(valconst.ErrChildrenRange)
	}
	if err := ValidateAmount(room.PricePerNight); err != nil {
		return fmt.Errorf(valconst.ErrWrapWithField, valconst.FieldRoomPricePerNight, err)
	}

	return nil
}
