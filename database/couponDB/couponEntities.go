package database

import "time"

type Coupon struct {
	ID            int64     `json:"id"`
	Code          string    `json:"code"`
	DiscoutType   string    `json:"discountType"`
	DiscountValue float64   `json:"discountValue"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	MinCartValue  float64   `json:"minCartValue"`
	AppliesToItem []string  `json:"appliesToItem"`
}
