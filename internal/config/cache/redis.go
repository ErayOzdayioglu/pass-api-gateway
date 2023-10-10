package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var Client *redis.Client

func ConnectRedisClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if status := Client.Ping(context.Background()); status.Err() != nil {
		log.Fatal("Failed to connect redis.")
	}
	log.Println("Connected to redis.")

}

type ServiceEntity struct {
	ServiceName string            `json:"serviceName"`
	IpAddresses []IpAddressEntity `json:"ipAddresses"`
}

type IpAddressEntity struct {
	IpAddress   string `json:"addr"`
	IsAvailable bool   `json:"isAvailable"`
}
