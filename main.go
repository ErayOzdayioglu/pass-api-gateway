package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/sa", selam)
	err := r.Run()
	if err != nil {
		return
	}
}

func selam(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "as",
	})

}
