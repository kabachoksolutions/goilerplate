package RedisCache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(connUrl string, db int) (*RedisCache, error) {
	opts, err := redis.ParseURL(connUrl)
	if err != nil {
		return nil, fmt.Errorf("RedisCache: failed to parse URL: %w", err)
	}
	opts.DB = db

	rdb := redis.NewClient(opts)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("RedisCache: failed to ping redis: %w", err)
	}

	return &RedisCache{
		client: rdb,
	}, nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if err := c.client.Set(ctx, key, value, ttl).Err(); err != nil {
		return fmt.Errorf("RedisCache: failed to set key: %w", err)
	}

	return nil
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("RedisCache: failed to get key: %w", err)
	}

	return val, nil
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	if err := c.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("RedisCache: failed to delete key: %w", err)
	}

	return nil
}

func (c *RedisCache) CountKeys(ctx context.Context) (int, error) {
	keys, err := c.client.Keys(ctx, "*").Result()
	if err != nil {
		return 0, fmt.Errorf("RedisCache: failed to count keys: %w", err)
	}

	return len(keys), nil
}
