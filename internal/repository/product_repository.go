package repository

import (
	"fmt"
	"golang-marketplace/internal/entity"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) List(db *sqlx.DB) ([]entity.Product, error) {
	var books []entity.Product

	// ToDo: Get by request param
	// query := db.Preload("Category")

	// if *filter.CategoryId != 0 {
	// 	query.Where("category_id = ?", &filter.CategoryId)
	// }

	// if *filter.Keyword != "" {
	// 	query.Where("title LIKE ?", "%"+*filter.Keyword+"%")
	// }

	// if err := query.Find(&books).Error; err != nil {
	// 	return nil, err
	// }

	return books, nil
}

func (r *ProductRepository) Get(db *sqlx.DB, id int, book *entity.Product) error {

	query := fmt.Sprintf("SELECT * FROM PRODUCT WHERE id = %d ", id)

	_, err := db.Queryx(query)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Create(db *sqlx.DB, request *entity.Product) error {
	query := fmt.Sprintf("INSERT INTO products (id, name, price, image_url, stock, condition, is_purchasable, purchase_count) VALUES (%s, %s, %s, %s, %d, %s, %t, %d)", request.ID, request.Name, request.Price, request.ImageUrl, request.Stock, request.Condition, request.IsPurchasable, request.PurchaseCount)

	_, err := db.Exec(query)
	return err
}

func (r *ProductRepository) Update(db *sqlx.DB, id int, request *entity.Product) error {
	query := fmt.Sprintf(`UPDATE products SET name = %s, price = %s, image_url = %s, stock = %d, condition = %s, is_purchasable=%t WHERE id = %s`, request.Name, request.Price, request.ImageUrl, request.Stock, request.Condition, request.IsPurchasable, request.ID)

	_, err := db.Exec(query)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

func (r *ProductRepository) Delete(db *sqlx.DB, request *entity.Product) error {
	query := fmt.Sprintf("DELETE FROM producs WHERE id=%s", request.ID)
	// Send query to database.
	_, err := db.Exec(query)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
