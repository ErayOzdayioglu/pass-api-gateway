package internal

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/proxy"
	"github.com/ErayOzdayioglu/api-gateway/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ApplicationRouter(engine *gin.Engine, db *mongo.Database) {
	proxy.Router(engine.Group("/"), db)
	service.Router(engine.Group("/service"), db)
}
