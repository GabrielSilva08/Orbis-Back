package tasksService

import "github.com/GabrielSilva08/Orbis/internal/models/tasksModel"

type TaskListServiceInterface interface {
	Create(taskList tasksModel.TaskList) (tasksModel.TaskList, error)
}
