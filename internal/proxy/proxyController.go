package proxy

import (
	"encoding/json"
	"github.com/ErayOzdayioglu/api-gateway/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func CreateReverseProxy(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		//targetURL, _ := url.Parse(target)
		//log.Printf("Target Url is : %s\n", targetURL.String())
		// Create the reverse proxy
		serviceName := c.Param("name")
		//path := c.Param("path")
		val, err := redisClient.Get(c, serviceName).Result()

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "There is no service in registry : " + serviceName,
			})
			return
		}
		var serviceRegistryInstance model.GetServiceRegistryResponse

		err = json.Unmarshal([]byte(val), &serviceRegistryInstance)
		if err != nil {
			return
		}

		urlString := serviceRegistryInstance.IpAddress + ":" + strconv.Itoa(serviceRegistryInstance.Port)
		targetURL, _ := url.Parse(urlString)
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		//c.Request.URL.Scheme = targetURL.Scheme
		//c.Request.URL.Host = targetURL.Host
		// c.Request.URL.Path = path
		//c.Request.URL = targetURL
		log.Printf("Forwarding request to %s\n", c.Request.URL.String())
		// Let the reverse proxy do its job
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
