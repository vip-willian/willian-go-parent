package _gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func GlobalMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		t := time.Now()
		log.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")

		status := c.Writer.Status()
		log.Println("中间件执行完毕", status)

		t2 := time.Since(t)
		log.Println("time: ", t2)

		c.Next()
		// c.Next() 放行
		// c.Abort() 中断
	}
}

func R2Middleware1() gin.HandlerFunc {

	return func(c *gin.Context) {
		log.Println("R2 第一个局部中间件 执行完成")
		c.Next()
	}
}

func R2Middleware2() gin.HandlerFunc {

	return func(c *gin.Context) {
		log.Println("R2 第二个局部中间件 执行完成")
		c.Next()
	}
}

func GinMiddleMain() {
	// 创建路由
	router := gin.Default()
	// 注册中间件
	router.Use(GlobalMiddleware())
	// 定义请求
	router.GET("/r1", func(c *gin.Context) {
		log.Println("在中间件执行完之后执行 r1")
		value, _ := c.Get("request")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("hello middle r1 %v", value),
		})
	})
	// 添加局部中间件 执行顺序跟添加顺序有关
	router.GET("/r2", R2Middleware1(), R2Middleware2(), func(c *gin.Context) {
		log.Println("在中间件执行完之后执行 r2")
		value, _ := c.Get("request")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("hello middle r2 %v", value),
		})
	})

	router.Run(":9320")
}
