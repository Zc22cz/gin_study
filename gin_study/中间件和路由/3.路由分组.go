package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	r := router.Group("/api")
	r.GET("/index", func(c *gin.Context) {
		c.String(200, "index")
	})
	r.GET("/home", func(c *gin.Context) {
		c.String(200, "home")
	})

	router.Run(":80")
}
