package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 定义一个中间件 indexHandler
func indexHandler(c *gin.Context) {
	name, ok := c.Get("name") //	从上下文中取值 (跨中间件存取值)
	if !ok {
		name = "你连名字都没有?!"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "this is index22.. " + fmt.Sprintf("%v", name),
	})
}

// 定义一个中间件 m1
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	//	计时
	start := time.Now()

	//go funcXX(c.Copy())		这里必须使用 c的拷贝

	c.Next() //	调用后续的处理函数
	//c.Abort() //	阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cosg:%v\n", cost)
	fmt.Println("m1 out...")
}

// 定义一个中间件m2
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "jchhh")
	//c.Abort()
	c.Next()
	fmt.Println("m2 out ...")
}

func authMiddleware(isLogin bool) gin.HandlerFunc {
	//	写成闭包的形式	可以做一些其他的逻辑
	//	比如连接数据库
	return func(c *gin.Context) {
		//	 是否登录的判断
		if isLogin {
			c.Next()
		} else {
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default() //	默认使用了 Logger 和 Recovery 中间件
	//	如果不想使用默认的中间件，那么请使用 gin.New()

	r.Use(m1, m2) //	全局注册中间件 m1 ,每个请求都可用

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/index2", indexHandler)

	//	路由组注册中间件方法 1
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}

	//	路由组注册中间件方法 2
	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
		})
	}

	r.Run(":9000")
}
