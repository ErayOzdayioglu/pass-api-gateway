package internal

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/proxy"
	"github.com/ErayOzdayioglu/api-gateway/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func ApplicationRouter(engine *gin.Engine, db *mongo.Database, cache *redis.Client) {
	proxy.Router(engine.Group("/"), db, cache)
	service.Router(engine.Group("/service"), db)
}
