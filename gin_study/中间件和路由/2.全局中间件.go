package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

func m10(c *gin.Context) {
	fmt.Println("m1 ...in")
	c.Set("user", User{"123", 21})
	c.Next()
	fmt.Println("m1 ...out")
}

func main() {
	router := gin.Default()

	router.Use(m10)
	router.GET("/", func(c *gin.Context) {
		fmt.Println("index ...in")
		_user, _ := c.Get("user")
		user := _user.(User) //将interface{}类型的断言为User类型
		fmt.Println(user.Name, user.Age)
		c.JSON(200, gin.H{"msg": "index"})
	})

	router.Run(":8080")

}
