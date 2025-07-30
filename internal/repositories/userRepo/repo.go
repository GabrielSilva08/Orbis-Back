package userRepo

import (
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/GabrielSilva08/Orbis/internal/repositories"
)

type userRepository struct {} //definindo a struct do repositório que irá implementar os métodos da interface

func (ur userRepository) Create(user models.User) (models.User, error)  { //implementando a função definida na interface
	result := db.Database.Create(&user)
	return user, result.Error
}

func NewUserRepository() UserRepositoryInterface { //construtor da struct, será chamado na main
	return &userRepository{}
}