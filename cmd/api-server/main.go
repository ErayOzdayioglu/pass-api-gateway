package main

import (
	"github.com/ErayOzdayioglu/api-gateway/internal"
	"github.com/ErayOzdayioglu/api-gateway/internal/config/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// redisClient := cache.ConnectRedisClient()
	db := database.ConnectToMongo()
	router := gin.Default()
	internal.ApplicationRouter(router, db)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
