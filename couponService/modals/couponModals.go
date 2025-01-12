package couponmodal

import (
	database "github.com/ujjawal0619/cm/database/couponDB"
)

type CouponWithBxGy struct {
	database.Coupon
	database.BxGy
}
