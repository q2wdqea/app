package main

import "github.com/gin-gonic/gin"

func main() {
	ginServer := gin.Default()
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "hello world"})
	})
	//服务器端口
	ginServer.Run(":9090") /*默认是8080*/
}
