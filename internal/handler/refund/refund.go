package refund

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"oppapi/internal/bank"
	"oppapi/internal/model"
	repository "oppapi/internal/repository/payment"
)

// CreateHandler creates a new refund using reference to the original payment.
//
// swagger:operation POST /refund/{id} refund CreateHandler
// Creates a new refund using reference to the original payment.
// ---
// produces:
//   - application/json
//
// responses:
//
//	  '200':
//		   description: OK
//	  '400':
//		   description: Bad Request
//	  '500':
//		   description: Internal Server Error
func CreateHandler(c *gin.Context) {
	// id is the reference to the original payment
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty payment id provided"})
		return
	}
	// The controller gives access to particular collection.
	tc, err := repository.NewPaymentRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Refund refers to a previous payment & transaction.
	payment, err := tc.ReadOne(id)
	if err != nil {
		// Handle error in repository read operation.
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var refund model.Payment = model.Payment{
		Type:     "REFUND",
		Amount:   payment.Amount,
		Currency: payment.Currency,
		Method:   payment.Method,
		Status:   "PENDING",
	}
	pval, err := tc.Create(&refund)
	if err != nil {
		// Handle error in object creation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tval, err := bank.Resolve(refund)
	if err != nil {
		// Handle error in refund resolution by marchant's bank.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pval.Status = "OK"
	err = tc.SetStatus(pval.ID.Hex(), pval.Status)
	if err != nil {
		// Handle error payment update.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status":      tval.Status,
		"Refund":      pval,
		"Transaction": tval,
	}
	c.JSON(http.StatusOK, r)
	return
}

// ReadHOneHandler reads one refund by id.
//
// swagger:operation GET /refund/{id} refund ReadOneHandler
// Reads one one refund by id.
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the refund
//     required: true
//     type: string
//
// produces:
// - application/json
// responses:
//
//	  '200':
//		   description: OK
//	  '400':
//		   description: Bad Request
//	  '404':
//	    description: Not Found
//	  '500':
//		   description: Internal Server Error
func ReadOneHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty refund id provided"})
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
	if rval.Type != "REFUND" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type, expected REFUND, found: " + rval.Type})
		return
	}
	r := map[string]interface{}{
		"Status": "Ok",
		"Refund": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}
