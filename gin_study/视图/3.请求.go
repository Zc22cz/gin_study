package main

import (
	"encoding/json"
	"fmt"
	main2 "gin_study"
	"github.com/gin-gonic/gin"
)

func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user"))
	fmt.Println(c.QueryMap("user"))
	fmt.Println(c.DefaultQuery("addr", "湖北省"))
}

func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "湖北省")) //如果用户没有传，就使用默认值
	forms, err := c.MultipartForm()               //接受所有form参数，包括文件
	fmt.Println(forms, err)
}

func _raw(c *gin.Context) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-type")
	switch contentType {
	case "application/json":
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		err := json.Unmarshal(body, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(user)
}

func main2.main() {
	router := gin.Default()
	router.GET("/query", _query)
	router.GET("/param/:user_id/", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)
	router.Run(":81")
}
