package tasksController

import (
	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	"github.com/GabrielSilva08/Orbis/internal/services/tasksService"
	"github.com/gofiber/fiber/v2"
)

type tagController struct {
	service tasksService.TagServiceInterface
}

func (tc tagController) Create(ctx *fiber.Ctx) error {
	var tagReq tasksModel.Tag

	if err := ctx.BodyParser(&tagReq); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Message:": err.Error()})
	}

	task, err := tc.service.Create(tagReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(task)
}

func (tc tagController) defineRoutes(router fiber.Router) {
	tagGroup := router.Group("/tags")

	tagGroup.Post("/", tc.Create)
}

func NewTagController(service tasksService.TagServiceInterface, router fiber.Router) TagControllerInterface {
	var tc = tagController{service: service}

	tc.defineRoutes(router)

	return &tc
}
