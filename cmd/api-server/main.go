package main

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/config/cache"
	"github.com/ErayOzdayioglu/api-gateway/internal/proxy"
	serviceregistry "github.com/ErayOzdayioglu/api-gateway/internal/service-registry"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	redisClient := cache.RedisClient()
	router := createRouters(redisClient)

	err := router.Run(":8000")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func createRouters(redisClient *redis.Client) *gin.Engine {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router = addServiceRegistryEndpoints(router, redisClient)
	router = reverseProxy(router, redisClient)
	return router
}

func addServiceRegistryEndpoints(router *gin.Engine, redisClient *redis.Client) *gin.Engine {
	router.POST("/service-registry", serviceregistry.PostServiceRegistry(redisClient))
	router.GET("/service-registry/:name", serviceregistry.GetServiceRegistry(redisClient))
	return router
}

func reverseProxy(router *gin.Engine, redisClient *redis.Client) *gin.Engine {
	router.GET("/api/:name/*path", proxy.CreateReverseProxy(redisClient))
	router.POST("/api/:name/*path", proxy.CreateReverseProxy(redisClient))
	router.DELETE("/api/:name/*path", proxy.CreateReverseProxy(redisClient))
	router.PUT("/api/:name/*path", proxy.CreateReverseProxy(redisClient))
	return router
}
