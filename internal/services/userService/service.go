package userService

import "github.com/GabrielSilva08/Orbis/internal/models"
import "github.com/GabrielSilva08/Orbis/internal/repositories/userRepo"

type UserService struct {
	repo userRepo.UserRepositoryInterface
}

func NewUserService(repo userRepo.UserRepositoryInterface) UserServiceInterface {
	return &UserService{repo: repo}
}

func (service UserService) Create(user models.User) (models.User, error) { //implementando a função definida na interface
	return service.repo.Create(user)
}