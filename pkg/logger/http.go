package logger

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func HttpLogger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		logrus.Infof("%s - :remote-user \"%s %s %s\" :status %d \":referrer\" \"%s\" - %s",
			r.Host,
			r.Method,
			r.RequestURI,
			r.Proto,
			r.ContentLength,
			r.Header.Get("User-Agent"),
			time.Since(start),
		)
	})
}
