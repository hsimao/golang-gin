package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Hello world",
		})
	})

	r.Run()
}
