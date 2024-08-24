package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction sent to the bank of the merchant in order to resolve the payment.
type Transaction struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Amount   string             `json:"amount" bson:"amount"`
	Currency ISOCurrencyCode    `json:"currency" bson:"currency"`
	Merchant MerchantParty      `json:"merchant" bson:"merchant"`
	Customer CustomerParty      `json:"customer" bson:"customer"`
	Status   string             `json:"status" bson:"status"`
	Created  time.Time          `json:"created" bson:"created"`
}
