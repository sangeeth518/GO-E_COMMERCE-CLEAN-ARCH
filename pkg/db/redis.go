package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
)

// ConnectRedis creates a new Redis client and verifies the connection with a PING.
// This is called once at startup (just like ConnectDB for PostgreSQL).
func ConnectRedis(cfg config.Config) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,             // e.g. "localhost:6379"
		Password: "",               // no password for local dev
		DB:       0,                // default DB
		PoolSize: 10,               // max connections in pool
		MinIdleConns: 5,            // keep 5 idle connections ready
	})

	// Verify connection with a 5-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis at %s: %w", addr, err)
	}

	fmt.Printf("✅ Redis connected successfully at %s\n", addr)
	return client, nil
}
