package tasksRepo

import (
	"errors"
	"time"

	taskdtos "github.com/GabrielSilva08/Orbis/internal/dtos/taskDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	db "github.com/GabrielSilva08/Orbis/internal/repositories"
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

func (tr taskRepository) Update(id uuid.UUID, request taskdtos.UpdateTaskDto) (models.Task, error) {
	var task models.Task

	readResult := db.Database.First(&task, "task_id = ?", id)

	if readResult.Error != nil {
		return task, readResult.Error
	}

	updates := make(map[string]interface{})

	if request.Title != nil {
		updates["title"] = *request.Title
	}

	if request.Description != nil {
		updates["description"] = *request.Description
	}

	if request.Deadline != nil {
		parsedDeadline, err := time.Parse(time.RFC3339, *request.Deadline)
		if err != nil {
			return models.Task{}, err
		}
		updates["deadline"] = parsedDeadline
	}

	if request.Priority != nil {
		updates["priority"] = *request.Priority
	}

	if request.Progress != nil {
		updates["progress"] = *request.Progress
	}

	if request.TagID != nil {
		updates["tag_id"] = *request.TagID
	}

	if request.ColumnID != nil {
		updates["column_id"] = *request.ColumnID
	}

	updateResult := db.Database.Model(&task).Updates(updates)
	if updateResult.Error != nil {
		return task, updateResult.Error
	}

	db.Database.First(&task, "task_id = ?", id) //buscando de novo para retornar a task atualizada

	return task, updateResult.Error
}

func NewTaskRepository() TaskRepositoryInterface {
	return &taskRepository{}
}
