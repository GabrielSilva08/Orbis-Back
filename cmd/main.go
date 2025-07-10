package main

import (
	"log"

	"github.com/GabrielSilva08/Orbis/internal/controllers/tagsController"
	"github.com/GabrielSilva08/Orbis/internal/controllers/tasksController"
	"github.com/GabrielSilva08/Orbis/internal/controllers/userController"
	"github.com/GabrielSilva08/Orbis/internal/models"
	db "github.com/GabrielSilva08/Orbis/internal/repositories"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tagsRepo"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tasksRepo"
	"github.com/GabrielSilva08/Orbis/internal/repositories/userRepo"
	"github.com/GabrielSilva08/Orbis/internal/services/tagsService"
	"github.com/GabrielSilva08/Orbis/internal/services/tasksService"
	"github.com/GabrielSilva08/Orbis/internal/services/userService"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	v1 := app.Group("/api/v1")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	db.Connect() //se conectando com o banco de dados
	db.Database.AutoMigrate(&models.User{})
	db.Database.AutoMigrate(&models.Tag{})
	db.Database.AutoMigrate(&models.Task{})

	userrepo := userRepo.NewUserRepository()
	userservice := userService.NewUserService(userrepo)
	userController.NewUserController(userservice, v1)

	tagrepo := tagsRepo.NewTagRepository()
	tagservice := tagsService.NewTagService(tagrepo)
	tagsController.NewTagController(tagservice, v1)

	taskRepo := tasksRepo.NewTaskRepository()
	taskService := tasksService.NewTaskService(taskRepo)
	tasksController.NewTaskController(taskService, v1)

	app.Listen("0.0.0.0:3000")
}
