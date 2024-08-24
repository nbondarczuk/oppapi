package model

// swagger:model customerparty
type CustomerParty struct {
	Identification PaymentMethod `json:"identification" bson:"identification"`
}
