package auth

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"

	"auth/internal/grpc/lib/interceptor"
)

func newGRPCServer(logger *slog.Logger, accessSecret string) *grpc.Server {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.FinishCall),
		logging.WithFieldsFromContext(
			func(ctx context.Context) logging.Fields {
				return logging.Fields{}
			},
		),
		logging.WithLevels(logging.DefaultServerCodeToLevel),
	}

	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptor.Logger(logger), opts...),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(interceptor.Logger(logger), opts...),
		),
		grpc.UnaryInterceptor(interceptor.AuthInterceptor(accessSecret)),
	)
}
