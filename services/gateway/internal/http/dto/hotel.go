package dto

import "time"

type LocationDTO struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type CreateHotelBody struct {
	Location    *LocationDTO `json:"location,omitempty"`
	Title       string       `json:"title"`
	Description string       `json:"description,omitempty"`
	Address     string       `json:"address"`
	OwnerId     int64        `json:"owner_id"`
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
