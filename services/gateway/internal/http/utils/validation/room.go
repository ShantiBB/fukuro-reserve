package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
)

func ValidateCreateRoomRequest(req dto.CreateRoomRequest) error {
	if strings.TrimSpace(req.RoomNumber) == "" {
		return errors.New(consts.ErrRoomNumberRequired)
	}
	if strings.TrimSpace(req.Type) == "" {
		return errors.New(consts.ErrRoomTypeRequired)
	}
	if strings.TrimSpace(req.HotelSlug) == "" {
		return errors.New(consts.ErrRoomHotelSlugReq)
	}
	if strings.TrimSpace(req.Title) == "" {
		return errors.New(consts.ErrTitleRequired)
	}
	if strings.TrimSpace(req.CountryCode) == "" {
		return errors.New(consts.ErrCountryCodeRequired)
	}
	if strings.TrimSpace(req.CitySlug) == "" {
		return errors.New(consts.ErrCitySlugRequired)
	}
	if err := ValidateAmount(req.Price); err != nil {
		return fmt.Errorf(consts.ErrWrapWithField, consts.FieldRoomPrice, err)
	}
	if req.Capacity < 1 || req.Capacity > 20 {
		return errors.New(consts.ErrRoomCapacityRange)
	}
	if req.Floor < 0 {
		return errors.New(consts.ErrRoomFloorNonNeg)
	}
	if req.AreaSqm <= 0 {
		return errors.New(consts.ErrRoomAreaPositive)
	}
	return nil
}

func ValidateUpdateRoomRequest(req dto.UpdateRoomRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New(consts.ErrTitleRequired)
	}
	if strings.TrimSpace(req.RoomNumber) == "" {
		return errors.New(consts.ErrRoomNumberRequired)
	}
	if strings.TrimSpace(req.Type) == "" {
		return errors.New(consts.ErrRoomTypeRequired)
	}
	if err := ValidateAmount(req.Price); err != nil {
		return fmt.Errorf(consts.ErrWrapWithField, consts.FieldRoomPrice, err)
	}
	if req.Capacity < 1 || req.Capacity > 20 {
		return errors.New(consts.ErrRoomCapacityRange)
	}
	if req.Floor < 0 {
		return errors.New(consts.ErrRoomFloorNonNeg)
	}
	if req.AreaSqm <= 0 {
		return errors.New(consts.ErrRoomAreaPositive)
	}
	return nil
}

func ValidateUpdateRoomStatusRequest(req dto.UpdateRoomStatusRequest) error {
	if strings.TrimSpace(req.Status) == "" {
		return errors.New(consts.ErrRoomStatusRequired)
	}
	return nil
}

func ValidateRoom(room *dto.CreateBookingRoomRequest) error {
	if err := ValidateUUID(room.RoomId); err != nil {
		return fmt.Errorf(consts.ErrWrapWithField, consts.FieldRoomID, err)
	}
	if room.Adults < 1 || room.Adults > 20 {
		return errors.New(consts.ErrAdultsRange)
	}
	if room.Children > 20 {
		return errors.New(consts.ErrChildrenRange)
	}
	if err := ValidateAmount(room.PricePerNight); err != nil {
		return fmt.Errorf(consts.ErrWrapWithField, consts.FieldRoomPricePerNight, err)
	}

	return nil
}
