package server

import (
	"context"
	"github.com/paulojr83/Go-Expert/Desafios-tcnicos/rate-limiter/limiter-config"
	"net/http"
	"strings"
)

func RateLimitMiddleware(limiter *limiter_config.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := strings.Split(r.RemoteAddr, ":")[0]
			token := r.Header.Get("Authorization")

			if token == "" {
				http.Error(w, "there is no token", http.StatusUnauthorized)
				return
			}
			key, limit := limiter.GetLimitKey(ip, token)
			allowed, err := limiter.AllowRequest(context.Background(), key, limit)

			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if !allowed {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("you have reached the maximum number of requests or actions allowed within a certain time frame"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
