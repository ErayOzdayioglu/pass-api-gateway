package proxy

import (
	"encoding/json"
	"github.com/ErayOzdayioglu/api-gateway/internal/model"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

type ReverseProxyController struct {
	Redis     *redis.Client
	Couchbase *gocb.Bucket
}

func (c *ReverseProxyController) RegisterRoutes(router *gin.Engine) {
	router.GET("/api/:name/*path", c.CreateReverseProxy())
	router.POST("/api/:name/*path", c.CreateReverseProxy())
	router.DELETE("/api/:name/*path", c.CreateReverseProxy())
	router.PUT("/api/:name/*path", c.CreateReverseProxy())
}

func (c *ReverseProxyController) CreateReverseProxy() gin.HandlerFunc {
	return func(context *gin.Context) {
		redisClient := c.Redis
		serviceName := context.Param("name")
		val, err := redisClient.Get(context, serviceName).Result()

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{
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

		log.Printf("Forwarding request to %s\n", context.Request.URL.String())
		// Let the reverse proxy do its job
		proxy.ServeHTTP(context.Writer, context.Request)
	}
}
