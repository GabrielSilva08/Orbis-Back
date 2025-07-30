package columnRepo

import (
	"errors"

	columnDtos "github.com/GabrielSilva08/Orbis/internal/dtos/columnDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	db "github.com/GabrielSilva08/Orbis/internal/repositories"
	"github.com/google/uuid"
)

type ColumnRepository struct {}

func NewColumnRepository() ColumnRepoInterface {
	return &ColumnRepository{}
}

func (ct ColumnRepository) Create(column models.Column) (models.Column, error) {
	if err := db.Database.Create(&column).Error; err != nil {
		return models.Column{}, err
	}

	return column, nil
}

func (ct ColumnRepository) DeleteColumnByID(id uuid.UUID) error {
	result := db.Database.Delete(&models.Column{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("column not found")
	}

	return nil
}

func (ct ColumnRepository) Update(request columnDtos.UpdateColumnDto) (models.Column, error) {
	var column models.Column

	// 1. Busca a coluna existente no banco de dados
	if err := db.Database.First(&column, "column_id = ?", request.ID).Error; err != nil {
		return column, err
	}

	// 2. Prepara os dados para atualização de forma segura (somente campos preenchidos)
	updateData := make(map[string]interface{})

	if request.Name != nil {
		updateData["Name"] = *request.Name
	}
	if request.Color != nil {
		updateData["Color"] = *request.Color
	}

	// 3. Atualiza apenas os campos presentes
	if err := db.Database.Model(&column).Updates(updateData).Error; err != nil {
		return column, err
	}

	// 4. Retorna a task atualizada
	if err := db.Database.First(&column, "column_id = ?", request.ID).Error; err != nil {
		return column, err
	}

	return column, nil
}

func (ct ColumnRepository) ListAllColumns(userID uuid.UUID) ([]models.Column, error) {
	var columns []models.Column
	err := db.Database.
		Where("user_id = ?", userID).
		Preload("Tasks").
		Find(&columns).Error

	return columns, err
}