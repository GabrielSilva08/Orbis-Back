package tasksRepo

import (
	"errors"

	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	db "github.com/GabrielSilva08/Orbis/internal/repositories"
	"github.com/google/uuid"
)

type taskRepository struct{}

func (tr taskRepository) Create(task tasksModel.Task) (tasksModel.Task, error) {
	// Criar a task
	if err := db.Database.Create(&task).Error; err != nil {
		return tasksModel.Task{}, err
	}

	// Recarregar com relacionamentos
	if err := db.Database.Preload("Tag").First(&task, task.ID).Error; err != nil {
		return task, err // Retorna a task criada mesmo se não conseguir carregar a tag
	}

	return task, nil
}

func (tr taskRepository) ListAllTasks() ([]tasksModel.Task, error) {
	var tasks []tasksModel.Task
	result := db.Database.Preload("Tag").Find(&tasks) // Preload para carregar as tags dentro de tasks
	return tasks, result.Error
}

func (tr taskRepository) GetTaskByID(id uuid.UUID) (tasksModel.Task, error) {
	var task tasksModel.Task
	result := db.Database.Preload("Tag").First(&task, id)
	return task, result.Error
}

func (tr taskRepository) DeleteTaskByID(id uuid.UUID) error {
	result := db.Database.Delete(&tasksModel.Task{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

func NewTaskRepository() TaskRepositoryInterface {
	return &taskRepository{}
}

type tagRepository struct{}

func (tr tagRepository) Create(tag tasksModel.Tag) (tasksModel.Tag, error) {
	result := db.Database.Create(&tag)
	return tag, result.Error
}

func NewTagRepository() TagRepositoryInterface {
	return &tagRepository{}
}
