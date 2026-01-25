package mapper

import (
	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/repository/models"
)

type locationGetter interface {
	GetLatitude() float32
	GetLongitude() float32
}

type hotelRefGetter interface {
	GetCountryCode() string
	GetCitySlug() string
	GetHotelSlug() string
}

func locationRequestToDomain[T locationGetter](req T) models.Location {
	return models.Location{
		Latitude:  req.GetLatitude(),
		Longitude: req.GetLongitude(),
	}
}

func GetHotelRefRequestToDomain[T hotelRefGetter](req T) models.HotelRef {
	return models.HotelRef{
		CountryCode: req.GetCountryCode(),
		CitySlug:    req.GetCitySlug(),
		HotelSlug:   req.GetHotelSlug(),
	}
}

func CreateHotelRequestToDomain(req *hotelv1.CreateHotelRequest) *models.CreateHotel {
	return &models.CreateHotel{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		Title:       req.Title,
		OwnerID:     req.OwnerId,
		Description: req.Description,
		Address:     req.Address,
		Location:    locationRequestToDomain(req.Location),
	}
}

func GetHotelsRequestToDomain(req *hotelv1.GetHotelsRequest) (uint64, uint64, models.HotelRef) {
	hotelInfo := models.HotelRef{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
	}
	return req.Page, req.Limit, hotelInfo
}

func UpdateHotelRequestToDomain(req *hotelv1.UpdateHotelRequest) models.UpdateHotel {
	return models.UpdateHotel{
		Description: req.Description,
		Address:     req.Address,
		Location:    locationRequestToDomain(req.Location),
	}
}

func UpdateHotelTitleRequestToDomain(req *hotelv1.UpdateHotelTitleRequest) models.UpdateHotelTitle {
	return models.UpdateHotelTitle{
		Title: req.Title,
	}
}
