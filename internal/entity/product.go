package entity

type Product struct {
	ID            string
	Name          string
	Price         string
	ImageUrl      string
	Stock         int16
	Condition     string
	IsPurchasable bool
	PurchaseCount int8
	UserId        string
}

func (prod *Product) TableName() string {
	return "products"
}
