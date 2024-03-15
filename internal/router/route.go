package router

import (
	"golang-marketplace/internal/handler"
	"golang-marketplace/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                *fiber.App
	ProductHandler     *handler.ProductHandler
	UserHandler        *handler.UserHandler
	BankAccountHandler *handler.BankAccountHandler
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
	product.Post("/:id/stock", c.ProductHandler.UpdateStock)

	bankAccount := c.App.Group("/api/bank/account", jwt)
	bankAccount.Get("/", c.BankAccountHandler.List)
	bankAccount.Get("/:id", c.BankAccountHandler.Get)
	bankAccount.Patch("/:id", c.BankAccountHandler.Update)
	bankAccount.Delete("/:id", c.BankAccountHandler.Delete)
	bankAccount.Post("/", c.BankAccountHandler.Create)

	//c.App.Get("/api/product", jwt, c.ProductHandler.List)
	//c.App.Get("/api/product/:id", c.ProductHandler.Get)
	//c.App.Post("/api/product", c.ProductHandler.Create)
	//c.App.Delete("/api/product/:id", c.ProductHandler.Delete)
	//c.App.Put("/api/product/:id", c.ProductHandler.Update)
}
