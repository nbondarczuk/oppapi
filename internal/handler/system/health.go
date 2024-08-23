package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler provides simple feedback about living server.
func HealthHandler(c *gin.Context) {
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}
