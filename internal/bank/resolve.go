package bank

import (
	"log/slog"
	"time"

	resty "github.com/go-resty/resty/v2"

	"oppapi/internal/logging"
	"oppapi/internal/model"
)

// Mocking merchnat
var merchant model.MerchantParty = model.MerchantParty{
	Name: "Merchan's name",
	Address: model.AddressDetails{
		Line1: "Address line 1",
		Line2: "Address line 2",
		Line3: "Address line 3",
	},
	Token: "fake-token",
	Identification: model.BankDetails{
		Name: "Merchant's name",
		IBAN: "PL1234567890",
	},
}

// Resolve does the clearing between Payment and Transaction using
// the payment details.
func Resolve(payment model.Payment) (model.Transaction, error) {
	logging.Logger.Info("Resolving payment",
		slog.String("payment.ID", string(payment.ID.Hex())))
	// Charging for the payment with transaction
	t := model.Transaction{
		// To make easy link between transaction and payment
		ID:       payment.ID,
		// Charged by full amount
		Amount:   payment.Amount,
		Currency: payment.Currency,
		// Using mocked merchant selling goods to the customer
		Merchant: merchant,
		// Using payment method to identify the customer
		Customer: model.CustomerParty{
			Identification: payment.Method,
		},
		Status:  "PENDING",
		Created: time.Now(),
	}
	// Hit the bank mock with payload of a transaction
	client := resty.New()
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(merchant.Token).
		Post("http://localhost:8080/transaction")
	if err != nil {
		t.Status = "ERROR"
	} else {
		t.Status = "OK"
	}
	logging.Logger.Info("Resolved payment",
		slog.String("payment.ID", string(payment.ID.Hex())),
		slog.String("transaction.ID", string(t.ID.Hex())),
		slog.String("transaction.Status", t.Status))
	return t, nil
}
