package tagsRepo

import (
	"errors"

	tagdtos "github.com/GabrielSilva08/Orbis/internal/dtos/tagDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	db "github.com/GabrielSilva08/Orbis/internal/repositories"
	"github.com/google/uuid"
)

type tagRepository struct{}

func (tr tagRepository) Create(tag models.Tag) (models.Tag, error) {
	result := db.Database.Create(&tag)
	return tag, result.Error
}

func (tr tagRepository) ListAll() ([]models.Tag, error) {
	var tags []models.Tag
	result := db.Database.Preload("Tasks").Preload("User").Find(&tags)
	return tags, result.Error
}

func (tr tagRepository) Delete(id uuid.UUID) (error) {
	result := db.Database.Delete(&models.Tag{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("tag não encontrada")
	}

	return nil
}

func (tr tagRepository) Update(request tagdtos.UpdateTagDto) (models.Tag, error) {
	var tag models.Tag
	readResult := db.Database.First(&tag, "tag_id = ?", request.Id)

	if readResult.Error != nil {
		return tag, readResult.Error
	}
	
	updateResult := db.Database.Model(&tag).Select("Name","Color").Updates(models.Tag{Name: request.Name, Color: request.Color})

	db.Database.First(&tag, "task_id = ?", request.Id) //buscando de novo para retornar a tag atualizada
	
	return tag, updateResult.Error
}

func NewTagRepository() TagRepositoryInterface {
	return &tagRepository{}
}