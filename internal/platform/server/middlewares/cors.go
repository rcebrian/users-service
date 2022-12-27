package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

// Cors allow all cors
func Cors(next http.Handler) http.Handler {
	var corsHandler = cors.AllowAll().Handler

	return corsHandler(next)
}
