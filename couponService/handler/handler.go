package couponhandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	couponservice "github.com/ujjawal0619/cm/couponService"
)

type CouponHandler struct {
	couponService couponservice.ICouponService
}

type ICouponHandler interface {
	AddCoupon(c *gin.Context)
	GetAllCoupon(c *gin.Context)
	GetCouponByID(c *gin.Context)
	UpdateCouponByID(c *gin.Context)
	DeleteCouponByID(c *gin.Context)
}

func InitHandler(couponService couponservice.ICouponService) ICouponHandler {
	return &CouponHandler{
		couponService,
	}
}

func (h *CouponHandler) AddCoupon(c *gin.Context) {
	err := h.couponService.AddCoupon(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "Coupon created successfully"})
	}
}
func (h *CouponHandler) GetAllCoupon(c *gin.Context) {
	coupons, err := h.couponService.GetAllCoupon(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "no coupon found"})
	} else {
		c.JSON(http.StatusOK, coupons)
	}
}

func (h *CouponHandler) GetCouponByID(c *gin.Context) {

}
func (h *CouponHandler) UpdateCouponByID(c *gin.Context) {

}
func (h *CouponHandler) DeleteCouponByID(c *gin.Context) {

}
