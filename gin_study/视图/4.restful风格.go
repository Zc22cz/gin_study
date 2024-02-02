package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func _getList(c *gin.Context) {
	articleList := []ArticleModel{
		{"GO", "文章是go语言"},
		{"JAVA", "文章是java语言"},
		{"PYTHON", "文章是python语言"},
	}
	c.JSON(200, Response{0, articleList, "成功"})
}

func _getDetail(c *gin.Context) {
	fmt.Println(c.Param("id"))
	article := ArticleModel{
		"GO", "文章是go语言",
	}
	c.JSON(200, Response{0, article, "成功"})
}

func _create(c *gin.Context) {

	var article ArticleModel

	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, Response{0, article, "添加成功"})
}

func _update(c *gin.Context) {
	fmt.Println(c.Param("id"))
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{0, article, "修改成功"})
}

func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(200, Response{0, map[string]string{}, "删除成功"})
}

func main() {
	router := gin.Default()
	router.GET("/articles", _getList)
	router.GET("/articles/:id", _getDetail)
	router.POST("/articles", _create)
	router.PUT("/articles/:id", _update)
	router.DELETE("/articles/:id", _delete)
	router.Run(":80")
}
