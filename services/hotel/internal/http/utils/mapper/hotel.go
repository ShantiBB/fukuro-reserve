package mapper

import (
	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
	"hotel/internal/repository/models"
)

func HotelCreateRequestToEntity(req request.HotelCreate) models.HotelCreate {
	location := models.Location{
		Latitude:  *req.Location.Latitude,
		Longitude: *req.Location.Longitude,
	}
	return models.HotelCreate{
		Title:       *req.Title,
		OwnerID:     *req.OwnerID,
		Description: req.Description,
		Address:     *req.Address,
		Location:    location,
	}
}

func HotelUpdateRequestToEntity(req request.HotelUpdate) models.HotelUpdate {
	location := models.Location{
		Latitude:  *req.Location.Latitude,
		Longitude: *req.Location.Longitude,
	}
	return models.HotelUpdate{
		Description: req.Description,
		Address:     *req.Address,
		Location:    location,
	}
}

func HotelTitleUpdateRequestToEntity(req request.HotelTitleUpdate) models.HotelTitleUpdate {
	return models.HotelTitleUpdate{
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
		Slug:        req.Slug,
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
		Slug:     req.Slug,
		OwnerID:  req.OwnerID,
		Address:  req.Address,
		Rating:   req.Rating,
		Location: location,
	}
}

func HotelUpdateEntityToResponse(req models.HotelUpdate) response.HotelUpdate {
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

func HotelTitleUpdateEntityToResponse(req models.HotelTitleUpdate) response.HotelTitleUpdate {
	return response.HotelTitleUpdate{
		Title: req.Title,
		Slug:  req.Slug,
	}
}

func HotelPathParamsToEntity(req request.HotelPathParams) models.HotelRef {
	return models.HotelRef{
		CountryCode: req.CountryCode,
		CitySlug:    req.CitySlug,
		HotelSlug:   req.HotelSlug,
	}
}
