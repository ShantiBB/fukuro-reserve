package clients

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	hotelv1 "github.com/ShantiBB/fukuro-reserve/services/hotel/api/hotel/v1"
)

type Clients struct {
	User    userv1.UserServiceClient
	Token   userv1.TokenServiceClient
	Hotel   hotelv1.HotelServiceClient
	Room    hotelv1.RoomServiceClient
	Booking bookingv1.BookingServiceClient

	authConn    *grpc.ClientConn
	hotelConn   *grpc.ClientConn
	bookingConn *grpc.ClientConn
}

func New(cfg *config.Config) (*Clients, error) {
	userClient, tokenClient, authConn, err := newAuthClients(cfg)
	if err != nil {
		return nil, err
	}

	hotelClient, roomClient, hotelConn, err := newHotelClients(cfg)
	if err != nil {
		_ = authConn.Close()
		return nil, err
	}

	bookingClient, bookingConn, err := newBookingClient(cfg)
	if err != nil {
		_ = authConn.Close()
		_ = hotelConn.Close()
		return nil, err
	}

	return &Clients{
		User:        userClient,
		Token:       tokenClient,
		Hotel:       hotelClient,
		Room:        roomClient,
		Booking:     bookingClient,
		authConn:    authConn,
		hotelConn:   hotelConn,
		bookingConn: bookingConn,
	}, nil
}

func newConn(host string, port int, serviceName string) (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s service: %w", serviceName, err)
	}

	return conn, nil
}

func (c *Clients) Close() error {
	var errs []error
	if err := c.authConn.Close(); err != nil {
		errs = append(errs, err)
	}
	if err := c.hotelConn.Close(); err != nil {
		errs = append(errs, err)
	}
	if err := c.bookingConn.Close(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return fmt.Errorf("errors closing connections: %v", errs)
	}

	return nil
}
