package srv

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const barrierLuaScript = `
redis.call('SADD', KEYS[1], ARGV[1])
redis.call('EXPIRE', KEYS[1], tonumber(ARGV[3]))
local count = redis.call('SCARD', KEYS[1])
if count >= tonumber(ARGV[2]) then
    return 1
else
    return 0
end
`

func WaitStartupBarrier(ctx context.Context, rdb *redis.Client) error {
	podName, err := os.Hostname()
	if err != nil {
		podName = "unknown-pod"
	}

	expectedReplicas := 0
	if val := os.Getenv("EXPECTED_REPLICAS"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			expectedReplicas = parsed
		}
	}

	if expectedReplicas <= 0 {
		log.Printf("[%s] Startup barrier is disabled (EXPECTED_REPLICAS not set). Skipping.", podName)
		return nil
	}

	barrierKey := "server:startup:barrier"
	ttlSec := 30

	log.Printf("[%s] Joining Redis startup barrier (expected: %d)...", podName, expectedReplicas)

	script := redis.NewScript(barrierLuaScript)

	for {
		res, err := script.Run(ctx, rdb, []string{barrierKey}, podName, expectedReplicas, ttlSec).Int()
		if err != nil {
			log.Printf("[%s] Error running barrier script: %v, retrying...", podName, err)
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if res == 1 {
			log.Printf("[%s] Barrier reached! All %d replicas are synchronized.", podName, expectedReplicas)
			break
		}

		log.Printf("[%s] Waiting for other replicas...", podName)
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}
