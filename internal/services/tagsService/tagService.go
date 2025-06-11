package tagsService

import (
	tagdtos "github.com/GabrielSilva08/Orbis/internal/dtos/tagDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tagsRepo"
)

type TagService struct {
	repo tagsRepo.TagRepositoryInterface
}

func NewTagService(repo tagsRepo.TagRepositoryInterface) TagServiceInterface {
	return &TagService{repo: repo}
}

func (service TagService) Create(request tagdtos.CreateTagDto) (models.Tag, error) {

	tag := models.Tag{ //instanciando o model a partir da request
		Name:   request.Name,
		Color:  request.Color,
		UserID: request.UserID,
	}

	return service.repo.Create(tag)
}
