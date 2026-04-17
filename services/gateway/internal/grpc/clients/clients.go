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
	// Auth service connection
	authAddr := fmt.Sprintf("%s:%d", cfg.Auth.Host, cfg.Auth.Port)
	authConn, err := grpc.NewClient(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}

	// Hotel service connection
	hotelAddr := fmt.Sprintf("%s:%d", cfg.Hotel.Host, cfg.Hotel.Port)
	hotelConn, err := grpc.NewClient(hotelAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		_ = authConn.Close()
		return nil, fmt.Errorf("failed to connect to hotel service: %w", err)
	}

	// Booking service connection
	bookingAddr := fmt.Sprintf("%s:%d", cfg.Booking.Host, cfg.Booking.Port)
	bookingConn, err := grpc.NewClient(bookingAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		_ = authConn.Close()
		_ = hotelConn.Close()
		return nil, fmt.Errorf("failed to connect to booking service: %w", err)
	}

	return &Clients{
		User:        userv1.NewUserServiceClient(authConn),
		Token:       userv1.NewTokenServiceClient(authConn),
		Hotel:       hotelv1.NewHotelServiceClient(hotelConn),
		Room:        hotelv1.NewRoomServiceClient(hotelConn),
		Booking:     bookingv1.NewBookingServiceClient(bookingConn),
		authConn:    authConn,
		hotelConn:   hotelConn,
		bookingConn: bookingConn,
	}, nil
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
