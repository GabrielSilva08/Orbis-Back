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
	result := db.Database.Find(&tags)
	return tags, result.Error
}

func (tr tagRepository) Delete(id uuid.UUID) error {
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

	updateData := make(map[string]interface{})
	if request.Name != nil {
		updateData["Name"] = *request.Name
	}
	if request.Color != nil {
		updateData["Color"] = *request.Color
	}

	if err := db.Database.Model(&tag).Updates(updateData).Error; err != nil {
		return tag, err
	}

	// 4. Retorna a task atualizada
	if err := db.Database.First(&tag, "tag_id = ?", request.Id).Error; err != nil {
		return tag, err
	}

	return tag, nil
}

func NewTagRepository() TagRepositoryInterface {
	return &tagRepository{}
}
