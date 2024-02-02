package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 函数名前+下划线表示该函数只能在本包内使用，不对外公开
func _string(c *gin.Context) {
	c.String(200, "zzz")
}

func _json(c *gin.Context) {
	type UserInfo struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		Password string `json:"-"`
	}
	user := UserInfo{"杰", 20, "123456"}
	c.JSON(200, user)
}

func _xml(c *gin.Context) {
	c.XML(200, gin.H{"user": "hanru", "message": "hey", "status": http.StatusOK, "data": gin.H{"user": "zzz"}})
}

func _yaml(c *gin.Context) {
	c.YAML(200, gin.H{"user": "hanru", "message": "hey", "status": http.StatusOK, "data": gin.H{"user": "zzz"}})
}

func _html(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"username": "杰"})
}

func _redirect(c *gin.Context) {
	c.Redirect(301, "http://www.baidu.com") //301永久，302临时
}

func main() {
	router := gin.Default()

	//加载
	router.LoadHTMLGlob("templates/*")
	//在golang中，没有相对文件路径，只有相对项目路径
	//配置单个文件，第一个参数是网页请求的路由，第二个参数是文件的路径
	router.StaticFile("/image", "static/image.jpg")
	//第一个参数是网页请求这个静态目录的前缀，第二个参数是该目录的路径，注意前缀不能重复
	router.StaticFS("/static", http.Dir("static/static"))

	router.GET("/string", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/baidu", _redirect)

	router.Run(":80")
}
