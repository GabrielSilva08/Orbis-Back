package tagsService

import "github.com/GabrielSilva08/Orbis/internal/models"

type TagServiceInterface interface {
	Create(tag models.Tag) (models.Tag, error)
}
