package middleware

import (
	"context"
	"log/slog"

	_ "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

func SlogInterceptor(logger *slog.Logger) grpczap.Logger {
	return grpczap.LoggerFunc(
		func(ctx context.Context, lvl grpczap.Level, msg string, fields ...any) {
			switch lvl {
			case grpczap.LevelError:
				logger.Error(msg, fields...)
			case grpczap.LevelWarn:
				logger.Warn(msg, fields...)
			default:
				logger.Info(msg, fields...)
			}
		},
	)
}
