package service

import (
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
)

type Auth struct {
	clients    *clients.Clients
	pagination config.PaginationConfig
}

type Hotel struct {
	clients    *clients.Clients
	pagination config.PaginationConfig
}

type Booking struct {
	clients    *clients.Clients
	pagination config.PaginationConfig
}

func NewAuth(clients *clients.Clients, pagination config.PaginationConfig) *Auth {
	return &Auth{clients: clients, pagination: pagination}
}

func NewHotel(clients *clients.Clients, pagination config.PaginationConfig) *Hotel {
	return &Hotel{clients: clients, pagination: pagination}
}

func NewBooking(clients *clients.Clients, pagination config.PaginationConfig) *Booking {
	return &Booking{clients: clients, pagination: pagination}
}
