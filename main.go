package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"auth_service/api/http/handler"
	"auth_service/api/http/router"
	"auth_service/internal/config"
	"auth_service/internal/database/postgres"
	"auth_service/internal/service"
)

func getProjectRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env is not found: %v", err)
	}

	rootPath := getProjectRoot()
	configFileName := os.Getenv("CONFIG_FILENAME")
	if configFileName == "" {
		configFileName = "local.yaml"
	}

	configPath := filepath.Join(rootPath, "internal", "config", configFileName)
	newConfig, err := config.New(configPath)
	if err != nil {
		panic(err.Error())
	}

	userRepo, err := postgres.NewRepository(newConfig)
	if err != nil {
		panic(err.Error())
	}

	userService := service.New(userRepo)
	userHandler := handler.New(userService)
	routerHandlers := router.NewHandlers(userHandler)

	r := chi.NewRouter()
	router.New(r, routerHandlers)

	server := fmt.Sprintf("%s:%d", newConfig.Server.Host, newConfig.Server.Port)
	fmt.Printf("Starting server on %s\n", server)
	if err = http.ListenAndServe(server, r); err != nil {
		panic(err.Error())
	}

	panic("unreachable")
}
