package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var client *redis.Client

func ConnectRedisClient() *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if status := client.Ping(context.Background()); status.Err() != nil {
		log.Fatal("Failed to connect redis.")
	}
	log.Println("Connected to redis.")

	return client
}
