package _gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func GoAsyncMain() {
	router := gin.Default()
	// 异步处理调用
	AsyncRequest(router)
	// 同步处理调用
	SyncRequest(router)
	err := router.Run(":4464")
	if err != nil {
		fmt.Println("启动失败,监听端口: 4464")
		panic(err)
	}
}

func SyncRequest(router *gin.Engine) {
	router.GET("/async/request", func(c *gin.Context) {
		// 异步执行，需要copy一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})
}

func AsyncRequest(router *gin.Engine) {
	router.GET("/sync/request", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})
}
