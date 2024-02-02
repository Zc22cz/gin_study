package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoggerWithFormatter(params gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"[ jie ] %s  | %d | \t %s | %s | %s \t  %s\n",
		params.TimeStamp.Format("2006/01/02 - 15:04:05"),
		params.StatusCode, // 状态码
		params.ClientIP,   // 客户端ip
		params.Latency,    // 请求耗时
		params.Method,     // 请求方法
		params.Path,       // 路径
	)
}

func main() {
	//1.输入到文件
	//f, _ := os.Create("gin.log")
	////gin.DefaultWriter = io.MultiWriter(f)
	//// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//2.定义路由格式
	//gin.DebugPrintRouteFunc = func(
	//	httpMethod,
	//	absolutePath,
	//	handlerName string,
	//	nuHandlers int) {
	//	log.Printf(
	//		"[ jie ] %v %v %v %v\n",
	//		httpMethod,
	//		absolutePath,
	//		handlerName,
	//		nuHandlers,
	//	)
	//}

	gin.SetMode(gin.ReleaseMode)

	//3.修改log的显示

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(LoggerWithFormatter))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "ok"})
	})
	router.Run()
}
