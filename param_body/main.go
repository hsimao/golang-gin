package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyByts, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}

		// 使用 ioutil.ReadAll 讀取完資料後就沒辦法用 PostForm 取得值
		// 所以要再將取出的 bodyByts 回寫進去 Request.Body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByts))

		name := c.PostForm("name")
		age := c.PostForm("age")
		c.String(http.StatusOK, "%s.%s,%s", name, age, string(bodyByts))
		// c.String(http.StatusOK, string(bodyByts))
	})

	r.Run()
}
