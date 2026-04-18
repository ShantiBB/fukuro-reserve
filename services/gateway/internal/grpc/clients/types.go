package clients

import (
	"google.golang.org/grpc"

	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	bookingv1 "github.com/ShantiBB/fukuro-reserve/services/booking/api/booking/v1"
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
