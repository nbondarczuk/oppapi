package payment

import "errors"

var (
	ErrEmptyPaymentId = errors.New("an empty payment id was provided")
)
