package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type operationResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// PanicRecovery prevent server shutdowns when a panic occurs
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				logrus.Errorf("panic recover: %+v", err)

				body := operationResponse{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
				res, _ := json.Marshal(body)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write(res)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
