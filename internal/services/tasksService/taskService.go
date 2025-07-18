package tasksService

import (
	"fmt"
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

	layouts := []string{
		time.RFC3339,
		"2006-01-02",
		"2006-01-02 15:04:05",
		"02/01/2006 15:04",
	}
	
	var deadline time.Time
	var err error
	for _, layout := range layouts {
		deadline, err = time.Parse(layout, request.DeadLine)
		if err == nil {
			break
		}
	}

	fmt.Print(deadline)

	if err != nil {
		return taskdtos.CreateTaskResponse{}, err
	}

	task := models.Task{
		Title:       request.Title,
		Description: request.Description,
		Deadline:    deadline,
		Priority:    request.Priority,
		Progress: 	 request.Progress,
		UserID: 	request.User,
		TagID: 		nil,
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
		TagID:       nil,
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

func (service TaskService) GetTasksByTag(tagId uuid.UUID) ([]models.Task, error) {
	return service.repo.GetTasksByTag(tagId)
}

func (service TaskService) DeleteTaskByID(id uuid.UUID) error {
	return service.repo.DeleteTaskByID(id)
}

func (service TaskService) Update(id uuid.UUID, request taskdtos.UpdateTaskDto) (models.Task, error) {

	return service.repo.Update(id, request)
}
