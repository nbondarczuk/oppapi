package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"oppapi/internal/config"
	"oppapi/internal/handler"
	"oppapi/internal/logging"
	"oppapi/internal/model"
	"oppapi/internal/repository"
	paymentrepository "oppapi/internal/repository/payment"
)

var (
	testDatabase *repository.TestDatabase
	pattern      = `application:
  name: oppapi3
server:
  http:
    address: localhost3
    port: 80903
log:
  level: DEBUG3
  format: text3
repsitory:
  dbname: mongo
  url: mongodb://localhost:27017
auth:
  x_api_key: something
bank:
  url: test
`
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

	setupPaymentTestRouter()
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

func setupPaymentTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(handler.TestRequestLogger())
	r.Use(handler.TestResponseLogger())
	r.POST("/payment", CreatePaymentHandler)
	r.GET("/payment/:id", ReadOnePaymentHandler)
	return r
}

func TestCreatePaymentWithEmptyPayload(t *testing.T) {
	// Prepare
	config.MakeTestConfigFile(t, pattern)
	config.Init()
	defer config.CleanupTestConfigFile(t)
	r := setupPaymentTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/payment", nil)
	req.Header.Set("Content-Type", "application/json")

	// Run
	r.ServeHTTP(w, req)

	// Check
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreatePaymentWithInvalidPayload(t *testing.T) {
	// Prepare
	config.MakeTestConfigFile(t, pattern)
	config.Init()
	defer config.CleanupTestConfigFile(t)
	r := setupPaymentTestRouter()
	w := httptest.NewRecorder()
	var payload = []byte("abracadabra")
	req, _ := http.NewRequest("POST", "/payment", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Run
	r.ServeHTTP(w, req)

	// Check
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

type PaymentRequestResult struct {
	Payment     model.Payment
	Status      string
	Transaction model.Transaction
}

func TestCreatePayment(t *testing.T) {
	// Prepare
	var body = `{
  "amount": "100.00",
  "currency": "USD",
  "method": {
    "creditcard": {
      "nameandsurename": "Good Customer",
      "cardno": "1234567890",
      "ccv": "123",
      "expirtdate": "01/25"
    }
  }
}`

	config.MakeTestConfigFile(t, pattern)
	config.Init()
	defer config.CleanupTestConfigFile(t)
	r := setupPaymentTestRouter()
	w := httptest.NewRecorder()
	var payload = []byte(body)
	req, _ := http.NewRequest("POST", "/payment", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Run
	r.ServeHTTP(w, req)

	// Check
	assert.Equal(t, http.StatusOK, w.Code)
	response := w.Body.String()
	assert.NotNil(t, response)
	assert.NotZero(t, len(response))
	var result PaymentRequestResult
	json.Unmarshal([]byte(response), &result)
	logging.Logger.Debug("Result", slog.String("body", fmt.Sprintf("%+v", result)))
	assert.Equal(t, "OK", result.Status)
	assert.NotNil(t, result.Payment.ID)
	assert.Equal(t, "OK", result.Payment.Status)
	assert.Equal(t, "REGULAR", result.Payment.Type)
	assert.Equal(t, "100.00", result.Payment.Amount)
	assert.Equal(t, model.ISOCurrencyCode("USD"), result.Payment.Currency)
	assert.NotNil(t, result.Transaction.ID)
	assert.Equal(t, "100.00", result.Transaction.Amount)
	assert.Equal(t, model.ISOCurrencyCode("USD"), result.Transaction.Currency)
	assert.Equal(t, "OK", result.Transaction.Status)
	assert.Equal(t, result.Payment.ID, result.Transaction.ID)
}

func TestReadOnePaymentWithInvalidID(t *testing.T) {
	// Prepare
	r := setupPaymentTestRouter()
	w := httptest.NewRecorder()
	str := "whatever"
	req, _ := http.NewRequest("GET", "/payment/"+str, nil)
	req.Header.Set("Content-Type", "application/json")

	// Run
	r.ServeHTTP(w, req)

	// Check
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestReadOnePayment(t *testing.T) {
	// Prepare
	var payment model.Payment = model.Payment{
		Type:     "REGULAR",
		Amount:   "100.00",
		Currency: "USD",
		Status:   "OK",
	}
	tc, _ := paymentrepository.NewPaymentRepository()
	pval, _ := tc.Create(&payment)
	r := setupPaymentTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/payment/"+pval.ID.Hex(), nil)
	logging.Logger.Debug("Request", slog.String("path", fmt.Sprintf("%s", "/payment/"+pval.ID.Hex())))
	req.Header.Set("Content-Type", "application/json")

	// Run
	r.ServeHTTP(w, req)

	// Check
	assert.Equal(t, http.StatusOK, w.Code)
	response := w.Body.String()
	assert.NotNil(t, response)
	assert.NotZero(t, len(response))
	var result PaymentRequestResult
	json.Unmarshal([]byte(response), &result)
	logging.Logger.Debug("Result", slog.String("body", fmt.Sprintf("%+v", result)))
	assert.Equal(t, "OK", result.Status)
	assert.NotNil(t, result.Payment.ID)
	assert.Equal(t, "OK", result.Payment.Status)
	assert.Equal(t, "REGULAR", result.Payment.Type)
	assert.Equal(t, "100.00", result.Payment.Amount)
	assert.Equal(t, model.ISOCurrencyCode("USD"), result.Payment.Currency)
}
