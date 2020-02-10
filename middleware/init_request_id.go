package middleware

//http middleware
//cmd/calc/http.goから呼び出す

import (
	"context"
	"net/http"
)

const RequestIDKey = "REQUEST-ID-KEY"

// InitRequestID is a HTTP server middleware that reads the value of
// the X-Request-Id header and if present writes it in the request context
func InitRequestID() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		// A HTTP handler is a function.
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := r

			// Grab X-Request-Id header and initialize request context with it
			if id := r.Header.Get("X-Request-Id"); id != "" {
				ctx := context.WithValue(r.Context(), RequestIDKey, id)
				req = r.WithContext(ctx)
			}

			//Call initialize handler
			h.ServeHTTP(w, req)
		})
	}
}
