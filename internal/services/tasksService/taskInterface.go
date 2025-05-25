package tasksService

import "github.com/GabrielSilva08/Orbis/internal/models/tasksModel"

type TaskServiceInterface interface {
	Create(task tasksModel.Task) (tasksModel.Task, error)
}
