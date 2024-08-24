package model

type AddressDetails struct {
	Line1 string `json:"line1" bson:"line1"`
	Line2 string `json:"line2" bson:"line2"`
	Line3 string `json:"line3" bson:"line3"`
}

type BankDetails struct {
	Name      string `json:"name" bson:"name"`
	IBAN      string `json:"iban" bson:"iban"`
}

type MerchantParty struct {
	Name           string         `json:"name" bson:"name"`
	Address        AddressDetails `json:"address" bson:"address"`
	Token          string         `json:"token" bson:"token"`
	Identification BankDetails    `json:"bank" bson:"bank"`
}
