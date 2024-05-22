package server

import (
	"context"
	"fmt"
	"github.com/paulojr83/Go-Expert/Desafios-tcnicos/rate-limiter/limiter-config"
	"net/http"
)

func RateLimitMiddleware(limiter *limiter_config.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := ReadUserIP(r)
			token := r.Header.Get("API_KEY")
			/*if token == "" {
				http.Error(w, "there is no token", http.StatusUnauthorized)
				return
			}*/
			key, limit := limiter.GetLimitKey(ip, token)
			allowed, err := limiter.AllowRequest(context.Background(), key, limit)

			if err != nil {
				http.Error(w, fmt.Sprintf("Internal Server Error %s", err), http.StatusInternalServerError)
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

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
