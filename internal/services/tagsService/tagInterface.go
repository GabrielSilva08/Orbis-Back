package tagsService

import (
	tagdtos "github.com/GabrielSilva08/Orbis/internal/dtos/tagDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
)

type TagServiceInterface interface {
	Create(tagdtos.CreateTagDto) (models.Tag, error)
}
