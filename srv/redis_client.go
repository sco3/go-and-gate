package srv

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	const REDIS_ADDR = "REDIS_ADDR"
	const DEFAULT_REDIS_ADDRESS = "localhost:6379"
	redisAddr := os.Getenv(REDIS_ADDR)
	if redisAddr == "" {
		log.Printf(
			"Redist address %s not set, default: %s",
			REDIS_ADDR, DEFAULT_REDIS_ADDRESS,
		)
		redisAddr = DEFAULT_REDIS_ADDRESS
	}

	return redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
}
