package clients

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

func newHotelClients(cfg *config.Config) (hotelv1.HotelServiceClient, hotelv1.RoomServiceClient, *grpc.ClientConn, error) {
	hotelAddr := fmt.Sprintf("%s:%d", cfg.Hotel.Host, cfg.Hotel.Port)
	hotelConn, err := grpc.NewClient(hotelAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to connect to hotel service: %w", err)
	}

	return hotelv1.NewHotelServiceClient(hotelConn), hotelv1.NewRoomServiceClient(hotelConn), hotelConn, nil
}
