package database

import (
	"encoding/json"
	"time"
)

type CouponType int

const (
	CART_WISE CouponType = iota
	PRODUCT_WISE
	BXGY
)

type Coupon struct {
	ID            int64           `json:"id"`
	Code          string          `json:"code"`
	DiscountType  CouponType      `json:"discountType"`
	DiscountValue float64         `json:"discountValue"`
	StartDate     time.Time       `json:"startDate"`
	EndDate       time.Time       `json:"endDate"`
	Details       json.RawMessage `json:"details"`
}

type BxGy struct {
	ID         int64    `json:"id"`
	CouponID   int64    `json:"couponId"`
	BxItemList []string `json:"bxItemList"`
	GyItemList []string `josn:"gyItemList"`
}
