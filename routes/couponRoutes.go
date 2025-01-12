package routes

import (
	"github.com/gin-gonic/gin"
	couponhandler "github.com/ujjawal0619/cm/couponService/handler"
)

func CouponRoutes(router *gin.Engine, handler couponhandler.ICouponHandler) {
	router.GET("/coupons", handler.GetAllCoupon)
	router.POST("/coupons", handler.AddCoupon)
}
