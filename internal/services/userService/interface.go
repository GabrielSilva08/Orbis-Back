package userService

import "github.com/GabrielSilva08/Orbis/internal/models/userModel"

type UserServiceInterface interface{
	Create(user userModel.User) (userModel.User, error)
}