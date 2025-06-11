package tasksController

import (
	"errors"

	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/GabrielSilva08/Orbis/internal/services/tasksService"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type taskController struct {
	service tasksService.TaskServiceInterface
}

func (tc taskController) Create(ctx *fiber.Ctx) error {
	var taskReq models.Task

	if err := ctx.BodyParser(&taskReq); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Message:": err.Error()})
	}

	task, err := tc.service.Create(taskReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(task)
}

func (tc taskController) ListAllTasks(ctx *fiber.Ctx) error {
	tasks, err := tc.service.ListAllTasks()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error:": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(tasks)
}

func (tc taskController) GetTaskByID(ctx *fiber.Ctx) error {
	taskIDString := ctx.Params("id")

	taskID, err := uuid.Parse(taskIDString)
	if err != nil {
		// UUID inválida
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error:": err.Error(),
		})
	}

	task, err := tc.service.GetTaskByID(taskID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// task não encontrada
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error:": err.Error()})
		} else {
			// outro erro qualquer
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error:": err.Error()})
		}
	}

	return ctx.Status(fiber.StatusFound).JSON(task)
}

func (tc taskController) DeleteTaskByID(ctx *fiber.Ctx) error {
	taskIDString := ctx.Params("id")

	taskID, err := uuid.Parse(taskIDString)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error:": err.Error(),
		})
	}

	err = tc.service.DeleteTaskByID(taskID)
	if err != nil {
		if err.Error() == "task not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error:": "Task not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error:": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message:": "Task deleted successfully"})
}

func (tc taskController) defineRoutes(router fiber.Router) {
	taskGroup := router.Group("/tasks")

	taskGroup.Post("/", tc.Create)
	taskGroup.Get("/", tc.ListAllTasks)
	taskGroup.Get("/:id", tc.GetTaskByID)
	taskGroup.Delete("/:id", tc.DeleteTaskByID)
}

func NewTaskController(service tasksService.TaskServiceInterface, router fiber.Router) TaskControllerInterface {
	var tc = taskController{service: service}

	tc.defineRoutes(router)

	return &tc
}
