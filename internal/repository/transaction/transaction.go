package transaction

type Transaction struct {
	// the id of the transaction
	// required: true
	ID primitive.ObjectID `json:"id" bson:"_id"`

	// the amount of the payment
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
}
