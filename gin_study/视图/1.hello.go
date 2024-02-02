package main

import "github.com/gin-gonic/gin"

func Index(context *gin.Context) {
	context.String(200, "Hello")
}

func main() {
	route := gin.Default()

	route.GET("/index", Index)

	route.Run(":8080")
}
