package utils

import (
	"context"
	"net/http"
)

func LocationMiddleware(loc string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), LocationKey, loc)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
