package router

import (
	"golang-marketplace/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	ProductHandler *handler.ProductHandler
}

func (c *RouteConfig) Setup() {
	c.App.Get("/api/products", c.ProductHandler.List)
	c.App.Get("/api/products/:id", c.ProductHandler.Get)
	c.App.Post("/api/products", c.ProductHandler.Create)
	c.App.Delete("/api/products/:id", c.ProductHandler.Delete)
	c.App.Put("/api/products/:id", c.ProductHandler.Update)
}
