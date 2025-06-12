package tagsRepo

import (
	tagdtos "github.com/GabrielSilva08/Orbis/internal/dtos/tagDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/google/uuid"
)

type TagRepositoryInterface interface {
	Create(tag models.Tag) (models.Tag, error)
	ListAll() ([]models.Tag, error)
	Delete(id uuid.UUID) error
	Update(request tagdtos.UpdateTagDto) (models.Tag, error)
}
