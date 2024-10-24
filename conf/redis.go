package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var (
	redisClient *redis.Client
)
var expire = 3600 * 24 * 30 * time.Second

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host"),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, value any) error {
	return redisClient.Set(context.Background(), key, value, expire).Err()
}

func (rc *RedisClient) Get(key string) (string, error) {
	return redisClient.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Delete(keys ...string) error {
	return redisClient.Del(context.Background(), keys...).Err()
}
