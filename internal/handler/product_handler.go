package handler

import (
	"fmt"
	"golang-marketplace/internal/model"
	"golang-marketplace/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sagikazarmark/slog-shim"
)

type ProductHandler struct {
	Service *service.ProductService
	Log     *slog.Logger
}

func NewProductHandler(s *service.ProductService, log *slog.Logger) *ProductHandler {
	return &ProductHandler{
		Service: s,
		Log:     log,
	}
}

func (b *ProductHandler) List(c *fiber.Ctx) error {
	keyword := c.Query("keyword")

	filter := &model.ProductFilter{
		Keyword: &keyword,
	}

	if err := c.QueryParser(filter); err != nil {
		b.Log.Error("failed to process request")
		return fiber.ErrBadRequest
	}

	products, err := b.Service.List(c.UserContext(), filter)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "list of products",
		"data":    products,
	})
}

func (b *ProductHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		b.Log.Error("failed parse param id")
		return fiber.ErrBadRequest
	}

	product, err := b.Service.Get(c.UserContext(), id)
	if err != nil {
		return fiber.ErrNotFound
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "detail of product",
		"data":    product,
	})
}

func (b *ProductHandler) Create(c *fiber.Ctx) error {
	request := new(model.ProductRequest)

	if err := c.BodyParser(request); err != nil {
		b.Log.Error("failed to process request")
		return fiber.ErrBadRequest
		// return &fiber.Error{Message: "Opppss", Code: 400}
	}

	err := b.Service.Create(c.UserContext(), request)
	if err != nil {
		// return fiber.ErrBadRequest
		return &fiber.Error{Message: err.Error(), Code: 400}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "success insert new product",
		"data":    request,
	})
}

func (b *ProductHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	if id == "" {
		b.Log.Error("failed parse param id")
		return fiber.ErrBadRequest
	}

	_, err := b.Service.Get(c.UserContext(), id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if err := b.Service.Delete(c.UserContext(), id); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "success delete a product",
		"data":    nil,
	})
}

func (b *ProductHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	if id == "" {
		b.Log.Error("failed parse param id")
		return fiber.ErrBadRequest
	}

	request := new(model.ProductRequest)
	if err := c.BodyParser(request); err != nil {
		b.Log.Error("failed to process request")
		return fiber.ErrBadRequest
	}

	if err := b.Service.Update(c.UserContext(), id, request); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "success Update a product",
		"data":    request,
	})
}
