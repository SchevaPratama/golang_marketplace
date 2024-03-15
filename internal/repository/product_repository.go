package repository

import (
	"fmt"
	"golang-marketplace/internal/entity"
	"golang-marketplace/internal/model"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) List(filter *model.ProductFilter) ([]entity.Product, error) {
	tx, _ := r.DB.Beginx()
	defer tx.Rollback()
	// product := []entity.Product{}

	query := `SELECT * FROM products`
	var filterValues []interface{}

	// fmt.Println(*filter.Tags)

	// Conditionally append filters
	if *filter.Keyword != "" {
		query += ` WHERE name LIKE CONCAT('%', $1::TEXT, '%')`
		filterValues = append(filterValues, *filter.Keyword)
	}

	if *filter.Condition != "" {
		if len(filterValues) > 0 {
			query += ` AND `
		} else {
			query += ` WHERE `
		}
		query += ` condition = $` + strconv.Itoa(len(filterValues)+1)
		filterValues = append(filterValues, *filter.Condition)
	}

	// Add sorting if SortField and SortOrder are provided
	if *filter.SortBy != "" && *filter.OrderBy != "" {
		query += ` ORDER BY ` + *filter.SortBy + ` ` + *filter.OrderBy
	}

	// Add price range if MinPrice and MaxPrice are provided
	if *filter.MinPrice != 0 && *filter.MaxPrice != 0 {
		if len(filterValues) > 0 {
			query += ` AND `
		} else {
			query += ` WHERE `
		}
		query += `price >= $` + strconv.Itoa(len(filterValues)+1) + ` AND price <= $` + strconv.Itoa(len(filterValues)+2)
		filterValues = append(filterValues, filter.MinPrice, filter.MaxPrice)
	}

	if filter.Tags != nil && len(*filter.Tags) > 0 {
		if len(filterValues) > 0 {
			query += ` AND `
		} else {
			query += ` WHERE `
		}

		query += `tags = ANY($` + strconv.Itoa(len(filterValues)+1) + `)`
		// query += `tags && $` + strconv.Itoa(len(filterValues)+1)
		filterValues = append(filterValues, pq.Array([]string{"perabotan"}))
	}

	fmt.Println(query)

	// Execute the query
	rows, err := r.DB.Query(query, filterValues...)
	if err != nil {
		// Handle error
		return nil, err
	}

	// Slice to hold the fetched products
	var products []entity.Product

	// Loop through the rows and scan each product into the slice
	for rows.Next() {
		var product entity.Product
		// Use pq.Array to scan the Tags column into the product.Tags slice
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.ImageUrl, &product.Stock, &product.Condition, &product.IsPurchasable, pq.Array(&product.Tags), &product.UserId)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	fmt.Println(filter.Tags)
	fmt.Println(products)

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// if err := query.Find(&product).Error; err != nil {
	// 	return nil, err
	// }

	return products, nil
}

func (r *ProductRepository) Get(id string, product *entity.Product) (entity.Product, error) {
	tx, _ := r.DB.Beginx()
	defer tx.Rollback()
	query := `SELECT * FROM products WHERE id = $1`
	productData := entity.Product{}

	// Execute the query
	rows, err := r.DB.Query(query, id)
	if err != nil {
		// Handle error
		return productData, err
	}

	// Loop through the rows and scan each product into the slice
	for rows.Next() {
		var product entity.Product
		// Use pq.Array to scan the Tags column into the product.Tags slice
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.ImageUrl, &product.Stock, &product.Condition, &product.IsPurchasable, pq.Array(&product.Tags), &product.UserId)
		if err != nil {
			return productData, err
		}
		productData = product
	}

	if err != nil {
		return productData, err
	}

	return productData, nil
}

func (r *ProductRepository) Create(request *entity.Product) error {
	query := `INSERT INTO products VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.DB.Exec(query, request.ID, request.Name, request.Price, request.ImageUrl, request.Stock, request.Condition, request.IsPurchasable, pq.Array(request.Tags), "1")
	return err
}

func (r *ProductRepository) Update(id string, request *entity.Product) error {
	query := `UPDATE products SET name = $2, price = $3, imageUrl = $4, condition = $5, isPurchasable = $6,tags = $7 WHERE id = $1`

	_, err := r.DB.Exec(query, id, request.Name, request.Price, request.ImageUrl, request.Condition, request.IsPurchasable, pq.Array(request.Tags))
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) UpdateStock(id string, request *entity.Product) error {
	query := `UPDATE products SET stock = $2 WHERE id = $1`

	_, err := r.DB.Exec(query, id, request.Stock)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Delete(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	// Send query to database.
	_, err := r.DB.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
