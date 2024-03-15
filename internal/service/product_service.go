package service

import (
	"context"

	"golang-marketplace/internal/entity"
	helpers "golang-marketplace/internal/helper"
	"golang-marketplace/internal/model"
	"golang-marketplace/internal/model/converter"
	"golang-marketplace/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/sagikazarmark/slog-shim"
)

type ProductService struct {
	Repository *repository.ProductRepository
	Validate   *validator.Validate
	Log        *slog.Logger
	DB         *sqlx.DB
}

func NewProductService(r *repository.ProductRepository, validate *validator.Validate, log *slog.Logger, db *sqlx.DB) *ProductService {
	return &ProductService{Repository: r, Validate: validate, Log: log, DB: db}
}

func (s *ProductService) List(ctx context.Context, filter *model.ProductFilter) ([]model.ProductRespone, error) {
	tx, _ := s.DB.Beginx()
	defer tx.Rollback()

	if err := helpers.ValidationError(s.Validate, filter); err != nil {
		s.Log.Error("failed to validate request query params")
		return nil, err
	}

	products, err := s.Repository.List(s.DB, filter)
	if err != nil {
		s.Log.Error("failed get product lists")
		return nil, err
	}

	newProducts := make([]model.ProductRespone, len(products))
	for i, product := range products {
		newProducts[i] = *converter.ProductConverter(&product)
	}

	return newProducts, nil
}

func (s *ProductService) Get(ctx context.Context, id string) (*model.ProductRespone, error) {
	tx, _ := s.DB.Beginx()
	defer tx.Rollback()

	product := new(entity.Product)

	productData, err := s.Repository.Get(s.DB, id, product)
	if err != nil {
		s.Log.Error("failed get product detail")
		return nil, err
	}

	return converter.ProductConverter(&productData), nil
}

func (s *ProductService) Create(ctx context.Context, request *model.ProductRequest) error {
	// if err := s.Validate.Struct(request); err != nil {
	if err := helpers.ValidationError(s.Validate, request); err != nil {
		//s.Log.Error("failed to validate request body")
		return err
	}

	tx, _ := s.DB.Beginx()
	defer tx.Rollback()

	newRequest := &entity.Product{
		Name:          request.Name,
		Price:         request.Price,
		ImageUrl:      request.ImageUrl,
		Stock:         request.Stock,
		Condition:     request.Condition,
		IsPurchasable: request.IsPurchasable,
	}

	err := s.Repository.Create(s.DB, newRequest)
	if err != nil {
		//s.Log.Error("failed to insert new data")
		return err
	}

	return nil
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	tx, _ := s.DB.Beginx()
	defer tx.Rollback()

	product := new(entity.Product)
	_, err := s.Repository.Get(s.DB, id, product)
	if err != nil {
		s.Log.Error("failed get product detail")
		return err
	}

	err = s.Repository.Delete(s.DB, product)
	if err != nil {
		s.Log.Error("failed to delete data")
		return err
	}

	return nil
}

func (s *ProductService) Update(ctx context.Context, id string, request *model.ProductRequest) error {
	// if err := s.Validate.Struct(request); err != nil {
	if err := helpers.ValidationError(s.Validate, request); err != nil {
		s.Log.Error("failed to validate request body")
		return err
	}

	tx, _ := s.DB.Beginx()
	defer tx.Rollback()

	product := new(entity.Product)
	_, err := s.Repository.Get(s.DB, id, product)
	if err != nil {
		s.Log.Error("failed get product detail")
		return err
	}

	product.Name = request.Name
	product.Price = request.Price
	product.ImageUrl = request.ImageUrl
	product.Condition = request.Condition
	product.IsPurchasable = request.IsPurchasable

	err = s.Repository.Update(s.DB, id, product)
	if err != nil {
		s.Log.Error("failed to update data")
		return err
	}

	return nil
}
