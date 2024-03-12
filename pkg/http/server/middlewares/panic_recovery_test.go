package middlewares

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func Test_Middleware_PanicRecovery_InternalServerError(t *testing.T) {
	logrus.SetOutput(io.Discard)

	mux := http.NewServeMux()
	nextHandler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		panic("foo")
	})

	mux.Handle("/", PanicRecovery(nextHandler))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	var got operationResponse
	_ = json.NewDecoder(res.Body).Decode(&got)

	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
	assert.Equal(t, operationResponse{Code: http.StatusInternalServerError, Message: "Internal Server Error"}, got)
}

func Test_Middleware_PanicRecovery_Ok(t *testing.T) {
	mux := http.NewServeMux()
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(nil)
	})

	mux.Handle("/", PanicRecovery(nextHandler))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
