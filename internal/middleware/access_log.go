package middleware

import (
	"go.uber.org/zap"
	"net/http"
)

func NewAccessLog(log *zap.Logger) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info("Request",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()))

			handler.ServeHTTP(w, r)
		})
	}
}
