package payment

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"oppapi/internal/logging"
	"oppapi/internal/model"
	"oppapi/internal/repository"
)

var (
	testDatabase   *repository.TestDatabase
	testRepository *PaymentRepository
)

func setup() {
	err := logging.Init("testing", "DEBUG", "text")
	if err != nil {
		panic(err)
	}
	testDatabase = repository.SetupTestDatabase()
	repository.InitWithMongo("testdb", testDatabase.DbAddress)
	_, err = repository.WithMongo()
	if err != nil {
		panic(err)
	}
	testRepository, err = NewPaymentRepository()
	if err != nil {
		panic(err)
	}
}

func tearDown() {
	testDatabase.TearDown()
}

// TestMain is the main entry point for testing and benchmarking.
func TestMain(m *testing.M) {
	setup()
	rc := m.Run()
	tearDown()
	os.Exit(rc)
}

func TestPaymentCreate(t *testing.T) {
	payment := model.Payment{
		Amount:   "100.00",
		Currency: "USD",
		Method: model.PaymentMethod{
			model.CreditCard{
				NameAndSurname: "Good Customer",
				CardNo:          "1234567890",
				CCV:             "123",
				ExpiryDate:      "01/25",
			},
		},
	}
	rv, err := testRepository.Create(&payment)
	if assert.Nil(t, err) {
		assert.False(t, rv.ID.IsZero())
	}
}

func TestPaymentReadOne(t *testing.T) {
	payment := model.Payment{
		Amount:   "100.00",
		Currency: "USD",
		Method: model.PaymentMethod{
			model.CreditCard{
				NameAndSurname: "Good Customer",
				CardNo:          "1234567890",
				CCV:             "123",
				ExpiryDate:      "01/25",
			},
		},
	}
	rv1, err1 := testRepository.Create(&payment)
	require.Nil(t, err1)
	require.False(t, rv1.ID.IsZero())
	rv2, err2 := testRepository.ReadOne(rv1.ID.Hex())
	if assert.Nil(t, err2) {
		assert.Equal(t, rv1.ID.Hex(), rv2.ID.Hex())
	}
}
