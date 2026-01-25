package main

import (
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"hotel/internal/app/hotel"
	"hotel/internal/config"
	"hotel/pkg/lib/logger"
)

//	@title			Swagger Hotel API
//	@version		1.0
//	@description	Hotel service for microservices.

//	@host		localhost:8082
//	@BasePath	/api/v1

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token.
func main() {
	if err := cleanenv.ReadConfig(".env", &struct{}{}); err != nil {
		slog.Warn("failed to load env", "error", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH is not set")
	}

	cfg, err := config.New(configPath)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	log := logger.New(cfg.Env, cfg.LogLevel)
	hotelApp := hotel.App{
		Config: cfg,
		Logger: log,
	}
	hotelApp.MustLoadGRPC()
}
