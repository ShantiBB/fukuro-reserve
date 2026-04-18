package validation

import (
	"errors"
	"strings"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
)

func ValidateCreateHotelRequest(req dto.CreateHotelRequest) error {
	if strings.TrimSpace(req.CountryCode) == "" {
		return errors.New(consts.ErrCountryCodeRequired)
	}
	if strings.TrimSpace(req.CitySlug) == "" {
		return errors.New(consts.ErrCitySlugRequired)
	}
	if strings.TrimSpace(req.Title) == "" {
		return errors.New(consts.ErrTitleRequired)
	}
	if strings.TrimSpace(req.Address) == "" {
		return errors.New(consts.ErrAddressRequired)
	}
	if req.OwnerId <= 0 {
		return errors.New(consts.ErrOwnerIDPositive)
	}
	if req.Location != nil {
		if req.Location.Latitude < -90 || req.Location.Latitude > 90 {
			return errors.New(consts.ErrLatitudeOutOfRange)
		}
		if req.Location.Longitude < -180 || req.Location.Longitude > 180 {
			return errors.New(consts.ErrLongitudeOutOfRange)
		}
	}
	return nil
}

func ValidateUpdateHotelRequest(req dto.UpdateHotelRequest) error {
	if strings.TrimSpace(req.Description) == "" && strings.TrimSpace(req.Address) == "" && req.Location == nil {
		return errors.New(consts.ErrHotelUpdateRequired)
	}
	if req.Location != nil {
		if req.Location.Latitude < -90 || req.Location.Latitude > 90 {
			return errors.New(consts.ErrLatitudeOutOfRange)
		}
		if req.Location.Longitude < -180 || req.Location.Longitude > 180 {
			return errors.New(consts.ErrLongitudeOutOfRange)
		}
	}
	return nil
}

func ValidateUpdateHotelTitleRequest(req dto.UpdateHotelTitleRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New(consts.ErrTitleRequired)
	}
	return nil
}
