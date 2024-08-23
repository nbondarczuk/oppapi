package payment

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"oppapi/internal/model"
	repository "oppapi/internal/repository/payment"
)

// CreateHandler creates a new payment.
//
// swagger:operation POST /payment payment CreateHandler
// Creates a new payment.
// ---
// produces:
//  - application/json
//
// responses:
//   '200':
//	   description: OK
//   '400':
//	   description: Bad Request
//   '500':
//	   description: Internal Server Error
//
func CreateHandler(c *gin.Context) {
	var payment model.Payment
	// Check input ie. new object attributes from request body.
	if err := c.ShouldBindJSON(&payment); err != nil {
		// Handle error in request body.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// The controller gives access to particular collection.
	tc, err := repository.NewPaymentRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rval, err := tc.Create(&payment)
	if err != nil {
		// Handle error in object creation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}

// ReadHOneHandler reads one payment by id.
//
// swagger:operation GET /payment/{id} payment ReadOneHandler
// Reads one one payment by id.
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the payment
//     required: true
//     type: string
// produces:
// - application/json
// responses:
//   '200':
//	   description: OK
//   '400':
//	   description: Bad Request
//   '404':
//     description: Not Found
//   '500':
//	   description: Internal Server Error
//
func ReadOneHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty payment id provided"})
		return
	}
	// The controlle gives access to particular collection.
	tc, err := repository.NewPaymentRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rval, err := tc.ReadOne(id)
	if err != nil {
		// Handle error in repository read operation.
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status": "Ok",
		"Object": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}
