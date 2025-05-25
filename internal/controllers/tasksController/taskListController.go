package tasksController

import (
	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	"github.com/GabrielSilva08/Orbis/internal/services/tasksService"
	"github.com/gofiber/fiber/v2"
)

type taskListController struct {
	service tasksService.TaskListServiceInterface
}

func (tlc taskListController) Create(ctx *fiber.Ctx) error {
	var taskListReq tasksModel.TaskList

	if err := ctx.BodyParser(&taskListReq); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Message:": err.Error()})
	}

	task, err := tlc.service.Create(taskListReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(task)
}

func (tlc taskListController) defineRoutes(router fiber.Router) {
	taskListGroup := router.Group("/taskLists")

	taskListGroup.Post("/", tlc.Create)
}

func NewTaskListController(service tasksService.TaskListServiceInterface, router fiber.Router) TaskListControllerInterface {
	var tlc = taskListController{service: service}

	tlc.defineRoutes(router)

	return &tlc
}
