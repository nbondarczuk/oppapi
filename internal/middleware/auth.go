package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"oppapi/internal/config"
	"oppapi/internal/logging"
)

// Auth checks X-API-KEY in the request header
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") != config.AuthXAPIKey() {
			logging.Logger.Info("Unauthorized client, aborting",
				slog.String("X-API-KEY", c.GetHeader("X-API-KEY")),
				slog.String("AuthXAPIKey", config.AuthXAPIKey()))
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		logging.Logger.Info("Client authorized",
			slog.String("X-API-KEY", c.GetHeader("X-API-KEY")))
	}
}
