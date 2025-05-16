package userController

import (
	"github.com/GabrielSilva08/Orbis/internal/models/userModel"
	"github.com/GabrielSilva08/Orbis/internal/services/userService"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	service userService.UserServiceInterface
}

func (uc userController) Create(ctx *fiber.Ctx) error {
	var userReq userModel.User

	if err := ctx.BodyParser(&userReq); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Message": err.Error()}) //se não conseguir transformar o JSON em user, erro unprocessable entidy
	}

	user, err := uc.service.Create(userReq)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()}) //tratando de um possível erro no backend, erro bad request
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func (uc userController) defineRoutes(router fiber.Router) {
	userGroup := router.Group("/users") //definindo o grupo de rotas

	userGroup.Post("/", uc.Create)
}

func NewUserController(service userService.UserServiceInterface, router fiber.Router) UserControllerInterface {
	var uc = userController{service: service}

	uc.defineRoutes(router)

	return &uc
}
