package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name    string `json:"name"` //	称为结构体的 tag
	Message string
	Age     int
}

func main() {

	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//	方法 1 使用 map
		//data := map[string]interface{}{
		//	"name":    "jchhh",
		//	"message": "哈哈哈哈",
		//	"age":     18,
		//}
		data := gin.H{
			"name":    "jchhh",
			"message": "哈哈哈哈",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/json2", func(c *gin.Context) {
		//	方法 2 使用 结构体
		data := Person{
			"jchhh2",
			"哈哈哈哈",
			18,
		}
		c.JSON(http.StatusOK, data)
	})

	err := r.Run(":9000")
	if err != nil {
		return
	}
}
