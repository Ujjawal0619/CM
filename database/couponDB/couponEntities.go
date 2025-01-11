package database

import "time"

type CouponType int

const (
	CART_WISE CouponType = iota
	PRODUCT_WISE
	BXGY
)

type Coupon struct {
	ID            int64      `json:"id"`
	Code          string     `json:"code"`
	DiscoutType   CouponType `json:"discountType"`
	DiscountValue float64    `json:"discountValue"`
	StartDate     time.Time  `json:"startDate"`
	EndDate       time.Time  `json:"endDate"`
}
