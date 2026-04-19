package dto

import "time"

type LocationDTO struct {
	Latitude  float32 `json:"latitude" example:"35.66"`
	Longitude float32 `json:"longitude" example:"139.70"`
}

type CreateHotelBody struct {
	Location    *LocationDTO `json:"location,omitempty"`
	Title       string       `json:"title" example:"Imperial Hotel Tokyo"`
	Description string       `json:"description,omitempty" example:"Luxury hotel in central Tokyo"`
	Address     string       `json:"address" example:"1 Chiyoda, Tokyo"`
	OwnerId     int64        `json:"owner_id" example:"1"`
}

type CreateHotelRequest struct {
	Location    *LocationDTO
	CountryCode string
	CitySlug    string
	Title       string
	Description string
	Address     string
	OwnerId     int64
}

type UpdateHotelRequest struct {
	Location    *LocationDTO `json:"location,omitempty"`
	Description string       `json:"description" example:"Renovated rooms and improved amenities"`
	Address     string       `json:"address" example:"2 Chiyoda, Tokyo"`
}

type UpdateHotelTitleRequest struct {
	Title string `json:"title" example:"Imperial Hotel Tokyo Annex"`
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
