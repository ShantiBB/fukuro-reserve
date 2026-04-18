package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpswagger "github.com/swaggo/http-swagger"

	_ "github.com/ShantiBB/fukuro-reserve/services/gateway/docs"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
)

func New(r chi.Router, httpCfg config.HTTPConfig, corsCfg config.CORSConfig, routes ...RouteRegistrar) {
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Timeout(time.Duration(httpCfg.RequestTimeoutSec) * time.Second))

	r.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   corsCfg.AllowedOrigins,
				AllowedMethods:   corsCfg.AllowedMethods,
				AllowedHeaders:   corsCfg.AllowedHeaders,
				ExposedHeaders:   corsCfg.ExposedHeaders,
				AllowCredentials: corsCfg.AllowCredentials,
				MaxAge:           corsCfg.MaxAge,
			},
		),
	)

	// Health check
	r.Get(
		httpCfg.HealthPath, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	)

	// API routes
	r.Route(
		httpCfg.APIPrefix, func(r chi.Router) {
			r.Get("/docs/swagger/*", httpswagger.WrapHandler)
			for _, route := range routes {
				route.Register(r)
			}
		},
	)
}
