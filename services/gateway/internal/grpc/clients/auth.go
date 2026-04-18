package clients

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
)

func newAuthClients(cfg *config.Config) (userv1.UserServiceClient, userv1.TokenServiceClient, *grpc.ClientConn, error) {
	authAddr := fmt.Sprintf("%s:%d", cfg.Auth.Host, cfg.Auth.Port)
	authConn, err := grpc.NewClient(authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}

	return userv1.NewUserServiceClient(authConn), userv1.NewTokenServiceClient(authConn), authConn, nil
}
