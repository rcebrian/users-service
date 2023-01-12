package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nelkinda/health-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	componentID      = "test"
	componentVersion = "v1.0.0-beta"
)

func Test_NewHealthService_Success(t *testing.T) {
	service := NewHealthService(componentID, componentVersion)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", service.Handler)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	var got health.Health
	err := json.NewDecoder(res.Body).Decode(&got)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, health.Pass, got.Status)
	assert.Equal(t, componentID, got.ServiceID)
	assert.Equal(t, componentVersion, got.Version)
}
