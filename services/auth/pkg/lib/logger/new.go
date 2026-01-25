package logger

import (
	"log/slog"
	"os"
	"strings"
)

func New(env, level string) *slog.Logger {
	logLevel := parseLogLevel(level)

	if env == "local" || env == "dev" {
		return slog.New(&SimpleHandler{out: os.Stdout, level: logLevel})
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
}

func parseLogLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
