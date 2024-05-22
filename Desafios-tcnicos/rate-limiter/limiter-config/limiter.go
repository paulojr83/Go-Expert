package limiter_config

import (
	"context"
	"fmt"
	"github.com/paulojr83/Go-Expert/Desafios-tcnicos/rate-limiter/config"
	"github.com/paulojr83/Go-Expert/Desafios-tcnicos/rate-limiter/storage"
	"time"
)

type Limiter struct {
	storage storage.Storage
	config  *config.Config
}

func NewLimiter(storage storage.Storage, config *config.Config) *Limiter {
	return &Limiter{storage: storage, config: config}
}

func (l *Limiter) AllowRequest(ctx context.Context, key string, limit int) (bool, error) {
	count, err := l.storage.Incr(ctx, key)
	if err != nil {
		return false, err
	}

	expireTime := time.Duration(l.config.BlockTime) * time.Second

	if count == 1 {
		l.storage.Expire(ctx, key, expireTime)
	}

	if count > int64(limit) {
		return false, nil
	}

	return true, nil
}

func (l *Limiter) GetLimitKey(ip, token string) (string, int) {
	if token != "" {
		return fmt.Sprintf("token:%s", token), l.config.RateLimitToken
	}
	return fmt.Sprintf("ip:%s", ip), l.config.RateLimitIP
}
