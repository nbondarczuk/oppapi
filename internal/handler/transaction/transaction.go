package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"oppapi/internal/model"
)

// CreateHandler mocks bank response with a successuful transaction.
func CreateHandler(c *gin.Context) {
	var transaction model.Transaction
	// Check input ie. new object attributes from request body.
	if err := c.ShouldBindJSON(&transaction); err != nil {
		// Handle error in request body.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status":      "OK",
		"Transaction": transaction,
	}
	c.JSON(http.StatusOK, r)
	return
}
