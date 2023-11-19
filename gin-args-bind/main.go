package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Passwd   string `form:"passwd"`
}

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		//	 shouldBind	能根据你请求的contentType的类型来分别解析你的数据
		//	可以是querystring，form表单，和json

		var u UserInfo          //	 声明一个 UserInfo类型的变量 u
		err := c.ShouldBind(&u) //	 进行绑定,要用引用传递

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {

			c.JSON(http.StatusOK, gin.H{
				"status":   "ok",
				"username": u.Username,
				"passwd":   u.Passwd,
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {
		//username := c.Query("username")
		//passwd := c.Query("passwd")
		//data := UserInfo{username, passwd}

		var u UserInfo          //	 声明一个 UserInfo类型的变量 u
		err := c.ShouldBind(&u) //	 进行绑定,要用引用传递

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {

			c.JSON(http.StatusOK, gin.H{
				"status":   "ok",
				"username": u.Username,
				"passwd":   u.Passwd,
			})
		}
	})

	err := r.Run(":9000")
	if err != nil {
		return
	}
}
