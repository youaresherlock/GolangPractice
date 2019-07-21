/*
To bind a request body into a type, use model binding. We currently support binding of JSON, XML, YAML and standard form values (foo=bar&boo=baz).
*/
package main

import "net/http"
import "github.com/gin-gonic/gin"

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login

		//绑定JSON的例子({"user": "manu", "password": "123"}) postman发送json后成功接收到json数据
		if c.BindJSON(&json) == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status" : "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status" : "unauthorized"})
			}
		}
	})

	//一个HTML表单绑定的示例 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 根据请求头中content-type 自动推断
		if c.Bind(&form) == nil {
			if form.User == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{
					"status" : "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status" : "unauthorized"})
			}
		}
	})

	router.Run(":8888")
}