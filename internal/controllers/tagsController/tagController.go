package tagsController

import (
	"regexp"

	tagdtos "github.com/GabrielSilva08/Orbis/internal/dtos/tagDtos"
	"github.com/GabrielSilva08/Orbis/internal/services/tagsService"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type tagController struct {
	service tagsService.TagServiceInterface
}

var validate = validator.New()

func init() { //usando a biblioteca validade para cadastrar uma validação usando regex na cor
	_ = validate.RegisterValidation("hexcolor", func(fl validator.FieldLevel) bool {
		color := fl.Field().String()
		match, _ := regexp.MatchString(`^#(?:[0-9a-fA-F]{3}){1,2}$`, color)
		return match
	})
}

func (tc tagController) Create(ctx *fiber.Ctx) error {
	var tagReq tagdtos.CreateTagDto

	if err := ctx.BodyParser(&tagReq); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Message:": err.Error()})
	}

	if err := validate.Struct(tagReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	tag, err := tc.service.Create(tagReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(tag)
}

func (tc tagController) ListAll(ctx *fiber.Ctx) error {
	tags, err := tc.service.ListAll()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(tags)
}

func (tc tagController) Delete(ctx *fiber.Ctx) error {
	tagId, idErr := uuid.Parse(ctx.Params("id"))

	if idErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": idErr.Error()})
	}

	var tagReq = tagdtos.DeleteTagDto{Id: tagId} //encapsulando o id em um DTO

	if err := validate.Struct(tagReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	err := tc.service.Delete(tagReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}

func (tc tagController) Update(ctx *fiber.Ctx) error {
	var tagReq tagdtos.UpdateTagDto

	if err := ctx.BodyParser(&tagReq); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Message:": err.Error()})
	}

	if err := validate.Struct(tagReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	tag, err := tc.service.Update(tagReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(tag)
}

func (tc tagController) defineRoutes(router fiber.Router) {
	tagGroup := router.Group("/tags")

	tagGroup.Post("/", tc.Create)
	tagGroup.Get("/", tc.ListAll)
	tagGroup.Delete("/:id", tc.Delete)
	tagGroup.Put("/", tc.Update)
}

func NewTagController(service tagsService.TagServiceInterface, router fiber.Router) TagControllerInterface {
	var tc = tagController{service: service}

	tc.defineRoutes(router)

	return &tc
}
