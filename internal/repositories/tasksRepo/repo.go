package tasksRepo

import (
	"errors"

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

	return task, nil
}

func (tr taskRepository) ListAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := db.Database.Find(&tasks)
	return tasks, result.Error
}

func (tr taskRepository) GetTaskByID(id uuid.UUID) (models.Task, error) {
	var task models.Task
	result := db.Database.First(&task, id)
	return task, result.Error
}

func (tr taskRepository) GetTasksByTag(tagId uuid.UUID) ([]models.Task, error) {
	var task []models.Task
	result := db.Database.Preload("Tag").Where("tag_id = ?", tagId).Find(&task)
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

	// 1. Busca a tarefa existente no banco de dados
	if err := db.Database.First(&task, "task_id = ?", id).Error; err != nil {
		return task, err
	}

	// 2. Prepara os dados para atualização de forma segura (somente campos preenchidos)
	updateData := make(map[string]interface{})

	if request.Title != nil {
		updateData["Title"] = *request.Title
	}
	if request.Description != nil {
		updateData["Description"] = *request.Description
	}
	if request.Progress != nil {
		updateData["Progress"] = *request.Progress
	}
	if request.Deadline != nil {
		updateData["Deadline"] = *request.Deadline
	}
	if request.Priority != nil {
		updateData["Priority"] = *request.Priority
	}
	if request.TagID != nil {
		updateData["TagId"] = *request.TagID
	}
	if request.ColumnID != nil {
		updateData["ColumnId"] = *request.ColumnID
	}
	if request.UserID != nil {
		updateData["UserId"] = *request.UserID
	}

	// 3. Atualiza apenas os campos presentes
	if err := db.Database.Model(&task).Updates(updateData).Error; err != nil {
		return task, err
	}

	// 4. Retorna a task atualizada
	if err := db.Database.First(&task, "task_id = ?", id).Error; err != nil {
		return task, err
	}

	return task, nil
}

func NewTaskRepository() TaskRepositoryInterface {
	return &taskRepository{}
}
