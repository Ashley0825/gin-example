package service

import (
	"example/model"
	"example/validator"
	"math"
)

func CheckIsValidDiscountPoint(discountPoints uint, user model.UserInDB, activity model.ActivityInDB) bool {
	if discountPoints > user.Points {
		return false
	}

	if (&activity) != nil {
		if activity.PointDiscount.Limit > 0 && discountPoints > 0 {
			return false
		}
	}
	return true
}

func Calculate(purchase validator.PurchaseRequest, product model.ProductInDB, user model.UserInDB, activity model.ActivityInDB) (uint, uint) {
	cost := float64(product.Price)
	discount := 1.0
	pointDiscount := 1.0
	if (&activity) != nil {
		pointDiscount = activity.PointDiscount.Discount
	}

	if user.VipId != "" {
		if (&activity) != nil {
			for _, d := range activity.VipDiscounts {
				if d.VipId == user.VipId {
					discount = d.Discount
				}
			}
		} else {
			for _, v := range model.Vips {
				if v.ID == user.VipId {
					discount = v.Discount
				}
			}
		}
		cost = cost * discount
	}

	finalDiscountPoints := uint(math.RoundToEven(float64(purchase.DiscountPoints) * pointDiscount))
	finalCost := uint(math.RoundToEven(cost)) - finalDiscountPoints

	return finalDiscountPoints, finalCost
}
