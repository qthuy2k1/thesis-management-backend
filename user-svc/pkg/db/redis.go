package db

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisInitialize(port, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("thesis-management-backend-user-redis-db-service:%s", port),
		Password: password,
		DB:       0,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Could not connect to redis server: %v", err)
	}

	return client
}
