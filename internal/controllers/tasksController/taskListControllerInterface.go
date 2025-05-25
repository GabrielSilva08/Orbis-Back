package tasksController

import "github.com/gofiber/fiber/v2"

type TaskListControllerInterface interface {
	Create(ctx *fiber.Ctx) error
}
