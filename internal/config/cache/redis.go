package cache

import (
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

type ServiceEntity struct {
	ServiceName string            `json:"serviceName"`
	IpAddresses []IpAddressEntity `json:"ipAddresses"`
}

type IpAddressEntity struct {
	IpAddress   string `json:"addr"`
	IsAvailable bool   `json:"isAvailable"`
}
