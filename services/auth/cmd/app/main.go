package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"auth/internal/app/auth"
	"auth/internal/config"
)

//	@title			Swagger Auth API
//	@version		1.0
//	@description	Auth service for microservices.

//	@host		localhost:8081
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
		panic("CONFIG_PATH is not set")
	}

	cfg, err := config.New(configPath)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	authApp := auth.App{Config: cfg}
	authApp.MustLoad()
}
