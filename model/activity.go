package model

import (
	"time"
)

type VipDiscount struct {
	VipId    string
	Discount float64
}

type PointDiscount struct {
	Discount float64
	Limit    uint
}

type ActivityInDB struct {
	ID            string
	VipDiscounts  []VipDiscount
	PointDiscount PointDiscount
	StartAt       time.Time
	EndAt         time.Time
}
