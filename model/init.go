package model

import "time"

var Users = []UserInDB{
	{ID: "1", Username: "user1", Password: "password1", VipId: "1", Points: 1000, Coins: 10},
	{ID: "2", Username: "user2", Password: "password2", VipId: "2", Points: 200, Coins: 25},
	{ID: "3", Username: "user3", Password: "password3", VipId: "", Points: 30, Coins: 10},
}

var Vips = []VipInDB{
	{ID: "1", Name: "VIP1", Discount: 0.95},
	{ID: "2", Name: "VIP2", Discount: 0.9},
	{ID: "3", Name: "VIP3", Discount: 0.85},
}

var vip1 = VipDiscount{"1", 0.9}
var vip2 = VipDiscount{"2", 0.85}
var vip3 = VipDiscount{"3", 0.8}
var pointDiscount = PointDiscount{1.0, 0}

var Activities = []ActivityInDB{
	{ID: "1", VipDiscounts: []VipDiscount{vip1, vip2, vip3}, PointDiscount: pointDiscount,
		StartAt: time.Date(2022, 5, 1, 0, 0, 0, 0, time.Local), EndAt: time.Date(2022, 5, 10, 23, 59, 59, 0, time.Local)},
}

var Products = []ProductInDB{
	{ID: "1", Name: "Product1", Description: "some description", Price: 10},
	{ID: "2", Name: "Product2", Description: "some description", Price: 35},
}
