/*路由 gin框架中采用的路由库是httprouter*/
package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"

func getting (c *gin.Context){
	name := c.Param("name")
	fmt.Printf("Hello %s", name)
	c.JSON(200, gin.H{
		"message" : "pong",
	})
}

func main() {
	// 带有默认中间件的路由     gin.New()不带中间件的路由
	router := gin.Default()

	router.GET("/string/:name", getting)
	/*请求响应匹配的URL: /welcome?firstname=Jane&lastname=Doe*/
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	router.Run(":8888")
}