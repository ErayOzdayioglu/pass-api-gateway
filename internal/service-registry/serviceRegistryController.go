package service_registry

import (
	"encoding/json"
	"fmt"
	"github.com/ErayOzdayioglu/api-gateway/internal/config/cache"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetServiceRegistry() gin.HandlerFunc {
	return func(context *gin.Context) {
		client := cache.Client
		serviceName := context.Param("name")

		val, err := client.Get(context, serviceName).Result()

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"message": "There is no service in registry : " + serviceName,
			})
			return
		}

		var serviceEntity cache.ServiceEntity

		err = json.Unmarshal([]byte(val), &serviceEntity)

		if err != nil {
			return
		}

		context.JSON(http.StatusOK, &serviceEntity)
	}
}

func PostServiceRegistry() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := cache.Client
		body := cache.ServiceEntity{}

		if err := c.BindJSON(&body); err != nil {
			log.Println(err.Error())
		}

		jsonData, err := json.Marshal(body)

		err = client.Set(c, body.ServiceName, jsonData, 0).Err()

		if err != nil {
			fmt.Println(body)
			log.Println(err.Error())
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Created",
		})
	}
}
