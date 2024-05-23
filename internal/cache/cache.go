package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type CacheClient interface {
	Get(key string) ([]map[string]string, bool)
	Set(key string, value interface{}, ttl time.Duration)
}

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(addr string) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis")
		panic(err)
	}
	return &RedisClient{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (rc *RedisClient) Get(key string) ([]map[string]string, bool) {
	val, err := rc.client.Get(rc.ctx, key).Result()
	if err == redis.Nil || err != nil {
		return nil, false
	}

	var result []map[string]string
	json.Unmarshal([]byte(val), &result)
	return result, true
}

func (rc *RedisClient) Set(key string, value interface{}, ttl time.Duration) {
	val, _ := json.Marshal(value)
	rc.client.Set(rc.ctx, key, val, ttl)
}
