package model

type ProductRespone struct {
	ProductId     string `json:"id"`
	Name          string `json:"name"`
	Price         int64  `json:"price"`
	ImageUrl      string `json:"imageUrl"`
	Stock         int16  `json:"stock"`
	Condition     string `json:"condition"`
	Tags          string `json:"tags"`
	IsPurchasable bool   `json:"isPurchasable"`
	PurchaseCount int8   `json:"purchaseCount"`
}

type ProductRequest struct {
	Name          string `json:"name"`
	Price         int64  `json:"price"`
	ImageUrl      string `json:"imageUrl"`
	Stock         int16  `json:"stock"`
	Condition     string `json:"condition"`
	Tags          string `json:"tags"`
	IsPurchasable bool   `json:"isPurchasable"`
}

type ProductFilter struct {
	Keyword *string `json:"keyword"`
}
