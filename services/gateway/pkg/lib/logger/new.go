package logger

import (
	"log/slog"
	"os"
)

func New(env, logLevel string) *slog.Logger {
	var logger *slog.Logger

	level := slog.LevelInfo
	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	opts := &slog.HandlerOptions{Level: level}

	if env == "local" {
		opts.AddSource = true
		logger = slog.New(NewPrettyHandler(os.Stdout, opts))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
	}

	return logger
}
