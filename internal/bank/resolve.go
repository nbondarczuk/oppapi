package bank

import (
	"fmt"
	"log/slog"
	"time"

	resty "github.com/go-resty/resty/v2"

	"oppapi/internal/config"
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
	var amount string = payment.Amount
	// In case of REFUND the reverse transaction is done (up to the bank interface)
	if payment.Type == "REFUND" {
		amount = "-" + amount
	}
	t := model.Transaction{
		// To make easy link between transaction and payment
		ID:       payment.ID,
		// Charged by full amount
		Amount:   amount,
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
	// Test mode: no requests
	if config.BankURL() == "test" {
		logging.Logger.Info("Skip request due to test mode")
		t.Status = "OK"
		return t, nil
	}
	// Hit the bank interface with payload of a transaction
	logging.Logger.Info("Sending request to bank", slog.String("url", config.BankURL()))
	client := resty.New()
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(t.Merchant.Token).
		Post(config.BankURL())
	if err != nil {
		t.Status = fmt.Sprintf("ERROR: %v", err)
	} else {
		t.Status = "OK"
	}
	// Anonymize after use not to reveal in transaction record returned from api.
	t.Merchant.Token = ""
	logging.Logger.Info("Resolved payment with transaction",
		slog.String("payment.ID", string(payment.ID.Hex())),
		slog.String("transaction.ID", string(t.ID.Hex())),
		slog.String("transaction.Status", t.Status))
	return t, nil
}
