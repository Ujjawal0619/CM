package routes

import (
	"github.com/gin-gonic/gin"
	couponhandler "github.com/ujjawal0619/cm/couponService/handler"
)

func CouponRoutes(router *gin.Engine, handler couponhandler.ICouponHandler) {
	router.POST("/coupons", handler.AddCoupon)
	router.GET("/coupons", handler.GetAllCoupon)
	router.GET("/coupons/:id", handler.GetCouponByID)
	router.POST("/coupons/bxgy", handler.AddBxGy)
}
