package database

import (
	"context"
	"net/http"

	"github.com/ShardulNalegave/PICT-EMS/utils"
	"gorm.io/gorm"
)

func DatabaseMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), utils.DatabaseKey, db)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
