package validation

import (
	"errors"
	"strings"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	valconst "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/utils/validation/constants"
)

func ValidateCreateHotelRequest(req dto.CreateHotelRequest) error {
	if strings.TrimSpace(req.CountryCode) == "" {
		return errors.New(valconst.ErrCountryCodeRequired)
	}
	if strings.TrimSpace(req.CitySlug) == "" {
		return errors.New(valconst.ErrCitySlugRequired)
	}
	if strings.TrimSpace(req.Title) == "" {
		return errors.New(valconst.ErrTitleRequired)
	}
	if strings.TrimSpace(req.Address) == "" {
		return errors.New(valconst.ErrAddressRequired)
	}
	if req.OwnerId <= 0 {
		return errors.New(valconst.ErrOwnerIDPositive)
	}
	if req.Location != nil {
		if req.Location.Latitude < -90 || req.Location.Latitude > 90 {
			return errors.New(valconst.ErrLatitudeOutOfRange)
		}
		if req.Location.Longitude < -180 || req.Location.Longitude > 180 {
			return errors.New(valconst.ErrLongitudeOutOfRange)
		}
	}
	return nil
}

func ValidateUpdateHotelRequest(req dto.UpdateHotelRequest) error {
	if strings.TrimSpace(req.Description) == "" && strings.TrimSpace(req.Address) == "" && req.Location == nil {
		return errors.New(valconst.ErrHotelUpdateRequired)
	}
	if req.Location != nil {
		if req.Location.Latitude < -90 || req.Location.Latitude > 90 {
			return errors.New(valconst.ErrLatitudeOutOfRange)
		}
		if req.Location.Longitude < -180 || req.Location.Longitude > 180 {
			return errors.New(valconst.ErrLongitudeOutOfRange)
		}
	}
	return nil
}

func ValidateUpdateHotelTitleRequest(req dto.UpdateHotelTitleRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New(valconst.ErrTitleRequired)
	}
	return nil
}
