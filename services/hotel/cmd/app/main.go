package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"hotel/internal/app/hotel"
	"hotel/internal/config"
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
	if err := godotenv.Load(); err != nil {
		slog.Warn("failed load env", "error", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_FILENAME is not set")
	}

	cfg, err := config.New(configPath)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	hotelApp := hotel.App{Config: cfg}
	hotelApp.MustLoad()
}
