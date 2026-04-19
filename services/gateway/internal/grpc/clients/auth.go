package clients

import (
	"google.golang.org/grpc"

	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
)

func newAuthClients(cfg *config.Config) (userv1.UserServiceClient, userv1.TokenServiceClient, *grpc.ClientConn, error) {
	authConn, err := newConn(cfg.Auth.Host, cfg.Auth.Port, "auth")
	if err != nil {
		return nil, nil, nil, err
	}

	return userv1.NewUserServiceClient(authConn), userv1.NewTokenServiceClient(authConn), authConn, nil
}
