package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/repository/models"
)

func locationResponseToProto(l *models.Location) *hotelv1.Location {
	return &hotelv1.Location{
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
	}
}

func CreateHotelResponseToProto(resp *models.Hotel) *hotelv1.CreateHotel {
	return &hotelv1.CreateHotel{
		Id:          resp.ID.String(),
		Slug:        resp.Slug,
		Title:       resp.Title,
		OwnerId:     resp.OwnerID,
		Description: *resp.Description,
		Address:     resp.Address,
		Location:    locationResponseToProto(&resp.Location),
		CreatedAt:   timestamppb.New(resp.CreatedAt),
		UpdatedAt:   timestamppb.New(resp.UpdatedAt),
	}
}

func HotelResponseToProto(resp *models.Hotel) *hotelv1.Hotel {
	return &hotelv1.Hotel{
		Id:          resp.ID.String(),
		Title:       resp.Title,
		OwnerId:     resp.OwnerID,
		Description: *resp.Description,
		Address:     resp.Address,
		Rating:      resp.Rating,
		Location:    locationResponseToProto(&resp.Location),
		CreatedAt:   timestamppb.New(resp.CreatedAt),
		UpdatedAt:   timestamppb.New(resp.UpdatedAt),
	}
}

func HotelShortResponseToProto(resp *models.HotelShort) *hotelv1.HotelShort {
	return &hotelv1.HotelShort{
		Id:       resp.ID.String(),
		Title:    resp.Title,
		Slug:     resp.Slug,
		OwnerId:  resp.OwnerID,
		Rating:   resp.Rating,
		Address:  resp.Address,
		Location: locationResponseToProto(&resp.Location),
	}
}

func UpdateHotelResponseToProto(resp models.UpdateHotel) *hotelv1.UpdateHotel {
	return &hotelv1.UpdateHotel{
		Description: *resp.Description,
		Address:     resp.Address,
		Location:    locationResponseToProto(&resp.Location),
	}
}

func UpdateHotelTitleResponseToProto(resp models.UpdateHotelTitle) *hotelv1.UpdateHotelTitle {
	return &hotelv1.UpdateHotelTitle{
		Title: resp.Title,
		Slug:  resp.Slug,
	}
}

func HotelsResponseToProto(resp []*models.HotelShort) []*hotelv1.HotelShort {
	hotels := make([]*hotelv1.HotelShort, len(resp))
	for i, h := range resp {
		hotels[i] = HotelShortResponseToProto(h)
	}

	return hotels
}
