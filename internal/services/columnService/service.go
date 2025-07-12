package columnservice

import (
	columndtos "github.com/GabrielSilva08/Orbis/internal/dtos/columnDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/GabrielSilva08/Orbis/internal/repositories/columnRepo"
	"github.com/google/uuid"
)

type ColumnService struct {
	repo columnRepo.ColumnRepoInterface
}

func NewColumnService(repo columnRepo.ColumnRepoInterface) ColumnServiceInterface {
	return &ColumnService{repo: repo}
}

func (cs ColumnService) Create(request columndtos.CreateColumnDto) (models.Column, error) {
	column := models.Column{
		Name: request.Name,
		Color: request.Color,
		UserID: request.UserID,
	}

	return cs.repo.Create(column)
}

func (cs ColumnService) Delete(id uuid.UUID) (error) {
	return cs.repo.DeleteColumnByID(id)
}

func (cs ColumnService) Update(request columndtos.UpdateColumnDto) (models.Column, error){
	return cs.repo.Update(request)
}

func (cs ColumnService) ListAll(userID uuid.UUID) ([]models.Column, error) {
	return cs.repo.ListAllColumns(userID)
}