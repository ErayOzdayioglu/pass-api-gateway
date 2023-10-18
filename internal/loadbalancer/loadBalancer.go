package loadbalancer

import (
	"context"
	"encoding/json"
	"github.com/ErayOzdayioglu/api-gateway/internal/service"
	"github.com/redis/go-redis/v9"
	"time"
)

type LoadBalancerI struct {
	cache *redis.Client
}

type LoadBalancer interface {
	SelectTheRouteWithRoundRobin(service *service.Service) (*service.Host, error)
}

func (lb *LoadBalancerI) SelectTheRouteWithRoundRobin(service *service.Service) (*service.Host, error) {
	cacheEntity, err := lb.findEntityFromCache(service.ServiceName)

	if err != nil {
		lb.saveToCache(service)
		return &service.Hosts[0], nil
	}
	oldestRequestedHost := findOldest(cacheEntity.Hosts)
	availableHosts := service.Hosts
	return findCorrectHost(oldestRequestedHost, availableHosts), nil
}

func findCorrectHost(host *RoundRobinEntity, hosts []service.Host) *service.Host {
	for _, h := range hosts {
		if h.Url == host.Url {
			return &h
		}
	}
	return &hosts[0]
}

func findOldest(hosts []RoundRobinEntity) *RoundRobinEntity {
	host := &RoundRobinEntity{
		Url:           "",
		LastRequestAt: time.Now(),
	}

	for _, h := range hosts {
		if h.LastRequestAt.Before(host.LastRequestAt) {
			host = &h
		}
	}
	return host
}

func (lb *LoadBalancerI) findEntityFromCache(name string) (*RRCacheEntity, error) {
	val, err := lb.cache.Get(context.Background(), name).Result()
	if err != nil {
		return nil, err
	}

	var cacheEntity RRCacheEntity
	err = json.Unmarshal([]byte(val), &cacheEntity)

	if err != nil {
		return nil, err
	}
	return &cacheEntity, nil
}

func (lb *LoadBalancerI) saveToCache(s *service.Service) {
	var hosts []RoundRobinEntity

	for _, h := range s.Hosts {
		hosts = append(hosts, RoundRobinEntity{Url: h.Url, LastRequestAt: time.Now()})
	}
	cacheEntity := &RRCacheEntity{
		Hosts: hosts,
	}
	lb.cache.Set(context.Background(), s.ServiceName, cacheEntity, time.Second*3600)
}

func NewLoadBalancer(client *redis.Client) LoadBalancer {
	return &LoadBalancerI{
		cache: client,
	}
}
