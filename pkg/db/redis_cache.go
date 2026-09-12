package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache wraps the Redis client and provides simple cache operations.
// It is injected into usecases that need caching.
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache creates a new RedisCache wrapper.
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

// Set serializes the value to JSON and stores it in Redis with a TTL.
// Example: cache.Set(ctx, "products:page:1:limit:10", productList, 10*time.Minute)
func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("redis cache marshal error: %w", err)
	}
	return r.client.Set(ctx, key, data, ttl).Err()
}

// Get retrieves a cached value by key and unmarshals it into dest.
// Returns true if found, false if cache miss (key doesn't exist).
// Example:
//
//	var products []models.Inventories
//	found, err := cache.Get(ctx, "products:page:1:limit:10", &products)
func (r *RedisCache) Get(ctx context.Context, key string, dest interface{}) (bool, error) {
	data, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		// Cache MISS — key doesn't exist
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("redis cache get error: %w", err)
	}

	// Cache HIT — unmarshal the JSON into dest
	if err := json.Unmarshal(data, dest); err != nil {
		return false, fmt.Errorf("redis cache unmarshal error: %w", err)
	}
	return true, nil
}

// DeleteByPattern deletes all keys matching a glob pattern.
// Example: cache.DeleteByPattern(ctx, "products:*")
// This is used for cache invalidation when products are added/deleted.
func (r *RedisCache) DeleteByPattern(ctx context.Context, pattern string) error {
	var cursor uint64
	for {
		// SCAN is non-blocking (unlike KEYS) — safe for production
		keys, nextCursor, err := r.client.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return fmt.Errorf("redis scan error: %w", err)
		}

		if len(keys) > 0 {
			if err := r.client.Del(ctx, keys...).Err(); err != nil {
				return fmt.Errorf("redis delete error: %w", err)
			}
		}

		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return nil
}
