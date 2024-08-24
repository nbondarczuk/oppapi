package model

type ISOCurrencyCode string

type CreditCard struct {
	NameAndSurname string `json:"nameandsurename" bson:"nameandsurename"`
	CardNo         string `json:"cardno" bson:"cardno"`
	CCV            string `json:"ccv" bson:"ccv"`
	ExpiryDate     string `json:"expirtdate" bson:"expirtdate"`
}

type PaymentMethod struct {
	CreditCard `json:"creditcard" bson:"creditcard"`
}
