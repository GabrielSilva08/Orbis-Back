package main

import (
	"log"

	"github.com/GabrielSilva08/Orbis/internal/controllers/tasksController"
	"github.com/GabrielSilva08/Orbis/internal/controllers/userController"
	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	"github.com/GabrielSilva08/Orbis/internal/models/userModel"
	db "github.com/GabrielSilva08/Orbis/internal/repositories"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tasksRepo"
	"github.com/GabrielSilva08/Orbis/internal/repositories/userRepo"
	"github.com/GabrielSilva08/Orbis/internal/services/tasksService"
	"github.com/GabrielSilva08/Orbis/internal/services/userService"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	v1 := app.Group("/api/v1")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	db.Connect() //se conectando com o banco de dados
	db.Database.AutoMigrate(userModel.User{})
	db.Database.AutoMigrate(tasksModel.Tag{})
	db.Database.AutoMigrate(tasksModel.Task{})

	userrepo := userRepo.NewUserRepository()
	userservice := userService.NewUserService(userrepo)
	userController.NewUserController(userservice, v1)

	taskRepo := tasksRepo.NewTaskRepository()
	taskService := tasksService.NewTaskService(taskRepo)
	tasksController.NewTaskController(taskService, v1)

	tagrepo := tasksRepo.NewTagRepository()
	tagservice := tasksService.NewTagService(tagrepo)
	tasksController.NewTagController(tagservice, v1)

	app.Listen(":3000")
}
