package _gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GinRouterMain() {
	// 1、创建路由
	r := gin.Default()
	// 路径参数
	GetPathVariableParam(r)
	// 请求参数
	GetRequestParam(r)
	// 表单参数
	FormParam(r)
	// 文件上传参数
	FileUploadParam(r)
	// 多文件上传参数
	MultiFileUploadParam(r)
	// 路由组
	RouterGroup(r)
	// 3.监听端口，默认在8080
	r.Run(":8082")
}

// GetPathVariableParam 获取路径参数
func GetPathVariableParam(r *gin.Engine) {
	r.GET("/user/v1/:username/:password/*other", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		other := c.Param("other") // /*other
		other = strings.Trim(other, "/")
		c.String(http.StatusOK, "%s:%s:%s", username, password, other)
	})
}

// GetRequestParam 请求参数
func GetRequestParam(r *gin.Engine) {
	r.GET("/user/v2", func(c *gin.Context) {
		username := c.DefaultQuery("username", "admin")
		password := c.DefaultQuery("password", "123456")
		c.String(http.StatusOK, "%s:%s", username, password)
	})
}

// FormParam 表单参数
func FormParam(r *gin.Engine) {
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"type":     types,
		})
	})
}

// FileUploadParam 单文件上传
func FileUploadParam(r *gin.Engine) {
	// 限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/file-upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusInternalServerError, "文件上传失败")
		}
		// 保存文件
		err = c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "文件上传失败")
		}
		// 需要补充，限制文件大小以及限制文件上传类型
		// 返回响应
		c.String(http.StatusOK, file.Filename)
	})
}

// MultiFileUploadParam 多文件上传参数
func MultiFileUploadParam(r *gin.Engine) {
	// 限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/multi-file-upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		for _, file := range files {
			// 保存文件
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		// 返回响应
		c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
	})
}
