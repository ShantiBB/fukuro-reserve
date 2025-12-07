package auth

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"auth/internal/config"
	"auth/internal/http/handler"
	"auth/internal/http/lib/jwt"
	"auth/internal/http/router"
	"auth/internal/repository/postgres"
	"auth/internal/service"
)

type App struct {
	Config *config.Config
}

func (app *App) MustLoad() {
	tokenCredentials := jwt.TokenCredentials{
		AccessSecret:  app.Config.JWT.AccessSecret,
		RefreshSecret: app.Config.JWT.RefreshSecret,
		AccessTTL:     app.Config.JWT.AccessTTL,
		RefreshTTL:    app.Config.JWT.RefreshTTL,
	}

	userRepo, err := postgres.NewRepository(app.Config)
	if err != nil {
		panic(err.Error())
	}

	userService := service.New(userRepo, &tokenCredentials)
	userHandler := handler.New(userService)

	r := chi.NewRouter()
	router.New(r, userHandler, app.Config.JWT.AccessSecret)

	server := fmt.Sprintf("%s:%d", app.Config.Server.Host, app.Config.Server.Port)
	slog.Info("Starting server", "address", server)
	if err = http.ListenAndServe(server, r); err != nil {
		panic(err.Error())
	}
}
