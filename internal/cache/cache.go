package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheClient struct {
	client *redis.Client
}

func NewCacheClient(addr string) CacheClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	if rdb == nil {
		panic("could not create cache client")
	}
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic("could not connect to cache")
	}
	return CacheClient{client: rdb}
}

func (r *CacheClient) Get(key string) ([]map[string]string, bool) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, false
	}

	var result []map[string]string
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, false
	}

	return result, true
}

func (r *CacheClient) Set(key string, value interface{}, ttl time.Duration) bool {
	ctx := context.Background()
	jsonData, err := json.Marshal(value)
	if err != nil {
		return false
	}

	err = r.client.Set(ctx, key, jsonData, ttl).Err()
	return err == nil
}
