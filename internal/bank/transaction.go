package bank

import (
	"oppapi/internal/repository/payment"
	"oppapi/internal/repository/transaction"
)

// Resolve does the clearing between Payment and Transaction using
// the payment details.
func Resolve(payment.Payment) (transaction.Transaction, error) {
	return transaction.Transaction{}, nil
}
