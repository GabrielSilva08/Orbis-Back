package columncontroller

import (
	"regexp"

	columnDtos "github.com/GabrielSilva08/Orbis/internal/dtos/columnDtos"
	columnservice "github.com/GabrielSilva08/Orbis/internal/services/columnService"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ColumnController struct {
	service columnservice.ColumnServiceInterface
}

var validate = validator.New()

func init() { //usando a biblioteca validade para cadastrar uma validação usando regex na cor
	_ = validate.RegisterValidation("hexcolor", func(fl validator.FieldLevel) bool {
		color := fl.Field().String()
		match, _ := regexp.MatchString(`^#(?:[0-9a-fA-F]{3}){1,2}$`, color)
		return match
	})
}

func (cc ColumnController) defineRoutes(router fiber.Router) {
	tagGroup := router.Group("/columns")

	tagGroup.Post("/", cc.Create)
	tagGroup.Get("/:id", cc.ListAll)
	tagGroup.Delete("/:id", cc.Delete)
	tagGroup.Patch("/", cc.Update)
}

func NewColumnController(service columnservice.ColumnServiceInterface, router fiber.Router) ColumnControllerInterface {
	var cc = ColumnController{service: service}

	cc.defineRoutes(router)

	return &cc
}

func (cc ColumnController) Create(ctx *fiber.Ctx) error {
	var colReq columnDtos.CreateColumnDto

	if err := ctx.BodyParser(&colReq); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Message:": err.Error()})
	}

	if err := validate.Struct(colReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	tag, err := cc.service.Create(colReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(tag)
}

func (cc ColumnController) Delete(ctx *fiber.Ctx) error {
	colId, idErr := uuid.Parse(ctx.Params("id"))

	if idErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": idErr.Error()})
	}

	err := cc.service.Delete(colId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}

func (cc ColumnController) ListAll(ctx *fiber.Ctx) error {
	userId, idErr := uuid.Parse(ctx.Params("id"))

	if idErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": idErr.Error()})
	}

	cols, err := cc.service.ListAll(userId)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(cols)
}

func (cc ColumnController) Update(ctx *fiber.Ctx) error {
	var colReq columnDtos.UpdateColumnDto

	if err := ctx.BodyParser(&colReq); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Message:": err.Error()})
	}

	if err := validate.Struct(colReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	col, err := cc.service.Update(colReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(col)
}