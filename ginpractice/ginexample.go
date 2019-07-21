/*assume the following codes in example.go file*/
package main 

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message" : "pong",
		})
	})
	// r.Run() // 在 0.0.0.0:8080 上监听并服务
	http.ListenAndServe(":8888", r)
}