package clients

import (
	"google.golang.org/grpc"

	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
)

func newBookingClient(cfg *config.Config) (bookingv1.BookingServiceClient, *grpc.ClientConn, error) {
	bookingConn, err := newConn(cfg.Booking.Host, cfg.Booking.Port, "booking")
	if err != nil {
		return nil, nil, err
	}

	return bookingv1.NewBookingServiceClient(bookingConn), bookingConn, nil
}
