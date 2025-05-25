package tasksController

import "github.com/gofiber/fiber/v2"

type TaskControllerInterface interface {
	Create(ctx *fiber.Ctx) error
}
