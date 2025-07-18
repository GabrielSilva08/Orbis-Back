package main

import (
	"flag"
	"log"
	"time"

	columncontroller "github.com/GabrielSilva08/Orbis/internal/controllers/columnController"
	"github.com/GabrielSilva08/Orbis/internal/controllers/tagsController"
	"github.com/GabrielSilva08/Orbis/internal/controllers/tasksController"
	"github.com/GabrielSilva08/Orbis/internal/controllers/userController"
	columndtos "github.com/GabrielSilva08/Orbis/internal/dtos/columnDtos"
	tagdtos "github.com/GabrielSilva08/Orbis/internal/dtos/tagDtos"
	taskdtos "github.com/GabrielSilva08/Orbis/internal/dtos/taskDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	db "github.com/GabrielSilva08/Orbis/internal/repositories"
	"github.com/GabrielSilva08/Orbis/internal/repositories/columnRepo"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tagsRepo"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tasksRepo"
	"github.com/GabrielSilva08/Orbis/internal/repositories/userRepo"
	columnservice "github.com/GabrielSilva08/Orbis/internal/services/columnService"
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

	seedFlag := flag.Bool("seed", false, "Criar dados iniciais")
	flag.Parse()

	db.Connect() //se conectando com o banco de dados
	db.Database.AutoMigrate(&models.User{})
	db.Database.AutoMigrate(&models.Tag{})
	db.Database.AutoMigrate(&models.Column{})
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

	columnrepo := columnRepo.NewColumnRepository()
	columnservice := columnservice.NewColumnService(columnrepo)
	columncontroller.NewColumnController(columnservice, v1)

	if *seedFlag {
		// Criando um usuário após configurar o serviço
		newUser := models.User{
			UserID: "67594f2b-5fff-419d-894d-f1486ba4aee1",
			Name: "Marco Túlio",
		}

		createdUser, err := userservice.Create(newUser)
		if err != nil {
			log.Printf("Erro ao criar usuário: %v", err)
		} else {
			log.Printf("Usuário criado com sucesso: %+v", createdUser)
		}

		userId := createdUser.UserID

		// Criando tag
		tag := tagdtos.CreateTagDto{
			Name:   "UFC",
			Color:  "#00B7EB",
			UserID: userId,
		}

		createdTag, err := tagservice.Create(tag)
		if err != nil {
			log.Printf("Erro ao criar tag: %v", err)
		} else {
			log.Printf("Tag criada com sucesso: %+v", createdTag)
		}

		// Criando coluna
		column := columndtos.CreateColumnDto{
			Name:   "UFC",
			Color:  "#00B7EB",
			UserID: userId,
		}

		createdColumn, err := columnservice.Create(column)
		if err != nil {
			log.Printf("Erro ao criar column: %v", err)
		} else {
			log.Printf("Column criada com sucesso: %+v", createdColumn)
		}

		// Criando tasks
		task1 := taskdtos.CreateTaskDto{
			Title:       "Implementar autenticação JWT",
			Description: "Desenvolver sistema de autenticação usando JWT tokens para proteger as rotas da API",
			DeadLine:    time.Now().Add(3 * 24 * time.Hour).Format(time.RFC3339),
			Priority:    models.PriorityHigh,
			Progress:    true,
			User:        userId,
			Tag:         &createdTag.TagID,
			Column:      &createdColumn.ColumnID,
		}

		task2 := taskdtos.CreateTaskDto{
			Title:       "Criar documentação da API",
			Description: "Escrever documentação completa dos endpoints da API usando OpenAPI/Swagger",
			DeadLine:    time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339),
			Priority:    models.PriorityMedium,
			Progress:    false,
			User:        userId,
			Tag:         &createdTag.TagID,
			Column:      &createdColumn.ColumnID,
		}

		createdTask1, err := taskService.Create(task1)
		if err != nil {
			log.Printf("Erro ao criar task 1: %v", err)
		} else {
			log.Printf("Task 1 criada com sucesso: %+v", createdTask1)
		}

		createdTask2, err := taskService.Create(task2)
		if err != nil {
			log.Printf("Erro ao criar task 2: %v", err)
		} else {
			log.Printf("Task 2 criada com sucesso: %+v", createdTask2)
		}
	}

	app.Listen("0.0.0.0:3000")
}
