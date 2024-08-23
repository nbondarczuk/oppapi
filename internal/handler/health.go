package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler provides simple feedback about living server.
//
// swagger:operation GET /health health HealthHandler
// Check server status.
// ---
// produces:
// - application/json
// responses:
//   '200':
//	   description: OK
//   '500':
//	   description: Internal Server Error
//
func HealthHandler(c *gin.Context) {
	r := map[string]interface{}{
		"Status": "Ok",
	}
	c.JSON(http.StatusOK, r)
	return
}
