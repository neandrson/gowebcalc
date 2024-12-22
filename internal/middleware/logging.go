package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

func Logging(next http.HandlerFunc, log *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.With(zap.String("method", r.Method), zap.String("path", r.URL.Path)).Info("incoming request")
		next.ServeHTTP(w, r)
	}
}
