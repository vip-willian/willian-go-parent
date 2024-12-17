package _gin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func GinResponseMain() {
	router := gin.Default()
	// json数据格式响应
	jsonResponse(router)
	// 结构体数据格式响应
	structResponse(router)
	// xml数据格式响应
	xmlResponse(router)
	// yml数据格式响应
	ymlResponse(router)
	// protobuf数据格式响应
	protobufResponse(router)
	router.Run(":8090")
}

func jsonResponse(router *gin.Engine) {
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello json", "status": 200})
	})
}

func structResponse(router *gin.Engine) {
	router.GET("/struct", func(c *gin.Context) {
		var msg struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		msg.Code = 1001
		msg.Message = "hello world"
		c.JSON(200, msg)
	})
}

func xmlResponse(router *gin.Engine) {
	router.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "hello world"})
	})
}

func ymlResponse(router *gin.Engine) {
	router.GET("/yml", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "willian"})
	})
}

func protobufResponse(router *gin.Engine) {
	router.GET("/protobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
}
