package clients

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
)

func newBookingClient(cfg *config.Config) (bookingv1.BookingServiceClient, *grpc.ClientConn, error) {
	bookingAddr := fmt.Sprintf("%s:%d", cfg.Booking.Host, cfg.Booking.Port)
	bookingConn, err := grpc.NewClient(bookingAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to booking service: %w", err)
	}

	return bookingv1.NewBookingServiceClient(bookingConn), bookingConn, nil
}
