package tasksRepo

import (
	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	db "github.com/GabrielSilva08/Orbis/internal/repositories"
)

type taskRepository struct{}

func (tr taskRepository) Create(task tasksModel.Task) (tasksModel.Task, error) {
	result := db.Database.Create(&task)
	return task, result.Error
}

func NewTaskRepository() TaskRepositoryInterface {
	return &taskRepository{}
}

type taskListRepository struct{}

func (tlr taskListRepository) Create(taskList tasksModel.TaskList) (tasksModel.TaskList, error) {
	result := db.Database.Create(&taskList)
	return taskList, result.Error
}

func NewTaskListRepository() TaskListRepositoryInterface {
	return &taskListRepository{}
}
