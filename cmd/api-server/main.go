package main

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httputil"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	router := createRouter()
	log.Println("starting server")
	err := router.Run(":8000")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func createRouter() *gin.Engine {

	serviceMap := getServices()
	log.Println("service info readed")
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	for _, service := range serviceMap {
		serviceName := "/" + service.Name
		serviceUrl := service.Url
		router.GET(serviceName, createReverseProxy(serviceUrl, serviceName))
		router.POST(serviceName, createReverseProxy(serviceUrl, serviceName))
	}
	return router
}

func createReverseProxy(target string, path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the target URL
		targetURL, _ := url.Parse(target)

		// Create the reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// Modify the request
		c.Request.URL.Scheme = targetURL.Scheme
		c.Request.URL.Host = targetURL.Host
		c.Request.URL.Path = path

		// Let the reverse proxy do its job
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func getServices() map[string]model.ServiceConfig {

	yamlFile, err := os.ReadFile("resources/config.yaml")

	if err != nil {
		log.Fatalln(err.Error())
	}

	data := make(map[string]model.ServiceConfig)

	err2 := yaml.Unmarshal(yamlFile, &data)

	if err2 != nil {
		log.Fatalln(err2.Error())
	}
	return data
}
