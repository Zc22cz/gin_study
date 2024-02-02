package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	fmt.Println("m1...in")
	c.Next()
	fmt.Println("m1...out")
}

func index(c *gin.Context) {
	fmt.Println("index...in")
	c.JSON(200, gin.H{"msg": "index"})
	c.Next()
	fmt.Println("index...out")
}

func m2(c *gin.Context) {
	fmt.Println("m2...in")
	c.Next()
	fmt.Println("m2...out")
}

func main() {
	router := gin.Default()
	router.GET("/", m1, index, m2)
	router.Run(":80")
}
