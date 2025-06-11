package tasksController

import "github.com/gofiber/fiber/v2"

type TagControllerInterface interface {
	Create(ctx *fiber.Ctx) error
}
