package tasksService

import (
	"time"

	taskdtos "github.com/GabrielSilva08/Orbis/internal/dtos/taskDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tasksRepo"
	"github.com/google/uuid"
)

type TaskService struct {
	repo tasksRepo.TaskRepositoryInterface
}

func NewTaskService(repo tasksRepo.TaskRepositoryInterface) TaskServiceInterface {
	return &TaskService{repo: repo}
}

func (service TaskService) Create(request taskdtos.CreateTaskDto) (taskdtos.CreateTaskResponse, error) {

	deadline, err := time.Parse(time.RFC3339, request.DeadLine)

	if err != nil {
		return taskdtos.CreateTaskResponse{}, err
	}

	task := models.Task{
		Title:       request.Title,
		Description: request.Description,
		Deadline:    deadline,
		Priority:    request.Priority,
		Progress:    request.Progress,
	}

	createdTask, err := service.repo.Create(task)
	if err != nil {
		return taskdtos.CreateTaskResponse{}, err
	}

	response := taskdtos.CreateTaskResponse{
		ID:          createdTask.TaskID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		DeadLine:    createdTask.Deadline,
		Priority:    string(createdTask.Priority),
		Progress:    createdTask.Progress,
		TagID:       createdTask.TagID,
		ColumnID:    createdTask.ColumnID,
		CreatedAt:   createdTask.CreatedAt,
		UpdatedAt:   createdTask.UpdatedAt,
	}

	return response, err
}

func (service TaskService) ListAllTasks() ([]models.Task, error) {
	return service.repo.ListAllTasks()
}

func (service TaskService) GetTaskByID(id uuid.UUID) (models.Task, error) {
	return service.repo.GetTaskByID(id)
}

func (service TaskService) DeleteTaskByID(id uuid.UUID) error {
	return service.repo.DeleteTaskByID(id)
}

func (service TaskService) Update(id uuid.UUID, request taskdtos.UpdateTaskDto) (models.Task, error) {

	return service.repo.Update(id, request)
}
