package proxy

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/loadbalancer"
	"github.com/ErayOzdayioglu/api-gateway/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(group *gin.RouterGroup, db *mongo.Database, cache *redis.Client) {
	serviceRepository := service.NewServiceRepository(db)
	loadBalancer := loadbalancer.NewLoadBalancer(cache)
	proxyController := ReverseProxyController{
		ServiceRepository: serviceRepository,
		LoadBalancer:      loadBalancer,
	}

	group.GET("/:name/*path", proxyController.CreateReverseProxy)
	group.POST("/:name/*path", proxyController.CreateReverseProxy)
	group.DELETE("/:name/*path", proxyController.CreateReverseProxy)
	group.PUT("/:name/*path", proxyController.CreateReverseProxy)
}
