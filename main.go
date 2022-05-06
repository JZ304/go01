package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	// r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"Hello": "World!", "name": "zjx",
	// 	})
	// })
	// r.Run()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "world!",
		})
	})

	// 404处理
	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error_code": 404,
				"error_info": "路由未定义",
			})
		}
	})

	r.Run(":8000")
}
