package cache

import (
	"context"
	"errors"
	"time"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

var ErrCacheMiss = errors.New("cache miss")

type Cacher interface {
	Get(ctx context.Context, key string, dest interface{}) error
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisCache{client: rdb}
}

func (r *RedisCache) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return ErrCacheMiss
	} else if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, bytes, ttl).Err()
}
