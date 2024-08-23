package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction sent to the bank of the merchant in order to resolve the payment.
// swagger:model transaction
type Transaction struct {
	// the id of the transaction
	// required: true
	ID primitive.ObjectID `json:"id" bson:"_id"`

	// the amount to be debited
	// required: true
	Amount string `json:"amount" bson:"amount"`

	// the currency of the amount of the payment
	// required: true
	Currency ISOCurrencyCode `json:"currency" bson:"currency"`

	// the status of the transaction after clearing
	// required: true
	Status string

	// date of creation of the payment
	// required: true
	Created time.Time `json:"created" bson:"created"`

	// merchant data used in bank account determination
	// required: true
	Merchant MerchantParty `json:"merchant" bson:"merchant"`

	// customer data used payment method indentifcation to be resolved by bank
	// required: true
	Customer CustomerParty `json:"customer" bson:"customer"`
}
