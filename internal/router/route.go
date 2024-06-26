package router

import (
	"github.com/gofiber/fiber/v2"
	"golang-marketplace/internal/handler"
	"golang-marketplace/internal/middleware"
	"net/http"
)

type RouteConfig struct {
	App                *fiber.App
	ProductHandler     *handler.ProductHandler
	UserHandler        *handler.UserHandler
	ImageHandler       *handler.ImageHandler
	BankAccountHandler *handler.BankAccountHandler
}

func (c *RouteConfig) Setup() {

	authMiddleware := middleware.NewAuthMiddleware()

	c.App.Post("/v1/user/register", c.UserHandler.Register)
	c.App.Post("/v1/user/login", c.UserHandler.Login)

	image := c.App.Group("/v1/image", authMiddleware)
	image.Post("/", c.ImageHandler.Upload)

	product := c.App.Group("/v1/product", authMiddleware)
	product.Get("", c.ProductHandler.List)
	product.Post("", c.ProductHandler.Create)
	product.Get("/:id", c.ProductHandler.Get)
	product.Delete("/:id", c.ProductHandler.Delete)
	product.Put("/:id", c.ProductHandler.Update)
	product.Post("/:id/stock", c.ProductHandler.UpdateStock)
	product.Post("/:id/buy", c.ProductHandler.Buy)

	c.App.Patch("/v1/bank/account", authMiddleware, func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})
	bankAccount := c.App.Group("/v1/bank/account", authMiddleware)
	bankAccount.Get("/", c.BankAccountHandler.List)
	bankAccount.Get("/:id", c.BankAccountHandler.Get)
	bankAccount.Patch("/:id", c.BankAccountHandler.Update)
	bankAccount.Delete("/:id", c.BankAccountHandler.Delete)
	bankAccount.Post("/", c.BankAccountHandler.Create)

}
