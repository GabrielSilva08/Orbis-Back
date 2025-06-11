package tagsRepo

import (
	"github.com/GabrielSilva08/Orbis/internal/models"
)

type TagRepositoryInterface interface {
	Create(tag models.Tag) (models.Tag, error) // Cria uma tag
}
