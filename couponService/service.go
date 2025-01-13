package couponservice

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	couponmodal "github.com/ujjawal0619/cm/couponService/modals"
	dbm "github.com/ujjawal0619/cm/database/couponDB"
)

type CouponService struct {
	DB dbm.Storage
}

type ICouponService interface {
	AddCoupon(*gin.Context) error
	GetAllCoupon(*gin.Context) ([]*dbm.Coupon, error)
	GetCouponByID(*gin.Context) (*dbm.Coupon, error)
	UpdateCouponByID(*gin.Context) error
	DeleteCouponByID(*gin.Context) error
	GetApplicableCoupons(*gin.Context) (*couponmodal.ApplicableCoupons, error)
	ApplyCouponByID(*gin.Context) (*couponmodal.Cart, error)
	validateCoupon(*dbm.Coupon) bool
}

func InitCouponService(db dbm.Storage) ICouponService {
	return &CouponService{
		db,
	}
}

func (s *CouponService) AddCoupon(c *gin.Context) error {
	var coupon dbm.Coupon
	var bxgy dbm.BxGy
	var couponWithBxgy couponmodal.CouponWithBxGy

	if err := c.BindJSON(&couponWithBxgy); err != nil {
		return err
	}

	coupon.Code = couponWithBxgy.Code
	coupon.DiscountType = couponWithBxgy.DiscountType
	coupon.DiscountValue = couponWithBxgy.DiscountValue
	coupon.Details = couponWithBxgy.Details
	coupon.StartDate = couponWithBxgy.StartDate
	coupon.EndDate = couponWithBxgy.EndDate

	if err := s.DB.CreateCoupon(&coupon); err != nil {
		return err
	}

	if couponWithBxgy.DiscountType == dbm.BXGY {
		newCoupon, err := s.DB.GetCouponByCode(coupon.Code)
		if err != nil {
			return err
		}
		bxgy.CouponID = newCoupon.ID
		bxgy.BxItemList = couponWithBxgy.BxItemList
		bxgy.GyItemList = couponWithBxgy.GyItemList

		if err := s.DB.CreateBxGyItem(&bxgy); err != nil {
			return err
		}
	}

	return nil
}

func (s *CouponService) GetAllCoupon(c *gin.Context) ([]*dbm.Coupon, error) {
	return s.DB.GetCoupons()
}

func (s *CouponService) GetCouponByID(c *gin.Context) (*dbm.Coupon, error) {
	couponId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}
	return s.DB.GetCouponByID(couponId)
}

func (s *CouponService) UpdateCouponByID(c *gin.Context) error {
	couponId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	var coupon dbm.Coupon
	if err := c.BindJSON(&coupon); err != nil {
		return err
	}

	return s.DB.UpdateCouponByID(couponId, &coupon)
}

func (s *CouponService) DeleteCouponByID(c *gin.Context) error {
	couponId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	return s.DB.DeleteCouponByID(couponId)
}

func (s *CouponService) GetApplicableCoupons(c *gin.Context) (*couponmodal.ApplicableCoupons, error) {
	applicableCoupon := &couponmodal.ApplicableCoupons{}
	var requestData struct {
		Cart couponmodal.Cart `json:"cart"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		return nil, err
	}

	applicableCoupon.Cart = requestData.Cart

	itemsSKU := map[string]bool{}

	for _, item := range applicableCoupon.Cart.Items {
		itemsSKU[item.SKU] = true
	}

	for k, v := range itemsSKU {
		fmt.Printf("key: %s, val: %v\n", k, v)
	}

	allCoupons, err := s.GetAllCoupon(c)
	if err != nil {
		return nil, err
	}
	for _, coupon := range allCoupons {

		if !s.validateCoupon(coupon) {
			continue
		}

		if coupon.DiscountType == dbm.CART_WISE {
			applicableCoupon.ApplicableCoupons = append(applicableCoupon.ApplicableCoupons, *coupon)
		}

		if coupon.DiscountType == dbm.PRODUCT_WISE {
			for _, sku := range coupon.Details {
				if _, isExist := itemsSKU[string(sku)]; isExist {
					applicableCoupon.ApplicableCoupons = append(applicableCoupon.ApplicableCoupons, *coupon)
				}
			}
		}
	}

	return applicableCoupon, nil
}

func (s *CouponService) ApplyCouponByID(c *gin.Context) (*couponmodal.Cart, error) {
	coupon, err := s.GetCouponByID(c)
	if err != nil {
		return nil, err
	}

	if !s.validateCoupon(coupon) {
		return nil, errors.New("Invalid Coupon")
	}

	var cart couponmodal.Cart
	if err = c.BindJSON(&cart); err != nil {
		return nil, err
	}

	var totalDiscount float64
	switch coupon.DiscountType {
	case dbm.CART_WISE:
		totalDiscount = 0 // TODO
	case dbm.PRODUCT_WISE:
		totalDiscount = 1 // TODO
	case dbm.BXGY:
		totalDiscount = 3 // TODO
	default:
		totalDiscount = 0
	}

	cart.Discount = totalDiscount

	return &cart, nil
}

func (s *CouponService) AddBxGy(c *gin.Context) {

}

func (s *CouponService) validateCoupon(c *dbm.Coupon) bool {
	now := time.Now()
	if c.EndDate.After(now) {
		log.Println("Coupon is not available at the moment")
		return false
	}

	if c.EndDate.Before(now) {
		log.Println("Coupon has expired")
		return false
	}

	if c.DiscountType == dbm.BXGY {
		bxgy, err := s.DB.GetBxGyItemsByID(int(c.ID))
		if err != nil || len(bxgy.BxItemList) == 0 || len(bxgy.GyItemList) == 0 {
			log.Printf("BxGy items does not exist for couponId: %d\n", c.ID)
			return false
		}
	}

	return true
}

// {
// 	"code": "NEWYEAR20",
// 	"discountType": 2,
// 	"discountValue": 20,
// 	"startDate": "2024-02-15T00:00:00Z",
// 	"endDate": "2025-03-15T00:00:00Z",
// 	"details": {"disc": "use bxgy_items table for more details"},
// 	"couponId": 15,
// 	"bxItemList": ["SKU4", "SKU5"],
// 	"gyItemList": ["SKU1", "SKU2", "SKU3"]
//   }
