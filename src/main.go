package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {
		log.Println("start gin hello")
		context.JSON(200,gin.H{
			"code":200,
			"success": true,
		})
	})
	// 指定地址和端口号
	router.Run("localhost:9090")
}
