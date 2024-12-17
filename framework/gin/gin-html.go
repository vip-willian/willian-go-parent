package _gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinHtmlMain() {
	router := gin.Default()
	// 加载静态资源文件
	router.Static("/static", "framework/template/static")
	// 加载html页面模板
	router.LoadHTMLGlob("framework/template/html/*")
	router.GET("/index", func(c *gin.Context) {
		// 返回html
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Willian Page", "content": "123456"})
	})
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	router.Run(":9988")
}
