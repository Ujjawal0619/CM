package couponservice

import (
	"time"

	"github.com/gin-gonic/gin"
	database "github.com/ujjawal0619/cm/database/couponDB"
)

type CouponService struct {
	DB database.Storage
}

type ICouponService interface {
	AddCoupon(c *gin.Context)
	GetAllCoupon(c *gin.Context) ([]database.Coupon, error)
	GetCouponByID(c *gin.Context)
	UpdateCouponByID(c *gin.Context)
	DeleteCouponByID(c *gin.Context)
}

func InitCityService(db database.Storage) ICouponService {
	return &CouponService{
		db,
	}
}

func (h *CouponService) AddCoupon(c *gin.Context) {

}
func (h *CouponService) GetAllCoupon(c *gin.Context) ([]database.Coupon, error) {
	return []database.Coupon{
		{
			ID:            123,
			Code:          "XYZ",
			DiscoutType:   database.PRODUCT_WISE,
			DiscountValue: 10.0,
			StartDate:     time.Now(),
			EndDate:       time.Now(),
			MinCartValue:  100.0,
			AppliesToItem: []string{"first", "second"},
		},
	}, nil
}
func (h *CouponService) GetCouponByID(c *gin.Context) {

}
func (h *CouponService) UpdateCouponByID(c *gin.Context) {

}
func (h *CouponService) DeleteCouponByID(c *gin.Context) {

}
