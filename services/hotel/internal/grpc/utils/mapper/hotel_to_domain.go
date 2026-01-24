package mapper

import (
	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/repository/models"
)

func locationRequestToDomain(req *hotelv1.CreateBookingLocationRequest) models.Location {
	return models.Location{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
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
