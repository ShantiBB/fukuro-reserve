package logger

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startedAt := time.Now()
		requestPath := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(startedAt)
		statusCode := c.Writer.Status()

		logAttrs := []any{
			"method", c.Request.Method,
			"path", requestPath,
			"status", statusCode,
			"latency", latency,
			"client_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"response_size", c.Writer.Size(),
		}
		if rawQuery != "" {
			logAttrs = append(logAttrs, "query", rawQuery)
		}
		if len(c.Errors) > 0 {
			logAttrs = append(logAttrs, "errors", c.Errors.String())
		}

		msg := "HTTP request"
		switch {
		case statusCode >= http.StatusInternalServerError:
			slog.ErrorContext(c.Request.Context(), msg, logAttrs...)
		case statusCode >= http.StatusBadRequest:
			slog.WarnContext(c.Request.Context(), msg, logAttrs...)
		default:
			slog.InfoContext(c.Request.Context(), msg, logAttrs...)
		}
	}
}

func RecoveryWithLogger() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		slog.ErrorContext(
			c.Request.Context(),
			"panic recovered",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"client_ip", c.ClientIP(),
			"panic", fmt.Sprint(recovered),
			"stack", string(debug.Stack()),
		)

		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
