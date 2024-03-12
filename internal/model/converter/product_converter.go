package converter

import (
	"golang-marketplace/internal/entity"
	"golang-marketplace/internal/model"
)

func ProductConverter(product *entity.Product) *model.ProductRespone {
	return &model.ProductRespone{
		ProductId: product.ID,
		Name:      product.Name,
		Price:     product.Price,
		ImageUrl:  product.ImageUrl,
		Stock:     product.Stock,
		Condition: product.Condition,
		// Tags:          product.,
		IsPurchasable: product.IsPurchasable,
		PurchaseCount: product.PurchaseCount,
	}
}
