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
	GetApplicableCoupons(c *gin.Context)
	ApplyCouponByID(c *gin.Context)
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
	coupon, err := h.couponService.GetCouponByID(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "coupon does not exist with id:" + c.Param("id")})
	} else {
		c.JSON(http.StatusOK, coupon)
	}
}
func (h *CouponHandler) UpdateCouponByID(c *gin.Context) {
	err := h.couponService.UpdateCouponByID(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid coupon id:" + c.Param("id")})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "coupon id:" + c.Param("id") + " updated"})
	}
}
func (h *CouponHandler) DeleteCouponByID(c *gin.Context) {
	err := h.couponService.DeleteCouponByID(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid coupon id:" + c.Param("id")})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "coupon id:" + c.Param("id") + " deleted"})
	}
}

func (h *CouponHandler) GetApplicableCoupons(c *gin.Context) {
	applicableCoupon, err := h.couponService.GetApplicableCoupons(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Something went wrong."})
	} else {
		c.JSON(http.StatusOK, applicableCoupon)
	}
}

func (h *CouponHandler) ApplyCouponByID(c *gin.Context) {
	// TODO
}
