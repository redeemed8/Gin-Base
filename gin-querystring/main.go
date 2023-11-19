package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//	获取浏览器发请求时携带的querystring参数
		//	这个参数就是指的 url中 ?后面的参数
		name := c.Query("name")
		name2 := c.DefaultQuery("name2", "sb2")
		name3, ok := c.GetQuery("name3")
		if !ok {
			name3 = "sb3"
		}
		c.JSON(http.StatusOK, gin.H{
			"name":  name,
			"name2": name2,
			"name3": name3,
		})
	})

	err := r.Run(":9000")
	if err != nil {
		return
	}
}
