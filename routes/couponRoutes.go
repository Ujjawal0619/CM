package routes

import (
	"github.com/gin-gonic/gin"
	couponhandler "github.com/ujjawal0619/cm/couponService/handler"
)

func CouponRoutes(router *gin.Engine, handler couponhandler.ICouponHandler) {
	router.POST("/coupons", handler.AddCoupon)
	router.GET("/coupons", handler.GetAllCoupon)
	router.GET("/coupons/:id", handler.GetCouponByID)
	router.PUT("coupons/:id", handler.UpdateCouponByID)
	router.DELETE("coupons/:id", handler.DeleteCouponByID)
	router.POST("applicable-coupons", handler.GetApplicableCoupons)
	router.POST("apply-coupon/:id", handler.ApplyCouponByID)
}
