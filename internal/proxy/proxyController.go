package proxy

import (
	"encoding/json"
	"github.com/ErayOzdayioglu/api-gateway/internal/config/cache"
	"github.com/ErayOzdayioglu/api-gateway/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func CreateReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		redisClient := cache.Client
		serviceName := c.Param("name")
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
		log.Printf("Proxy at : %s", targetURL.String())
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		log.Printf("Forwarding request to %s\n", c.Request.URL.String())
		// Let the reverse proxy do its job
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
