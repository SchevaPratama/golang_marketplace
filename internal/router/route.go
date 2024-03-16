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
	ImageHandler       *handler.ImageHandler
	BankAccountHandler *handler.BankAccountHandler
}

func (c *RouteConfig) Setup() {

	authMiddleware := middleware.NewAuthMiddleware("secret")

	c.App.Post("/api/user/register", c.UserHandler.Register)
	c.App.Post("api/user/login", c.UserHandler.Login)

	image := c.App.Group("/api/image", authMiddleware)
	image.Post("/", c.ImageHandler.Upload)

	product := c.App.Group("/api/product", authMiddleware)
	product.Get("", c.ProductHandler.List)
	product.Post("", c.ProductHandler.Create)
	product.Get("/:id", c.ProductHandler.Get)
	product.Delete("/:id", c.ProductHandler.Delete)
	product.Put("/:id", c.ProductHandler.Update)
	product.Post("/:id/stock", c.ProductHandler.UpdateStock)

	bankAccount := c.App.Group("/api/bank/account", authMiddleware)
	bankAccount.Get("/", c.BankAccountHandler.List)
	bankAccount.Get("/:id", c.BankAccountHandler.Get)
	bankAccount.Patch("/:id", c.BankAccountHandler.Update)
	bankAccount.Delete("/:id", c.BankAccountHandler.Delete)
	bankAccount.Post("/", c.BankAccountHandler.Create)
}
