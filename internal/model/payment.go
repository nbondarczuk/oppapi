package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Payment is the entity mnaged by the repository.
// swagger:model payment
type Payment struct {
	// the id of the payment
	// required: true
	ID primitive.ObjectID `json:"id" bson:"_id"`

	// the amount of the payment
	// required: true
	Amount string `json:"amount" bson:"amount"`

	// the currency of the amount of the payment
	// required: true
	Currency ISOCurrencyCode `json:"currency" bson:"currency"`

	// the payment methid, it may be Payment card
	Method PaymentMethod `json:"method" bson:"method"`

	// date of creation of the payment
	// required: true
	Created time.Time `json:"created" bson:"created"`
}
