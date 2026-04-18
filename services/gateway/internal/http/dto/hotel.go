package dto

import (
	"time"

	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

// Hotel DTOs.
type LocationDTO struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type CreateHotelRequest struct {
	Location    *LocationDTO `json:"location,omitempty"`
	CountryCode string       `json:"country_code"`
	CitySlug    string       `json:"city_slug"`
	Title       string       `json:"title"`
	Description string       `json:"description,omitempty"`
	Address     string       `json:"address"`
	OwnerId     int64        `json:"owner_id"`
}

type UpdateHotelRequest struct {
	Location    *LocationDTO `json:"location,omitempty"`
	Description string       `json:"description"`
	Address     string       `json:"address"`
}

type UpdateHotelTitleRequest struct {
	Title string `json:"title"`
}

type HotelResponse struct {
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Location    *LocationDTO `json:"location,omitempty"`
	Id          string       `json:"id"`
	Title       string       `json:"title"`
	HotelSlug   string       `json:"hotel_slug,omitempty"`
	Description string       `json:"description"`
	Address     string       `json:"address"`
	OwnerId     int64        `json:"owner_id"`
	Rating      float32      `json:"rating,omitempty"`
}

type HotelShortResponse struct {
	Location  *LocationDTO `json:"location,omitempty"`
	Id        string       `json:"id"`
	Title     string       `json:"title"`
	HotelSlug string       `json:"hotel_slug"`
	Address   string       `json:"address"`
	OwnerId   int64        `json:"owner_id"`
	Rating    float32      `json:"rating,omitempty"`
}

type HotelsResponse struct {
	Hotels []*HotelShortResponse `json:"hotels"`
}

func locationDTOFromProto(loc *hotelv1.Location) *LocationDTO {
	if loc == nil {
		return nil
	}

	return &LocationDTO{
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
	}
}

func HotelResponseFromProto(hotel *hotelv1.CreateHotel) *HotelResponse {
	if hotel == nil {
		return nil
	}

	resp := &HotelResponse{
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

func HotelDetailResponseFromProto(resp *hotelv1.GetHotelResponse) *HotelResponse {
	if resp == nil || resp.Hotel == nil {
		return nil
	}

	hotel := resp.Hotel
	result := &HotelResponse{
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

func hotelShortResponseFromProto(hotel *hotelv1.HotelShort) *HotelShortResponse {
	if hotel == nil {
		return nil
	}

	resp := &HotelShortResponse{
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

func HotelsShortResponseFromProto(resp *hotelv1.GetHotelsResponse) *HotelsResponse {
	if resp == nil {
		return nil
	}

	hotels := make([]*HotelShortResponse, len(resp.Hotels))
	for i, h := range resp.Hotels {
		hotels[i] = hotelShortResponseFromProto(h)
	}

	return &HotelsResponse{Hotels: hotels}
}

func UpdateHotelResponseFromProto(hotel *hotelv1.UpdateHotel) *HotelResponse {
	if hotel == nil {
		return nil
	}

	result := &HotelResponse{
		Description: hotel.Description,
		Address:     hotel.Address,
		Location:    locationDTOFromProto(hotel.Location),
	}
	if hotel.Location != nil {
		result.Location = locationDTOFromProto(hotel.Location)
	}

	return result
}

func UpdateHotelTitleResponseFromProto(hotel *hotelv1.UpdateHotelTitle) *HotelResponse {
	if hotel == nil {
		return nil
	}

	return &HotelResponse{
		Title:     hotel.Title,
		HotelSlug: hotel.HotelSlug,
	}
}
