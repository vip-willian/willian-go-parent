package _gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLogin struct {
	UserName string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func GinDataBindMain() {
	router := gin.Default()
	// Json数据绑定与解析
	jsonDataBind(router)
	// FORM数据绑定与解析
	formDataBind(router)
	// URI数据绑定与解析
	uriDataBind(router)
	router.Run(":8089")
}

// Json数据绑定与解析
func jsonDataBind(router *gin.Engine) {
	router.POST("/user-login/v1", func(c *gin.Context) {
		var user UserLogin
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if user.UserName != "root" || user.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}

// FORM数据绑定与解析
func formDataBind(router *gin.Engine) {
	router.POST("/user-login/v2", func(c *gin.Context) {
		var user UserLogin
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if user.UserName != "root" || user.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}

// URI数据绑定与解析
func uriDataBind(router *gin.Engine) {
	router.GET("/user-login/v3/:username/:password", func(c *gin.Context) {
		var user UserLogin
		if err := c.ShouldBindUri(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if user.UserName != "root" || user.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}
