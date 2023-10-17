package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(group *gin.RouterGroup, db *mongo.Database) {
	serviceRepository := NewServiceRepository(db)
	serviceController := &ServiceController{
		serviceRepository,
	}

	group.POST("/", serviceController.CreateService)
	group.GET("/:name", serviceController.FindByName)
}
