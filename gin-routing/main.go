package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//	访问 /index的 GET请求会走这一条处理逻辑
	//	路由
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	//	any能处理所有的请求方式
	r.Any("/all", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
		default:
			c.JSON(http.StatusOK, gin.H{"method": "OTHERS"})
		}
	})

	//	所有未定义的路由默认是404, 但是可以用这个匹配所有的未定义的url
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "ggaggaa"})
	})

	//	路由组
	//	把公用的前缀提出来每创建一个路由组
	userGroup := r.Group("/user")

	//	每次调用路由组
	//	/user/xxx
	userGroup.GET("/xxx", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "GET = =`"})
	})
	//	/user/jchhh
	userGroup.POST("/jchhh", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "POST = =||"})
	})

	r.Run(":9000")
}
