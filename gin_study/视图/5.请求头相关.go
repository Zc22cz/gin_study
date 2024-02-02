package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()

	//请求头的各种获取方式
	router.GET("/", func(c *gin.Context) {
		//首字母不分大小写，单词与单词用-连接
		//用于获取一个请求头
		fmt.Println(c.GetHeader("User-Agent"))
		//fmt.Println(c.GetHeader("user-Agent"))
		//fmt.Println(c.GetHeader("user-agent"))
		//GetHeader和Get不用区分大小写，并且返回第一个value
		fmt.Println(c.Request.Header.Get("User-Agent"))
		//用map的取值方式就要注意大小写
		fmt.Println(c.Request.Header["User-Agent"])
		fmt.Println(c.Request.Header["User-agent"])
		c.JSON(200, gin.H{"msg": "成功"})
	})

	//爬虫和用户区别对待
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		if strings.Contains(userAgent, "python") {
			c.JSON(0, gin.H{"data": "这是响应给爬虫的数据"})
			return
		}
		c.JSON(0, gin.H{"data": "这是响应给用户的数据"})
	})

	//设置响应头
	router.GET("/res", func(c *gin.Context) {
		c.Header("Token", "231dqewqweqweqw")
		c.Header("Content-Type", "application/text; charset=utf-8")
		c.JSON(0, gin.H{"data": "看看响应头"})
	})

	router.Run(":80")
}
