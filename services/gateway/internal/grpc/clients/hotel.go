package clients

import (
	"google.golang.org/grpc"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

func newHotelClients(cfg *config.Config) (hotelv1.HotelServiceClient, hotelv1.RoomServiceClient, *grpc.ClientConn, error) {
	hotelConn, err := newConn(cfg.Hotel.Host, cfg.Hotel.Port, "hotel")
	if err != nil {
		return nil, nil, nil, err
	}

	return hotelv1.NewHotelServiceClient(hotelConn), hotelv1.NewRoomServiceClient(hotelConn), hotelConn, nil
}
