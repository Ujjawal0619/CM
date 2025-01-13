package couponmodal

import (
	database "github.com/ujjawal0619/cm/database/couponDB"
)

type CouponWithBxGy struct {
	database.Coupon
	database.BxGy
}

type Item struct {
	SKU      string  `json:"sku"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Cart struct {
	Items     []Item  `json:"items"`
	CartTotal float64 `json:"cartTotal"`
	Discount  float64 `json:"discount"`
}
