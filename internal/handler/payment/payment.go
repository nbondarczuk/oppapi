package payment

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"oppapi/internal/bank"
	"oppapi/internal/logging"
	"oppapi/internal/model"
	repository "oppapi/internal/repository/payment"
)

// CreatePaymentHandler creates a new payment.
//
// swagger:operation POST /payment payment CreatePaymentHandler
// Creates a new payment. The required fields in the request input payload
// are briefly validated. The new payment is created with a status PENDING.
// The type of the payment is REGULAR. The payment stored as it is in a collection.
// The transaction is issued to be bank of the merchant so that the payment can be cleared.
// The result of the transaction is stored with the payment status
// as the collection record gets updated. The payment and the transaction details
// are returned to be client plus status code of the whole operation.
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
func CreatePaymentHandler(c *gin.Context) {
	var payment model.Payment
	// Check input ie. new object attributes from request body.
	if err := c.ShouldBindJSON(&payment); err != nil {
		// Handle error in request body.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logging.Logger.Debug("Payment create",
		slog.String("Amount", payment.Amount),
		slog.String("Currency", string(payment.Currency)))
	// The controller gives access to particular collection.
	tc, err := repository.NewPaymentRepository()
	if err != nil {
		// Handle error in repository allocation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	payment.Type = "REGULAR"
	payment.Status = "PENDING"
	pval, err := tc.Create(&payment)
	if err != nil {
		// Handle error in object creation.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tval, err := bank.Resolve(payment)
	if err != nil {
		pval.Status = fmt.Sprintf("ERROR: %v", err)
	} else {
		pval.Status = "OK"
	}
	err = tc.SetStatus(pval.ID.Hex(), pval.Status)
	if err != nil {
		// Handle error payment update
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	r := map[string]interface{}{
		"Status":      tval.Status,
		"Payment":     pval,
		"Transaction": tval,
	}
	c.JSON(http.StatusOK, r)
	return
}

// ReadHOnePaymentHandler reads one payment by id.
//
// swagger:operation GET /payment/{id} payment ReadOnePaymentHandler
// Reads one payment by id. This one must be provided in the path. An error
// is returned if it is not provided. The repository is queried for a given
// payment id. An error is retuned if it is not found. After successful
// read operation in the repository the payment details is retned with a status code
// of he whole operation.
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the payment
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
func ReadOnePaymentHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty payment id provided"})
		return
	}
	logging.Logger.Debug("Payment read", slog.String("id", id))
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
		"Status":  "OK",
		"Payment": rval,
	}
	c.JSON(http.StatusOK, r)
	return
}
