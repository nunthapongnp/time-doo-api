package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func InitRedis(env Env) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.RedisAddress,
		Password: env.RedisPassword,
		DB:       0,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
	return client
}
