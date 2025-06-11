package tagsService

import (
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tagsRepo"
)

type TagService struct {
	repo tagsRepo.TagRepositoryInterface
}

func NewTagService(repo tagsRepo.TagRepositoryInterface) TagServiceInterface {
	return &TagService{repo: repo}
}

func (service TagService) Create(task models.Tag) (models.Tag, error) {
	return service.repo.Create(task)
}
