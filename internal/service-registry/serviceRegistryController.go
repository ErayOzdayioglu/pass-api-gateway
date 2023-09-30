package service_registry

import (
	"encoding/json"
	"github.com/ErayOzdayioglu/api-gateway/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

func GetServiceRegistry(client *redis.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceName := context.Param("name")

		val, err := client.Get(context, serviceName).Result()

		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		jsonData, err := json.Marshal(val)

		context.JSON(http.StatusOK, &jsonData)
	}
}

func PostServiceRegistry(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := model.AddToServiceRegistryRequest{}

		if err := c.BindJSON(&body); err != nil {
			log.Fatalln(err.Error())
		}

		jsonData, err := json.Marshal(model.AddToServiceRegistryRequest{
			IpAddress: body.IpAddress,
			Port:      body.Port})

		if err != nil {
			log.Fatalln(err.Error())
		}

		err = client.Set(c, body.ServiceName, jsonData, 0).Err()

		if err != nil {
			log.Fatalln(err.Error())
		}

		c.String(http.StatusCreated, "Created")
	}
}
