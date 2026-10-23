package srv

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	const EnvRedisAddr = "REDIS_ADDR"
	const EnvDefaultRedisAddress = "localhost:6379"
	redisAddr := os.Getenv(EnvRedisAddr)
	if redisAddr == "" {
		log.Printf(
			"Redist address %s not set, default: %s",
			EnvRedisAddr, EnvDefaultRedisAddress,
		)
		redisAddr = EnvDefaultRedisAddress
	}

	return redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
}
