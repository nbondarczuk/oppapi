package bank

import (
	"oppapi/internal/model"
)

// Resolve does the clearing between Payment and Transaction using
// the payment details.
func Resolve(model.Payment) (model.Transaction, error) {
	return model.Transaction{}, nil
}
