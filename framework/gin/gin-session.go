package _gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinCookieSessionMain() {
	// 1、创建路由
	router := gin.Default()
	// 2、获取cookie
	GetCookie(router)
	// 3.监听端口，默认在8080
	router.Run(":8086")
}

func GetCookie(router *gin.Engine) {

	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("willian_cookie")
		if err != nil {
			cookie = "Not Found"
			// 设置cookie
			// maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			// secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("willian_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		// 获取Cookie的值
		fmt.Printf("cookie的值是: %s\n", cookie)
	})

	router.GET("/home", CookieAuthMiddleMare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})

	router.GET("/login", func(c *gin.Context) {
		// 跳过账号密码验证
		// 登录成功后 设置cookie
		c.SetCookie("user", "willian", 60, "/", "localhost", false, true)
		c.JSON(200, gin.H{"data": "login successful"})
	})
}

// CookieAuthMiddleMare 认证中间件
func CookieAuthMiddleMare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取cookie信息
		if user, err := c.Cookie("user"); err == nil {
			if user == "willian" {
				// 用户等于willian放行
				c.Next()
				return
			}
		}
		// 进行拦截
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		c.Abort()
		return
	}
}
