package main

import (
	"github.com/GabrielSilva08/Orbis/internal/controllers/userController"
	"github.com/GabrielSilva08/Orbis/internal/repositories/userRepo"
	"github.com/GabrielSilva08/Orbis/internal/services/userService"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	v1 := app.Group("/api/v1")

	userrepo := userRepo.NewUserRepository()
	userservice := userService.NewUserService(userrepo)

	userController.NewUserController(userservice, v1)

	app.Listen(":3000")
}
