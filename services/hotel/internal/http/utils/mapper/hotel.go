package mapper

import (
	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
	"hotel/internal/repository/models"
)

func HotelCreateRequestToEntity(req request.HotelCreate) models.CreateHotel {
	location := models.Location{
		Latitude:  *req.Location.Latitude,
		Longitude: *req.Location.Longitude,
	}
	return models.CreateHotel{
		Title:       *req.Title,
		OwnerID:     *req.OwnerID,
		Description: req.Description,
		Address:     *req.Address,
		Location:    location,
	}
}

func HotelUpdateRequestToEntity(req request.HotelUpdate) models.UpdateHotel {
	location := models.Location{
		Latitude:  *req.Location.Latitude,
		Longitude: *req.Location.Longitude,
	}
	return models.UpdateHotel{
		Description: req.Description,
		Address:     *req.Address,
		Location:    location,
	}
}

func HotelTitleUpdateRequestToEntity(req request.HotelTitleUpdate) models.UpdateHotelTitle {
	return models.UpdateHotelTitle{
		Title: *req.Title,
	}
}

func HotelCreateEntityToResponse(req models.Hotel) response.HotelCreate {
	location := response.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return response.HotelCreate{
		ID:          req.ID,
		Title:       req.Title,
		Slug:        req.HotelSlug,
		OwnerID:     req.OwnerID,
		Description: req.Description,
		Address:     req.Address,
		Location:    location,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}
}

func HotelGetEntityToResponse(req models.Hotel) response.Hotel {
	location := response.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return response.Hotel{
		ID:          req.ID,
		Title:       req.Title,
		OwnerID:     req.OwnerID,
		Description: req.Description,
		Address:     req.Address,
		Location:    location,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}
}

func HotelShortEntityToShortResponse(req models.HotelShort) response.HotelShort {
	location := response.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return response.HotelShort{
		ID:       req.ID,
		Title:    req.Title,
		Slug:     req.HotelSlug,
		OwnerID:  req.OwnerID,
		Address:  req.Address,
		Rating:   req.Rating,
		Location: location,
	}
}

func HotelUpdateEntityToResponse(req models.UpdateHotel) response.HotelUpdate {
	location := response.Location{
		Latitude:  req.Location.Latitude,
		Longitude: req.Location.Longitude,
	}
	return response.HotelUpdate{
		Description: req.Description,
		Address:     req.Address,
		Location:    location,
	}
}

func HotelTitleUpdateEntityToResponse(req models.UpdateHotelTitle) response.HotelTitleUpdate {
	return response.HotelTitleUpdate{
		Title: req.Title,
		Slug:  req.HotelSlug,
	}
}

func HotelPathParamsToEntity(req request.HotelPathParams) models.HotelRef {
	return models.HotelRef{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		HotelSlug:   req.HotelSlug,
	}
}
