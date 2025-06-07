package tasksService

import (
	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tasksRepo"
)

type TagService struct {
	repo tasksRepo.TagRepositoryInterface
}

func NewTagService(repo tasksRepo.TagRepositoryInterface) TagServiceInterface {
	return &TagService{repo: repo}
}

func (service TagService) Create(task tasksModel.Tag) (tasksModel.Tag, error) {
	return service.repo.Create(task)
}
