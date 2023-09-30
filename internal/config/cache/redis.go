package cache

import (
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

type ServiceRegistryInstance struct {
	IpAddress string `json:"ipAddress"`
	Port      int    `json:"port"`
}
