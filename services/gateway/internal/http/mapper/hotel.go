package mapper

import (
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/dto"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

func HotelResponseFromProto(hotel *hotelv1.CreateHotel) *dto.HotelResponse {
	if hotel == nil {
		return nil
	}

	resp := &dto.HotelResponse{
		Id:          hotel.Id,
		Title:       hotel.Title,
		HotelSlug:   hotel.HotelSlug,
		OwnerId:     hotel.OwnerId,
		Description: hotel.Description,
		Address:     hotel.Address,
		Location:    locationDTOFromProto(hotel.Location),
	}
	if hotel.CreatedAt != nil {
		resp.CreatedAt = hotel.CreatedAt.AsTime()
	}
	if hotel.UpdatedAt != nil {
		resp.UpdatedAt = hotel.UpdatedAt.AsTime()
	}

	return resp
}

func HotelDetailResponseFromProto(resp *hotelv1.GetHotelResponse) *dto.HotelResponse {
	if resp == nil || resp.Hotel == nil {
		return nil
	}

	hotel := resp.Hotel
	result := &dto.HotelResponse{
		Id:          hotel.Id,
		Title:       hotel.Title,
		OwnerId:     hotel.OwnerId,
		Description: hotel.Description,
		Address:     hotel.Address,
		Location:    locationDTOFromProto(hotel.Location),
	}
	if hotel.Rating != nil {
		result.Rating = *hotel.Rating
	}
	if hotel.CreatedAt != nil {
		result.CreatedAt = hotel.CreatedAt.AsTime()
	}
	if hotel.UpdatedAt != nil {
		result.UpdatedAt = hotel.UpdatedAt.AsTime()
	}

	return result
}

func HotelsShortResponseFromProto(resp *hotelv1.GetHotelsResponse) *dto.HotelsResponse {
	if resp == nil {
		return nil
	}

	hotels := make([]*dto.HotelShortResponse, len(resp.Hotels))
	for i, hotel := range resp.Hotels {
		hotels[i] = hotelShortResponseFromProto(hotel)
	}

	return &dto.HotelsResponse{Hotels: hotels}
}

func UpdateHotelResponseFromProto(hotel *hotelv1.UpdateHotel) *dto.HotelResponse {
	if hotel == nil {
		return nil
	}

	result := &dto.HotelResponse{
		Description: hotel.Description,
		Address:     hotel.Address,
		Location:    locationDTOFromProto(hotel.Location),
	}
	if hotel.Location != nil {
		result.Location = locationDTOFromProto(hotel.Location)
	}

	return result
}

func UpdateHotelTitleResponseFromProto(hotel *hotelv1.UpdateHotelTitle) *dto.HotelResponse {
	if hotel == nil {
		return nil
	}

	return &dto.HotelResponse{
		Title:     hotel.Title,
		HotelSlug: hotel.HotelSlug,
	}
}

func locationDTOFromProto(loc *hotelv1.Location) *dto.LocationDTO {
	if loc == nil {
		return nil
	}

	return &dto.LocationDTO{
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
	}
}

func hotelShortResponseFromProto(hotel *hotelv1.HotelShort) *dto.HotelShortResponse {
	if hotel == nil {
		return nil
	}

	resp := &dto.HotelShortResponse{
		Id:        hotel.Id,
		Title:     hotel.Title,
		HotelSlug: hotel.HotelSlug,
		OwnerId:   hotel.OwnerId,
		Address:   hotel.Address,
		Location:  locationDTOFromProto(hotel.Location),
	}
	if hotel.Rating != nil {
		resp.Rating = *hotel.Rating
	}

	return resp
}
