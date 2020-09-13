package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})

	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})

	// 自動創建所有 metdhos get、post、put、patch、delete...
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})

	r.Run()
}
