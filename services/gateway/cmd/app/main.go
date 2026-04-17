package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/app/gateway"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/pkg/lib/logger"
)

//	@title			Swagger Gateway API
//	@version		1.0
//	@description	REST API Gateway for Fukuro Reserve microservices.

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
		configPath = "services/gateway/config/local.yaml"
	}

	cfg, err := config.New(configPath)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	log := logger.New(cfg.Env, cfg.LogLevel)

	app, err := gateway.New(cfg, log)
	if err != nil {
		panic("failed to create gateway app: " + err.Error())
	}

	app.MustRun()
}
