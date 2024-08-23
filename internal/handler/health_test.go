package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupHealthTestRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(testRequestLogger())
	r.Use(testResponseLogger())
	r.GET("/health", HealthHandler)
	return r
}

func TestHealthRoute(t *testing.T) {
	r := setupHealthTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"Status\":\"Ok\"}", w.Body.String())
}
