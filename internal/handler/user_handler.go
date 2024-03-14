package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sagikazarmark/slog-shim"
	"golang-marketplace/internal/model"
	"golang-marketplace/internal/service"
)

type UserHandler struct {
	Service *service.UserService
	Log     *slog.Logger
}

func NewUserHandler(s *service.UserService, log *slog.Logger) *UserHandler {
	return &UserHandler{
		Service: s,
		Log:     log,
	}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	request := new(model.RegisterRequest)

	if err := c.BodyParser(request); err != nil {
		h.Log.Error("failed to process request")
		return fiber.ErrBadRequest
	}

	resp, err := h.Service.Register(c.UserContext(), request)
	if err != nil {
		h.Log.Error("failed to process request")
		return fiber.ErrBadRequest
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    resp,
	})
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	request := new(model.RegisterRequest)

	if err := c.BodyParser(request); err != nil {
		h.Log.Error("failed to process request")
		return fiber.ErrBadRequest
	}

	resp, err := h.Service.Register(c.UserContext(), request)
	if err != nil {
		h.Log.Error("failed to process request")
		return fiber.ErrBadRequest
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    resp,
	})
}
