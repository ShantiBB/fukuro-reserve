package clients

import "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"

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
