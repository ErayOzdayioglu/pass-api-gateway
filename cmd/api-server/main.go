package main

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/config/cache"
	"github.com/ErayOzdayioglu/api-gateway/internal/config/database"
	"github.com/ErayOzdayioglu/api-gateway/internal/proxy"
	serviceregistry "github.com/ErayOzdayioglu/api-gateway/internal/service-registry"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cache.ConnectRedisClient()
	database.ConnectDB()
	router := createRouters()

	err := router.Run(":8000")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func createRouters() *gin.Engine {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router = addServiceRegistryEndpoints(router)
	router = reverseProxy(router)
	return router
}

func addServiceRegistryEndpoints(router *gin.Engine) *gin.Engine {
	router.POST("/service-registry", serviceregistry.PostServiceRegistry())
	router.GET("/service-registry/:name", serviceregistry.GetServiceRegistry())
	return router
}

func reverseProxy(router *gin.Engine) *gin.Engine {
	router.GET("/api/:name/*path", proxy.CreateReverseProxy())
	router.POST("/api/:name/*path", proxy.CreateReverseProxy())
	router.DELETE("/api/:name/*path", proxy.CreateReverseProxy())
	router.PUT("/api/:name/*path", proxy.CreateReverseProxy())
	return router
}
