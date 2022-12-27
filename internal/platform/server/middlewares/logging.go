package middlewares

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

// Logging middleware that logs all requests
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &StatusRecorder{ResponseWriter: w, Status: 200}
		next.ServeHTTP(recorder, r)
		logrus.Infof("%s \"%s %s %s\" %d %d \"%s\" %s",
			r.RemoteAddr,
			r.Method, r.RequestURI, r.Proto,
			recorder.Status,
			r.ContentLength,
			r.Header.Get("User-Agent"),
			time.Since(start),
		)
	})
}
