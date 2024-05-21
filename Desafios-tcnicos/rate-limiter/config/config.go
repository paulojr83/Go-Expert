package config

import (
	"os"
	"strconv"
)

type Config struct {
	StorageType    string
	RateLimitIP    int
	RateLimitToken int
	BlockTime      int
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
}

func LoadConfig() (*Config, error) {
	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "" {
		storageType = "memory"
	}

	rateLimitIP, err := strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	if err != nil {
		return nil, err
	}

	rateLimitToken, err := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	if err != nil {
		return nil, err
	}

	blockTime, err := strconv.Atoi(os.Getenv("BLOCK_TIME"))
	if err != nil {
		return nil, err
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}

	return &Config{
		StorageType:    storageType,
		RateLimitIP:    rateLimitIP,
		RateLimitToken: rateLimitToken,
		BlockTime:      blockTime,
		RedisAddr:      os.Getenv("REDIS_ADDR"),
		RedisPassword:  os.Getenv("REDIS_PASSWORD"),
		RedisDB:        redisDB,
	}, nil
}
