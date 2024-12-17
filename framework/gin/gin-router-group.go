package _gin

import (
	"github.com/gin-gonic/gin"
	"willian-go-parent/framework/gin/goods"
	"willian-go-parent/framework/gin/order"
	"willian-go-parent/framework/gin/user"
)

func RouterGroup(r *gin.Engine) {

	// 用户接口分组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/get", user.GetUser)
		userGroup.POST("/add", user.AddUser)
		userGroup.POST("/login", user.Login)
	}
	// 订单接口分组
	orderGroup := r.Group("/order")
	{
		orderGroup.GET("/list", order.OrderList)
		orderGroup.POST("/create", order.CreateOrder)
	}
	// 接口版本管理
	goodsGroup1 := r.Group("/goods/v1")
	{
		goodsGroup1.GET("/query", goods.QueryGoodsV1)
		goodsGroup1.POST("/add", goods.AddGoodsV1)
	}
	goodsGroup2 := r.Group("/goods/v2")
	{
		goodsGroup2.GET("/query", goods.QueryGoodsV2)
		goodsGroup2.POST("/add", goods.QueryGoodsV2)
	}
}
