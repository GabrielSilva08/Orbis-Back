package tasksController

import (
	"errors"
	"fmt"

	taskdtos "github.com/GabrielSilva08/Orbis/internal/dtos/taskDtos"
	"github.com/GabrielSilva08/Orbis/internal/services/tasksService"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type taskController struct {
	service tasksService.TaskServiceInterface
}

var validate = validator.New()

func (tc taskController) Create(ctx *fiber.Ctx) error {
	var taskReq taskdtos.CreateTaskDto

	if err := ctx.BodyParser(&taskReq); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Message:": err.Error()})
	}

	if err := validate.Struct(taskReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
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

func (tc taskController) GetTasksByTag(ctx *fiber.Ctx) error {
	tagIDString := ctx.Params("id")

	tagID, err := uuid.Parse(tagIDString)
	if err != nil {
		// UUID inválida
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error:": err.Error(),
		})
	}

	task, err := tc.service.GetTaskByID(tagID)

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

func (tc taskController) Update(ctx *fiber.Ctx) error {
	taskIDString := ctx.Params("id")

	taskID, err := uuid.Parse(taskIDString)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error:": err.Error(),
		})
	}
	var taskReq taskdtos.UpdateTaskDto

	fmt.Printf("Body recebido: %s\n", string(ctx.Body()))

	if err := ctx.BodyParser(&taskReq); err != nil {
		fmt.Printf("Erro no BodyParser: %v\n", err)
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "Invalid JSON format",
			"details": err.Error(),
		})
	}

	fmt.Print(taskReq)

	if err := validate.Struct(taskReq); err != nil {
		fmt.Printf("Erro na validação: %v\n", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}
	task, err := tc.service.Update(taskID, taskReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}

func (tc taskController) defineRoutes(router fiber.Router) {
	taskGroup := router.Group("/tasks")

	taskGroup.Post("/", tc.Create)
	taskGroup.Get("/", tc.ListAllTasks)
	taskGroup.Get("/:id", tc.GetTaskByID)
	taskGroup.Get("/tag/:id", tc.GetTasksByTag)
	taskGroup.Delete("/:id", tc.DeleteTaskByID)
	taskGroup.Patch("/:id", tc.Update)
}

func NewTaskController(service tasksService.TaskServiceInterface, router fiber.Router) TaskControllerInterface {
	var tc = taskController{service: service}

	tc.defineRoutes(router)

	return &tc
}
