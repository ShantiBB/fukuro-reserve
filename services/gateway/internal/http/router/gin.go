package router

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	httpswagger "github.com/swaggo/http-swagger"

	_ "github.com/ShantiBB/fukuro-reserve/services/gateway/docs"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/config"
)

type RouteRegistrar interface {
	Register(r *gin.RouterGroup)
}

func New(r *gin.Engine, httpCfg config.HTTPConfig, corsCfg config.CORSConfig, routes ...RouteRegistrar) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(requestTimeoutMiddleware(time.Duration(httpCfg.RequestTimeoutSec) * time.Second))
	r.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     corsCfg.AllowedOrigins,
				AllowMethods:     corsCfg.AllowedMethods,
				AllowHeaders:     corsCfg.AllowedHeaders,
				ExposeHeaders:    corsCfg.ExposedHeaders,
				AllowCredentials: corsCfg.AllowCredentials,
				MaxAge:           time.Duration(corsCfg.MaxAge) * time.Second,
			},
		),
	)

	r.GET(httpCfg.HealthPath, func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	api := r.Group(httpCfg.APIPrefix)
	api.GET("/docs/swagger/*any", gin.WrapH(httpswagger.WrapHandler))
	for _, route := range routes {
		route.Register(api)
	}
}

func requestTimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	if timeout <= 0 {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
