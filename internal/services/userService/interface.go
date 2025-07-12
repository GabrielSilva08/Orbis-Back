package userService

import "github.com/GabrielSilva08/Orbis/internal/models"

type UserServiceInterface interface{
	Create(user models.User) (models.User, error)
}