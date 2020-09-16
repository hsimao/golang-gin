package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 修改 gin Logger 中間件配置, 將 log 寫入到指定檔案
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(200, "%s", name)
	})

	r.Run()
}
