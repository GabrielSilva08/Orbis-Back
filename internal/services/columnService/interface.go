package columnservice

import (
	columndtos "github.com/GabrielSilva08/Orbis/internal/dtos/columnDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/google/uuid"
)

type ColumnServiceInterface interface {
	Create(request columndtos.CreateColumnDto) (models.Column, error)
	Delete(id uuid.UUID) (error)
	ListAll(userId uuid.UUID) ([]models.Column, error)
	Update(request columndtos.UpdateColumnDto) (models.Column, error)
}