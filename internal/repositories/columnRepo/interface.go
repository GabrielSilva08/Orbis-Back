package columnRepo

import (
	columnDtos "github.com/GabrielSilva08/Orbis/internal/dtos/columnDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/google/uuid"
)

type ColumnRepoInterface interface {
	Create(task models.Column) (models.Column, error)  // Cria coluna passando um objeto coluna
	ListAllColumns(userId uuid.UUID) ([]models.Column, error)          // Lista todas as colunas por usuário
	DeleteColumnByID(id uuid.UUID) error             // Deleta uma coluna pelo ID
	Update(request columnDtos.UpdateColumnDto) (models.Column, error)
}