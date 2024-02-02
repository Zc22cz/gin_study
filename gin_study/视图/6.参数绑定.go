package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindJSON(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)
	})

	router.POST("/query", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindQuery(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)
	})

	router.POST("/uri/:name/:age/:sex", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindUri(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)
	})

	router.POST("/form", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBind(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)
	})

	router.Run(":80")
}
