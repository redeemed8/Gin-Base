package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//	获取请求的path(URI)参数，返回的都是字符串类型

func main() {
	r := gin.Default()

	r.GET("/web/:name/:age", func(c *gin.Context) {

		//	 获取路径参数
		name := c.Param("name")
		age := c.Param("age")

		data := gin.H{
			"name": name,
			"age":  age,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9000")
}
