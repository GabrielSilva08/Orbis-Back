package tasksRepo

import (
	"errors"

	"github.com/GabrielSilva08/Orbis/internal/models"
	db "github.com/GabrielSilva08/Orbis/internal/repositories"
	taskdtos "github.com/GabrielSilva08/Orbis/internal/dtos/taskDtos"
	"github.com/google/uuid"
)

type taskRepository struct{}

func (tr taskRepository) Create(task models.Task) (models.Task, error) {
	// Criar a task
	if err := db.Database.Create(&task).Error; err != nil {
		return models.Task{}, err
	}

	// Recarregar com relacionamentos
	if err := db.Database.Preload("Tag").First(&task, task.TaskID).Error; err != nil {
		return task, err // Retorna a task criada mesmo se não conseguir carregar a tag
	}

	return task, nil
}

func (tr taskRepository) ListAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := db.Database.Preload("Tag").Find(&tasks) // Preload para carregar as tags dentro de tasks
	return tasks, result.Error
}

func (tr taskRepository) GetTaskByID(id uuid.UUID) (models.Task, error) {
	var task models.Task
	result := db.Database.Preload("Tag").First(&task, id)
	return task, result.Error
}

func (tr taskRepository) DeleteTaskByID(id uuid.UUID) error {
	result := db.Database.Delete(&models.Task{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (tr taskRepository) Update(request taskdtos.UpdateTaskDto) (models.Task, error) {
	var task models.Task
	readResult := db.Database.First(&task, "task_id = ?", request.TaskID)

	if readResult.Error != nil {
		return task, readResult.Error
	}
	
	updateResult := db.Database.Model(&task).
	Select("Title", "Description", "Deadline", "Priority", "Progress", "TagID", "ColumnID").
	Updates(models.Task{
		Title:       *request.Title,
		Description: *request.Description,
		Deadline:    request.Deadline,
		Priority:    request.Priority,
		Progress:    *request.Progress,
		TagID:       request.TagID,
		ColumnID:    request.ColumnID,
	})

	db.Database.First(&task, "task_id = ?", request.TaskID) //buscando de novo para retornar a task atualizada
	
	return task, updateResult.Error
}

func NewTaskRepository() TaskRepositoryInterface {
	return &taskRepository{}
}