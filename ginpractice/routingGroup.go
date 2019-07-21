/*把一个模块相关的方法都写在一个路由下,便于管理和查找相关的代码*/
package main

import "github.com/gin-gonic/gin"
import "fmt"

func loginEndpoint(c *gin.Context){
	fmt.Println("这是login方法")
}

func submitEndpoint(c *gin.Context){
	fmt.Println("这是submit方法")
}

func readEndpoint(c *gin.Context){
	fmt.Println("这是read方法")
}

func main() {
	router := gin.Default()
	//v1组路由
	v1:=router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}

	//v2组路由
	v2:=router.Group("/v2")
	{
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.GET("/read", readEndpoint)
	}
	router.Run(":8888")
}

/*
浏览器访问
http://localhost:8080/v1/login
http://localhost:8080/v1/submit
http://localhost:8080/v1/read
http://localhost:8080/v2/login
http://localhost:8080/v2/submit
http://localhost:8080/v2/read
*/
