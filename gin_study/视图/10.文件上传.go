package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		//1.SaveUploadedFile保存在指定路径
		//fmt.Println(file.Filename)
		//fmt.Println(file.Size / 1024)
		//c.SaveUploadedFile(file, "uploads/1.png")
		//2.读取上传文件
		//readerFile, _ := file.Open()
		//data, _ := io.ReadAll(readerFile)
		//fmt.Println(string(data))
		//3.上传单文件
		readerFile, _ := file.Open()
		writerFile, _ := os.Create("uploads/2.jpg")
		defer writerFile.Close()
		io.Copy(writerFile, readerFile)
		c.JSON(200, gin.H{"msg": "上传成功"})
	})
	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files, _ := form.File["upload[]"]
		for _, file := range files {
			c.SaveUploadedFile(file, "uploads/"+file.Filename)
		}
		c.JSON(200, gin.H{"msg": fmt.Sprintf("成功上传%d个文件", len(files))})
	})
	router.Run(":8080")
}
