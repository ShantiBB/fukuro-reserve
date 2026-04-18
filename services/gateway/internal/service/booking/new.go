package booking

import (
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/grpc/clients"
)

type Service struct {
	clients    *clients.Clients
	pagination config.PaginationConfig
}

func New(clients *clients.Clients, pagination config.PaginationConfig) *Service {
	return &Service{clients: clients, pagination: pagination}
}
