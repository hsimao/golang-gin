package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)

	r.Run()
}

func testing(c *gin.Context) {
	var person Person
	// 會自動依據 Content-type 來做不同的數據 binding
	// 因此就算使用 get 或 post 的 param 參數或 body 都可解析到
	// curl -X GET 'localhost:8080/testing?name=Mars&address=高雄&birthday=2020-09-30'
	// curl -X POST 'localhost:8080/testing' -d 'name=Mars&address=高雄&birthday=2020-09-30'
	// curl -H 'Content-Type:application/json' -X POST 'localhost:8080/testing' -d '{"name":"Mars"}'
	if err := c.ShouldBind(&person); err != nil {
		c.String(200, "person bind error:%v", person)
	} else {
		c.String(200, "%v", person)
	}
}
