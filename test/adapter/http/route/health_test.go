package route_test

import (
	"encoding/json"
	"go-base/internal/adapter/http/route"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	route.ProviderHealthRoute(router)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &response)

	assert.NoError(t, err)

	assert.Equal(t, "UP", response["status"])
	assert.Equal(t, "Service is running", response["message"])
}

func TestHealthRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	healthRoute := route.ProviderHealthRoute(router)

	assert.Implements(t, (*route.HealthRoute)(nil), healthRoute)
}
