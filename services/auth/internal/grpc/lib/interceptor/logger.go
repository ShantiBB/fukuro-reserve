package interceptor

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

func Logger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(
		func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
			filtered := make([]any, 0, len(fields))
			for i := 0; i < len(fields); i += 2 {
				if i+1 >= len(fields) {
					break
				}
				key := fields[i].(string)

				switch key {
				case "grpc.service", "grpc.method", "grpc.code", "grpc.time_ms":
					filtered = append(filtered, key, fields[i+1])
				}
			}

			l.Log(ctx, slog.Level(lvl), msg, filtered...)
		},
	)
}
