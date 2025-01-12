package couponservice

import (
	"fmt"

	"github.com/gin-gonic/gin"
	couponmodal "github.com/ujjawal0619/cm/couponService/modals"
	database "github.com/ujjawal0619/cm/database/couponDB"
)

type CouponService struct {
	DB database.Storage
}

type ICouponService interface {
	AddCoupon(c *gin.Context) error
	GetAllCoupon(c *gin.Context) ([]*database.Coupon, error)
	GetCouponByID(c *gin.Context)
	UpdateCouponByID(c *gin.Context)
	DeleteCouponByID(c *gin.Context)
}

func InitCouponService(db database.Storage) ICouponService {
	return &CouponService{
		db,
	}
}

func (h *CouponService) AddCoupon(c *gin.Context) error {
	var coupon database.Coupon
	var bxgy database.BxGy
	var couponWithBxgy couponmodal.CouponWithBxGy

	if err := c.BindJSON(&couponWithBxgy); err != nil {
		return err
	}

	coupon.Code = couponWithBxgy.Code
	coupon.DiscoutType = couponWithBxgy.DiscoutType
	coupon.DiscountValue = couponWithBxgy.DiscountValue
	coupon.StartDate = couponWithBxgy.StartDate
	coupon.EndDate = couponWithBxgy.EndDate

	bxgy.CouponID = couponWithBxgy.CouponID
	bxgy.BxItemList = couponWithBxgy.BxItemList
	bxgy.GyItemList = couponWithBxgy.GyItemList

	fmt.Println(couponWithBxgy, coupon, bxgy)

	if err := h.DB.CreateCoupon(&coupon); err != nil {
		return err
	}

	if err := h.DB.CreateBxGyItem(&bxgy); err != nil {
		return err
	}

	return nil
}

func (h *CouponService) GetAllCoupon(c *gin.Context) ([]*database.Coupon, error) {
	return h.DB.GetCoupons()
}

func (h *CouponService) GetCouponByID(c *gin.Context) {

}

func (h *CouponService) UpdateCouponByID(c *gin.Context) {

}

func (h *CouponService) DeleteCouponByID(c *gin.Context) {

}
