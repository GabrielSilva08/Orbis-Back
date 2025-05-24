package tasksRepo

import "github.com/GabrielSilva08/Orbis/internal/models/tasksModel"

type TaskRepositoryInterface interface {
	Create(user tasksModel.Task) (tasksModel.Task, error)
}

type TaskListRepositoryInterface interface {
	Create(user tasksModel.TaskList) (tasksModel.TaskList, error)
}
