package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" validate:"required,bookableDate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" validate:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// 自定義驗證
func bookableDate(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

func main() {
	r := gin.Default()

	// 註冊自定義驗證 bookableDate func
	validate := validator.New()
	validate.RegisterValidation("bookableDate", bookableDate)

	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(&b); err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := validate.Struct(b); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		c.JSON(200, gin.H{
			"message": "ok!",
			"booking": b,
		})
	})

	r.Run()
}
