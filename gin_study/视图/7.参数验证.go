package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name     string   `json:"name" binding:"endswith=f"`
	Age      int      `json:"age" binding:"lt=30,gt=18"`
	Sex      string   `json:"sex" binding:"oneof=man women"`
	LikeList []string `json:"like_list" binding:"dive,startswith=like"`
	IP       string   `json:"ip" binding:"ip"`
	Url      string   `json:"url" binding:"url"`
	Uri      string   `json:"uri" binding:"uri"`
	Date     string   `json:"date" binding:"datetime=2006-01-02 15:04:05"`
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": user})
	})
	router.Run(":80")
}
