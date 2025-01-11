package couponservice

import (
	"github.com/gin-gonic/gin"
	database "github.com/ujjawal0619/cm/database/couponDB"
)

type CouponService struct {
	DB database.Storage
}

type ICouponService interface {
	AddCoupon(c *gin.Context)
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

func (h *CouponService) AddCoupon(c *gin.Context) {

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
