package tagsRepo

import (
	"github.com/GabrielSilva08/Orbis/internal/models"
	db "github.com/GabrielSilva08/Orbis/internal/repositories"
)

type tagRepository struct{}

func (tr tagRepository) Create(tag models.Tag) (models.Tag, error) {
	result := db.Database.Create(&tag)
	return tag, result.Error
}

func NewTagRepository() TagRepositoryInterface {
	return &tagRepository{}
}