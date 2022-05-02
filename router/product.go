package router

import (
	"example/crud"
	"example/model"
	"example/service"
	"example/validator"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPong(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Purchase(request *gin.Context) {
	productID := request.Param("id")
	var purchase validator.PurchaseRequest
	if err := request.BindJSON(&purchase); err != nil {
		return
	}

	product := crud.GetProduct(productID)
	if product == (model.ProductInDB{}) {
		request.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	user := crud.GetUser(purchase.UserID)
	fmt.Println(user)
	if user == (model.UserInDB{}) {
		request.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "user not found"})
		return
	}

	activity := crud.GetActivity()
	discountPointValid := service.CheckIsValidDiscountPoint(purchase.DiscountPoints, user, activity)
	if discountPointValid == false {
		request.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "invalid discount points"})
	} else {
		finalDiscountPoints, finalCost := service.Calculate(purchase, product, user, activity)
		if leftCoin := user.Coins - finalCost; leftCoin < 0 {
			request.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "insufficient balance"})
			return
		} else {
			user.Points -= finalDiscountPoints
			user.Coins -= finalCost
			fmt.Println(user)
			request.IndentedJSON(http.StatusOK, gin.H{
				"message": "purchase successfully",
			})
		}
	}
}
