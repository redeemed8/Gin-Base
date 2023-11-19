package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	r.GET("/login", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		//	方法 1 直接获取
		//username := c.PostForm("username")
		//passwd := c.PostForm("passwd")

		//	方法 2 有默认值
		//	注意：这里的默认值是指找不到参数字段，而不是用户没填写用户名和密码
		//	用户名和密码没填的话，参数就是空字符串
		username := c.DefaultPostForm("username", "你没有名字吗?!")
		passwd := c.DefaultPostForm("passwd", "你连密码都没有吗?!")

		//	方法三
		//	username, ok := c.GetPostForm("username")
		//	passwd, _ := c.GetPostForm("passwd")

		data := gin.H{
			"username": username,
			"passwd":   passwd,
		}
		c.JSON(http.StatusOK, data)
	})

	err := r.Run(":9000")
	if err != nil {
		return
	}
}
