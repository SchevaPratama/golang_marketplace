package router

import (
	"golang-marketplace/internal/handler"
	"golang-marketplace/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	ProductHandler *handler.ProductHandler
	UserHandler    *handler.UserHandler
}

func (c *RouteConfig) Setup() {

	jwt := middleware.NewAuthMiddleware("secret")

	c.App.Post("/api/user/register", c.UserHandler.Register)
	c.App.Post("api/user/login", c.UserHandler.Login)

	product := c.App.Group("/api/product", jwt)
	product.Get("/", c.ProductHandler.List)
	product.Get("/:id", c.ProductHandler.Get)
	product.Post("/", c.ProductHandler.Create)
	product.Delete("/:id", c.ProductHandler.Delete)
	product.Put("/:id", c.ProductHandler.Update)

	//c.App.Get("/api/product", jwt, c.ProductHandler.List)
	//c.App.Get("/api/product/:id", c.ProductHandler.Get)
	//c.App.Post("/api/product", c.ProductHandler.Create)
	//c.App.Delete("/api/product/:id", c.ProductHandler.Delete)
	//c.App.Put("/api/product/:id", c.ProductHandler.Update)

}
