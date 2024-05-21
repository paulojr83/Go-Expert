package limiter_config

import (
	"context"
	"github.com/paulojr83/Go-Expert/Desafios-tcnicos/rate-limiter/config"
	"github.com/paulojr83/Go-Expert/Desafios-tcnicos/rate-limiter/storage"
	"testing"
)

func TestRateLimiterWithRedis(t *testing.T) {
	dataConfig := &config.Config{
		RateLimitIP:    5,
		RateLimitToken: 10,
		BlockTime:      300,
		RedisAddr:      "localhost:6379",
		RedisPassword:  "",
		RedisDB:        0,
	}

	redisClient := storage.NewRedisClient(dataConfig)
	testLimiter(t, redisClient, dataConfig)
}

func TestRateLimiterWithMemory(t *testing.T) {
	dataConfig := &config.Config{
		RateLimitIP:    5,
		RateLimitToken: 10,
		BlockTime:      300,
	}

	memoryStorage := storage.NewMemoryStorage()
	testLimiter(t, memoryStorage, dataConfig)
}

func testLimiter(t *testing.T, storage storage.Storage, config *config.Config) {
	limiter := NewLimiter(storage, config)

	ctx := context.Background()
	key := "ip:test"

	// Reset key
	storage.Set(ctx, key, 0, 0)

	for i := 0; i < config.RateLimitIP; i++ {
		allowed, _ := limiter.AllowRequest(ctx, key, config.RateLimitIP)
		if !allowed {
			t.Errorf("request %d should have been allowed", i+1)
		}
	}

	allowed, _ := limiter.AllowRequest(ctx, key, config.RateLimitIP)
	if allowed {
		t.Errorf("request should have been blocked")
	}
}
