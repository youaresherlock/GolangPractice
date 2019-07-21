/*upload single file*/
package main

import (
	"fmt"
	"log"
	"net/http"
	// "path/filepath"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB default is 32 MiB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context){
		file, _ := c.FormFile("files")
		log.Println(file.Filename)

		// upload the file to specific dst
		c.SaveUploadedFile(file, "test.txt") // 将文件下载到当前目录中名字为test.txt

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8888")
}