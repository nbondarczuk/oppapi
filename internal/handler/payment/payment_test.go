package payment

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"oppapi/internal/config"
	"oppapi/internal/handler"
	"oppapi/internal/logging"
	"oppapi/internal/repository"
)

var (
	testDatabase   *repository.TestDatabase
	pattern = `application:
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
	return r
}

func TestCreatePaymentRouteWithEmptyPayload(t *testing.T) {
	config.MakeTestConfigFile(t, pattern)
	config.Init()
	defer config.CleanupTestConfigFile(t)
	r := setupPaymentTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/payment", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreatePaymentRouteWithInvalidPayload(t *testing.T) {
	config.MakeTestConfigFile(t, pattern)
	config.Init()
	defer config.CleanupTestConfigFile(t)
	r := setupPaymentTestRouter()
	w := httptest.NewRecorder()
	var payload = []byte("abracadabra")
	req, _ := http.NewRequest("POST", "/payment", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}


var request = `{
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

func TestCreatePaymentRouteWith(t *testing.T) {
	config.MakeTestConfigFile(t, pattern)
	config.Init()
	defer config.CleanupTestConfigFile(t)
	r := setupPaymentTestRouter()
	w := httptest.NewRecorder()
	var payload = []byte(request)
	req, _ := http.NewRequest("POST", "/payment", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
