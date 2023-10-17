package proxy

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(group *gin.RouterGroup, db *mongo.Database) {
	serviceRepository := service.NewServiceRepository(db)
	proxyController := ReverseProxyController{
		ServiceRepository: serviceRepository,
	}

	group.GET("/api/:name/*path", proxyController.CreateReverseProxy)
	group.POST("/api/:name/*path", proxyController.CreateReverseProxy)
	group.DELETE("/api/:name/*path", proxyController.CreateReverseProxy)
	group.PUT("/api/:name/*path", proxyController.CreateReverseProxy)
}
