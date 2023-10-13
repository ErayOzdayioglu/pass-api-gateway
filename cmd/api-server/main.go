package main

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/config/cache"
	"github.com/ErayOzdayioglu/api-gateway/internal/config/database"
	"github.com/ErayOzdayioglu/api-gateway/internal/proxy"
	service_registry "github.com/ErayOzdayioglu/api-gateway/internal/service-registry"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	redisClient := cache.ConnectRedisClient()
	cb := database.InitCouchbaseBucket()
	router := createRouters(redisClient, cb)

	err := router.Run(":8000")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func createRouters(redisClient *redis.Client, cb *gocb.Bucket) *gin.Engine {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	proxyController := proxy.ReverseProxyController{
		Redis:     redisClient,
		Couchbase: cb,
	}
	serviceRegistryController := service_registry.ServiceRegistryController{
		Redis:     redisClient,
		Couchbase: cb,
	}
	proxyController.RegisterRoutes(router)
	serviceRegistryController.RegisterRoutes(router)
	return router
}
