package handler

import (
	"golang-marketplace/internal/model"
	"golang-marketplace/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	Service *service.ProductService
	Log     *logrus.Logger
}

func NewProductHandler(s *service.ProductService, log *logrus.Logger) *ProductHandler {
	return &ProductHandler{
		Service: s,
		Log:     log,
	}
}

func (b *ProductHandler) List(c *fiber.Ctx) error {
	keyword := c.Query("search")
	condition := c.Query("condition")
	sortBy := c.Query("sortBy")
	orderBy := c.Query("orderBy")
	maxPrice, _ := strconv.Atoi(c.Query("maxPrice"))
	minPrice, _ := strconv.Atoi(c.Query("minPrice"))

	filter := &model.ProductFilter{
		Condition: &condition,
		Keyword:   &keyword,
		SortBy:    &sortBy,
		OrderBy:   &orderBy,
		MaxPrice:  &maxPrice,
		MinPrice:  &minPrice,
	}

	if err := c.QueryParser(filter); err != nil {
		b.Log.WithError(err).Error("failed to process request")
		return fiber.ErrBadRequest
	}

	products, err := b.Service.List(c.UserContext(), filter)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    products,
	})
}

func (b *ProductHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	product, err := b.Service.Get(c.UserContext(), id.String())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
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
		b.Log.WithError(err).Error("failed to process request")
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
	id, errUuid := uuid.Parse(c.Params("id"))
	if errUuid != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   errUuid.Error(),
		})
	}

	_, err := b.Service.Get(c.UserContext(), id.String())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := b.Service.Delete(c.UserContext(), id.String()); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "success delete a product",
	})
}

func (b *ProductHandler) Update(c *fiber.Ctx) error {
	id, errUUID := uuid.Parse(c.Params("id"))
	if errUUID != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   errUUID.Error(),
		})
	}

	_, err := b.Service.Get(c.UserContext(), id.String())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	request := new(model.ProductRequest)
	if err := c.BodyParser(request); err != nil {
		b.Log.WithError(err).Error("failed to process request")
		return fiber.ErrBadRequest
	}

	if err := b.Service.Update(c.UserContext(), id.String(), request); err != nil {
		return &fiber.Error{Message: err.Error(), Code: 400}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "success Update a product",
		"data":    request,
	})
}
