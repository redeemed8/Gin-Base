package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/toBaidu", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})
		//	重定向,地址栏会发生变化
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.GET("/aaa", func(c *gin.Context) {
		//	转发给 bbb,地址栏不会发生变化
		c.Request.URL.Path = "/bbb"
		r.HandleContext(c) //	继续后续的处理
	})

	r.GET("/bbb", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"name":   "bbbbbbbbbbbbbbbbbb",
		})
	})

	r.Run(":9000")
}
