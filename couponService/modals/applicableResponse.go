package couponmodal

import (
	database "github.com/ujjawal0619/cm/database/couponDB"
)

type ApplicableCoupons struct {
	Cart              Cart              `json:"cart"`
	ApplicableCoupons []database.Coupon `json:"applicableCoupons"`
}
