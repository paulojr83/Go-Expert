package main

import (
	"github.com/joho/godotenv"
	"github.com/paulojr83/Go-Expert/Desafios-tcnicos/rate-limiter/config"
	"github.com/paulojr83/Go-Expert/Desafios-tcnicos/rate-limiter/limiter-config"
	"github.com/paulojr83/Go-Expert/Desafios-tcnicos/rate-limiter/server"
	storage "github.com/paulojr83/Go-Expert/Desafios-tcnicos/rate-limiter/storage"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dataConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	var storageType storage.Storage

	switch dataConfig.StorageType {
	case "redis":
		storageType = storage.NewRedisClient(dataConfig)
	case "memory":
		storageType = storage.NewMemoryStorage()
	default:
		log.Fatalf("Unknown storage type")
	}

	limiter := limiter_config.NewLimiter(storageType, dataConfig)

	http.Handle("/", server.RateLimitMiddleware(limiter)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
