package userRepo

import "github.com/GabrielSilva08/Orbis/internal/models"

type UserRepositoryInterface interface { //essa é a interface do repositório, contém os métodos que serão implementados
	Create(user models.User) (models.User, error)
}