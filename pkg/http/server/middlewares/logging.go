package middlewares

import (
	"net/http"
	"time"

	joonix "github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
	Size   int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *StatusRecorder) Write(buf []byte) (int, error) {
	r.Size = len(buf)
	return r.ResponseWriter.Write(buf)
}

// Logging middleware that logs all requests
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &StatusRecorder{ResponseWriter: w, Status: 200}
		next.ServeHTTP(recorder, r)

		req := &joonix.HTTPRequest{
			Request:      r,
			RequestSize:  r.ContentLength,
			Status:       recorder.Status,
			ResponseSize: int64(recorder.Size),
			Latency:      time.Since(start),
			LocalIP:      r.Host,
			RemoteIP:     r.RemoteAddr,
		}

		logrus.WithField("httpRequest", req).Info()
	})
}
