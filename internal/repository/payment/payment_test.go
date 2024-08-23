package payment

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"

	"oppapi/internal/model"
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
	cleanAll()
}

func tearDown() {
	cleanAll()
	testDatabase.TearDown()
}

func cleanAll() {
	err := testRepository.Drop()
	if err != nil {
		panic(err)
	}
}

// TestMain is the main entry point for testing and benchmarking.
func TestMain(m *testing.M) {
	setup()
	rc := m.Run()
	tearDown()
	os.Exit(rc)
}

func TestPaymentCreate(t *testing.T) {
	cleanAll()
	payment := model.Payment{}
	rv, err := testRepository.Create(&payment)
	if assert.Nil(t, err) {
		assert.False(t, rv.ID.IsZero())
		assert.Equal(t, payment.Label, rv.Label)
		assert.Equal(t, payment.Color, rv.Color)
	}
}

func TestPaymentReadOne(t *testing.T) {
	cleanAll()
	payment := model.Payment{}
	rv1, err1 := testRepository.Create(&payment)
	require.Nil(t, err1)
	require.False(t, rv1.ID.IsZero())
	rv2, err2 := testRepository.ReadOne(rv1.ID.Hex())
	if assert.Nil(t, err2) {
		assert.False(t, rv2.ID.IsZero())
		assert.Equal(t, rv1.ID.Hex(), rv2.ID.Hex())
		assert.Equal(t, rv1.Label, rv2.Label)
		assert.Equal(t, rv1.Color, rv2.Color)
	}
}
