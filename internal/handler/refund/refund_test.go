package refund

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

func setupRefundTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(handler.TestRequestLogger())
	r.Use(handler.TestResponseLogger())
	r.POST("/refund/:id", CreateRefundHandler)
	r.GET("/refund/:id", ReadOneRefundHandler)
	return r
}

type RefundRequestResult struct {
	Status      string
	Refund      model.Payment
	Transaction model.Transaction
}

func TestReadOneRefundWithInvalidID(t *testing.T) {
	// Prepare
	r := setupRefundTestRouter()
	w := httptest.NewRecorder()
	str := "whatever"
	req, _ := http.NewRequest("GET", "/refund/"+str, nil)
	req.Header.Set("Content-Type", "application/json")

	// Run
	r.ServeHTTP(w, req)

	// Check
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestReadOneRefund(t *testing.T) {
	// Prepare
	var refund model.Payment = model.Payment{
		Type:     "REFUND",
		Amount:   "100.00",
		Currency: "USD",
		Status:   "OK",
	}
	tc, _ := paymentrepository.NewPaymentRepository()
	pval, _ := tc.Create(&refund)
	r := setupRefundTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/refund/"+pval.ID.Hex(), nil)
	logging.Logger.Debug("Request", slog.String("path", fmt.Sprintf("%s", "/refund/"+pval.ID.Hex())))
	req.Header.Set("Content-Type", "application/json")

	// Run
	r.ServeHTTP(w, req)

	// Check
	assert.Equal(t, http.StatusOK, w.Code)
	response := w.Body.String()
	assert.NotNil(t, response)
	assert.NotZero(t, len(response))
	var result RefundRequestResult
	json.Unmarshal([]byte(response), &result)
	logging.Logger.Debug("Result", slog.String("body", fmt.Sprintf("%+v", result)))
	assert.Equal(t, "OK", result.Status)
	assert.NotNil(t, result.Refund.ID)
	assert.Equal(t, "OK", result.Refund.Status)
	assert.Equal(t, "REFUND", result.Refund.Type)
	assert.Equal(t, "100.00", result.Refund.Amount)
	assert.Equal(t, model.ISOCurrencyCode("USD"), result.Refund.Currency)
}

func TestCreateRefundWithInvalidPayload(t *testing.T) {
	// Prepare
	r := setupRefundTestRouter()
	w := httptest.NewRecorder()
	var payload = []byte("abracadabra")
	req, _ := http.NewRequest("POST", "/refund", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Run
	r.ServeHTTP(w, req)

	// Check
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateRefund(t *testing.T) {
	// Prepare
	config.MakeTestConfigFile(t, pattern)
	config.Init()
	defer config.CleanupTestConfigFile(t)
	var payment model.Payment = model.Payment{
		Type:     "REFUND",
		Amount:   "100.00",
		Currency: "USD",
		Status:   "OK",
	}
	tc, _ := paymentrepository.NewPaymentRepository()
	pval, _ := tc.Create(&payment)
	r := setupRefundTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/refund/"+pval.ID.Hex(), nil)
	logging.Logger.Debug("Request", slog.String("path", fmt.Sprintf("%s", "/refund/"+pval.ID.Hex())))
	req.Header.Set("Content-Type", "application/json")

	// Run
	r.ServeHTTP(w, req)

	// Check
	assert.Equal(t, http.StatusOK, w.Code)
	response := w.Body.String()
	assert.NotNil(t, response)
	assert.NotZero(t, len(response))
	var result RefundRequestResult
	json.Unmarshal([]byte(response), &result)
	logging.Logger.Debug("Result", slog.String("body", fmt.Sprintf("%+v", result)))
	assert.Equal(t, "OK", result.Status)
	assert.NotNil(t, result.Refund.ID)
	assert.Equal(t, "OK", result.Refund.Status)
	assert.Equal(t, "REFUND", result.Refund.Type)
	assert.Equal(t, "100.00", result.Refund.Amount)
	assert.Equal(t, model.ISOCurrencyCode("USD"), result.Refund.Currency)
	assert.NotNil(t, result.Transaction.ID)
	assert.Equal(t, "-100.00", result.Transaction.Amount)
	assert.Equal(t, model.ISOCurrencyCode("USD"), result.Transaction.Currency)
	assert.Equal(t, "OK", result.Transaction.Status)
	assert.Equal(t, result.Refund.ID, result.Transaction.ID)
}
