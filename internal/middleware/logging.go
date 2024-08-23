package middleware

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"

	"oppapi/internal/logging"
)

// RequestLogger logs incoming request.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		line := fmt.Sprintf("%s %s %s %s",
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			latency,
		)
		logging.Logger.Info("Received", slog.String("request", line))
	}
}

// ResponseLogger logs response.
func ResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Next()
		line := fmt.Sprintf("%d %s %s",
			c.Writer.Status(),
			c.Request.Method,
			c.Request.RequestURI,
		)
		logging.Logger.Info("Produced", slog.String("response", line))
	}
}
