package model

type CustomerParty struct {
	Identification PaymentMethod `json:"identification" bson:"identification"`
}
