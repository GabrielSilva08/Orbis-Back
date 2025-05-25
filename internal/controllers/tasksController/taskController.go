package tasksController

import (
	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	"github.com/GabrielSilva08/Orbis/internal/services/tasksService"
	"github.com/gofiber/fiber/v2"
)

type taskController struct {
	service tasksService.TaskServiceInterface
}

func (tc taskController) Create(ctx *fiber.Ctx) error {
	var taskReq tasksModel.Task

	if err := ctx.BodyParser(&taskReq); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Message:": err.Error()})
	}

	task, err := tc.service.Create(taskReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(task)
}

func (tc taskController) defineRoutes(router fiber.Router) {
	taskGroup := router.Group("/tasks")

	taskGroup.Post("/", tc.Create)
}

func NewTaskController(service tasksService.TaskServiceInterface, router fiber.Router) TaskControllerInterface {
	var tc = taskController{service: service}

	tc.defineRoutes(router)

	return &tc
}
