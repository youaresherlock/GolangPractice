package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//gin.H是一个map[string] interface{}的快捷方式
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message" : "hey", "status" : http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		//使用结构体
		var msg struct {
			Name string `json:"user"`
			Message string 
			Number int
		}
		msg.Name = "lena"
		msg.Message = "hey"
		msg.Number = 123
		//msg.Name在JSON中会变成"user"
		// 将会输出： {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message" : "hey", "status" : http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message" : "hey", "status" : http.StatusOK})
	})

	r.Run(":8888")
}