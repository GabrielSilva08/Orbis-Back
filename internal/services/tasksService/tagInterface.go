package tasksService

import "github.com/GabrielSilva08/Orbis/internal/models/tasksModel"

type TagServiceInterface interface {
	Create(tag tasksModel.Tag) (tasksModel.Tag, error)
}
