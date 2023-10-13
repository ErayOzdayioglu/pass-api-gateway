package service_registry

import (
	"encoding/json"
	"github.com/ErayOzdayioglu/api-gateway/internal/config/cache"
	"github.com/ErayOzdayioglu/api-gateway/internal/model"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

type ServiceRegistryController struct {
	Redis     *redis.Client
	Couchbase *gocb.Bucket
}

func (c *ServiceRegistryController) RegisterRoutes(router *gin.Engine) {
	router.POST("/service-registry", c.PostServiceRegistry())
	router.GET("/service-registry/:name", c.GetServiceRegistry())
}

func (c *ServiceRegistryController) GetServiceRegistry() gin.HandlerFunc {
	return func(context *gin.Context) {
		redisClient := c.Redis
		couchbase := c.Couchbase

		serviceName := context.Param("name")

		val, err := redisClient.Get(context, serviceName).Result()

		var serviceEntity cache.ServiceEntity

		err = json.Unmarshal([]byte(val), &serviceEntity)

		if err == nil {
			context.JSON(http.StatusOK, &serviceEntity)
			return
		}
		col := couchbase.Scope("registry").Collection("services")

		getResult, err := col.Get(serviceName, nil)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"message": "There is no service in registry : " + serviceName,
			})
			return
		}
		err = getResult.Content(&serviceEntity)

		if err != nil {
			log.Println(err.Error())
		}

		context.JSON(http.StatusOK, &serviceEntity)
	}
}

func (c *ServiceRegistryController) PostServiceRegistry() gin.HandlerFunc {
	return func(context *gin.Context) {
		redisClient := c.Redis
		couchbase := c.Couchbase
		col := couchbase.Scope("registry").Collection("services")

		body := model.ServiceEntityRequest{}

		if err := context.BindJSON(&body); err != nil {
			log.Println(err.Error())
		}
		setDefaults(&body)
		jsonData, err := json.Marshal(body)

		err = redisClient.Set(context, body.ServiceName, jsonData, 0).Err()

		if err != nil {
			log.Println(err.Error())
		}

		_, err = col.Upsert(body.ServiceName, body, nil)

		if err != nil {
			log.Println(err.Error())
		}

		context.JSON(http.StatusCreated, gin.H{
			"message": "Created",
		})
	}

}

func setDefaults(m *model.ServiceEntityRequest) {
	for _, item := range m.IpAddresses {
		item.IsAvailable = false
	}
}
