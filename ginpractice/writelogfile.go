package main

import "os"
import "io"
import "net/http"
import "github.com/gin-gonic/gin"

func main() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()

	//写入日志的文件
	f, _ := os.Create("gin.log")
	// By default gin.DefaultWriter = os.Stdout
	// 只写入日志文件gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run(":8888")
}
// 