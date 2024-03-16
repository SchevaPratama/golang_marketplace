package service

import (
	"context"
	"fmt"

	"golang-marketplace/internal/entity"
	helpers "golang-marketplace/internal/helper"
	"golang-marketplace/internal/model"
	"golang-marketplace/internal/model/converter"
	"golang-marketplace/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ProductService struct {
	Repository   *repository.ProductRepository
	ImageService *ImageService
	Validate     *validator.Validate
	Log          *logrus.Logger
}

func NewProductService(r *repository.ProductRepository, i *ImageService, validate *validator.Validate, log *logrus.Logger) *ProductService {
	return &ProductService{Repository: r, ImageService: i, Validate: validate, Log: log}
}

func (s *ProductService) List(ctx context.Context, filter *model.ProductFilter) ([]model.ProductRespone, error) {

	if err := helpers.ValidationError(s.Validate, filter); err != nil {
		s.Log.WithError(err).Error("failed to validate request query params")
		return nil, err
	}

	products, err := s.Repository.List(filter)
	if err != nil {
		s.Log.WithError(err).Error("failed get product lists")
		return nil, err
	}

	newProducts := make([]model.ProductRespone, len(products))
	for i, product := range products {
		newProducts[i] = *converter.ProductConverter(&product)
	}

	return newProducts, nil
}

func (s *ProductService) Get(ctx context.Context, id string) (*model.ProductRespone, error) {
	product := new(entity.Product)

	productData, err := s.Repository.Get(id, product)
	fmt.Println(err)
	if err != nil {
		s.Log.WithError(err).Error("failed get product detail")
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

	newRequest := &entity.Product{
		ID:            uuid.New().String(),
		Name:          request.Name,
		Price:         request.Price,
		ImageUrl:      request.ImageUrl,
		Stock:         request.Stock,
		Condition:     request.Condition,
		IsPurchasable: request.IsPurchasable,
		Tags:          request.Tags,
	}

	err := s.Repository.Create(newRequest)
	if err != nil {
		//s.Log.Error("failed to insert new data")
		return err
	}

	return nil
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	// product := new(entity.Product)
	// productData, err := s.Repository.Get(id, product)
	// if err != nil {
	// 	s.Log.WithError(err).Error("failed get product detail")
	// 	return err
	// }

	errDel := s.Repository.Delete(id)
	if errDel != nil {
		s.Log.WithError(errDel).Error("failed to delete data")
		return errDel
	}

	return nil
}

func (s *ProductService) Update(ctx context.Context, id string, request *model.ProductRequest) error {
	// if err := s.Validate.Struct(request); err != nil {
	if err := helpers.ValidationError(s.Validate, request); err != nil {
		s.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	product := new(entity.Product)
	_, err := s.Repository.Get(id, product)
	if err != nil {
		s.Log.WithError(err).Error("failed get product detail")
		return err
	}

	product.Name = request.Name
	product.Price = request.Price
	product.ImageUrl = request.ImageUrl
	product.Condition = request.Condition
	product.IsPurchasable = request.IsPurchasable
	product.Tags = request.Tags

	err = s.Repository.Update(id, product)
	if err != nil {
		s.Log.WithError(err).Error("failed to update data")
		return err
	}

	return nil
}
