package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Payment is the entity mnaged by the repository.
// swagger:parameters CreatePaymentHandler
type Payment struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Type     string             `json:"type" bson:"type"`
	Amount   string             `json:"amount" binding:"required" bson:"amount"`
	Currency ISOCurrencyCode    `json:"currency" binding:"required" bson:"currency"`
	Method   PaymentMethod      `json:"method" bson:"method"`
	Status   string             `json:"status" bson:"status"`
	Created  time.Time          `json:"created" bson:"created"`
	Modified time.Time          `json:"modified" bson:"modified"`
}
