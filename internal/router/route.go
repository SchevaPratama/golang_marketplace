package router

import (
	"golang-marketplace/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	ProductHandler *handler.ProductHandler
	UserHandler    *handler.UserHandler
}

func (c *RouteConfig) Setup() {

	c.App.Post("/api/user/register", c.UserHandler.Register)
	c.App.Post("api/user/login", c.UserHandler.Login)

	c.App.Get("/api/product", c.ProductHandler.List)
	c.App.Get("/api/product/:id", c.ProductHandler.Get)
	c.App.Post("/api/product", c.ProductHandler.Create)
	c.App.Delete("/api/product/:id", c.ProductHandler.Delete)
	c.App.Put("/api/product/:id", c.ProductHandler.Update)

}
