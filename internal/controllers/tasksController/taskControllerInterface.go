package tasksController

import "github.com/gofiber/fiber/v2"

type TaskControllerInterface interface {
	Create(ctx *fiber.Ctx) error
	ListAllTasks(ctx *fiber.Ctx) error
	GetTaskByID(ctx *fiber.Ctx) error
	DeleteTaskByID(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	GetTasksByTag(ctx *fiber.Ctx) error
}
