package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/download", func(c *gin.Context) {
		c.Header("Content-type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+"六花.png")
		c.File("uploads/1.png")
	})
	router.Run(":8080")
}
