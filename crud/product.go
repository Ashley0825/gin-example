package crud

import "example/model"

func GetProduct(productID string) model.ProductInDB {
	for _, p := range model.Products {
		if p.ID == productID {
			return p
		}
	}
	return model.ProductInDB{}
}
