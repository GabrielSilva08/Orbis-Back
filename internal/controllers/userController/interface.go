package userController

import "github.com/gofiber/fiber/v2"

type UserControllerInterface interface {
	Create(ctx *fiber.Ctx) error
}
