package repository

import (
	"fmt"
	"golang-marketplace/internal/entity"
	"golang-marketplace/internal/model"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) List(db *sqlx.DB, filter *model.ProductFilter) ([]entity.Product, error) {
	product := []entity.Product{}

	// ToDo: Get by request param
	// query := db.Preload("Category")

	// if *filter.CategoryId != 0 {
	// 	query.Where("category_id = ?", &filter.CategoryId)
	// }

	if *filter.Keyword != "" {
		query := `SELECT * FROM products WHERE name LIKE CONCAT('%', $1::TEXT, '%')`
		err := db.Select(&product, query, *filter.Keyword)
		if err != nil {
			// Return empty object and error.
			return nil, err
		}
	} else {
		query := `SELECT * FROM products`
		err := db.Select(&product, query)
		if err != nil {
			// Return empty object and error.
			return nil, err
		}
	}

	// if err := query.Find(&product).Error; err != nil {
	// 	return nil, err
	// }

	return product, nil
}

func (r *ProductRepository) Get(db *sqlx.DB, id string, product *entity.Product) (entity.Product, error) {
	query := `SELECT * FROM products WHERE id = $1`
	productData := entity.Product{}
	err := db.Get(&productData, query, id)

	if err != nil {
		return productData, err
	}

	return productData, nil
}

func (r *ProductRepository) Create(db *sqlx.DB, request *entity.Product) error {
	query := `INSERT INTO products VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := db.Exec(query, "dec7fff9-2842-41f2-812c-1ae219f95fba", request.Name, request.Price, request.ImageUrl, request.Stock, request.Condition, request.IsPurchasable, "1")
	return err
}

func (r *ProductRepository) Update(db *sqlx.DB, id string, request *entity.Product) error {

	query := `UPDATE products SET name = $2, price = $3, imageUrl = $4, condition = $5, isPurchasable = $6 WHERE id = $1`

	_, err := db.Exec(query, id, request.Name, request.Price, request.ImageUrl, request.Condition, request.IsPurchasable)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Delete(db *sqlx.DB, request *entity.Product) error {
	query := fmt.Sprintf("DELETE FROM products WHERE id=%s", request.ID)
	// Send query to database.
	_, err := db.Exec(query)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
