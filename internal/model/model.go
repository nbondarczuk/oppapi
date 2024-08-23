package model

// swagger:model isocurrencycode
type ISOCurrencyCode [3]byte

// swagger:model creditcard
type CreditCard struct {
	NameAndSurname string `json:"nameandsurename" bson:"nameandsurename"`
	CardNo         string `json:"cardno" bson:"cardno"`
	CCV            string `json:"ccv" bson:"ccv"`
	ExpiryDate     string `json:"expirtdate" bson:"expirtdate"`
}

// swagger:model paymentmethod
type PaymentMethod struct {
	CreditCard `json:"creditcard" bson:"creditcard"`
}

