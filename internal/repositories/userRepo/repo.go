package userRepo

import (
	"fmt"

	"github.com/GabrielSilva08/Orbis/internal/models/userModel"
)

type userRepository struct {} //definindo a struct do repositório que irá implementar os métodos da interface

func (ur userRepository) Create(user userModel.User) error  { //implementando a função definida na interface
	fmt.Println("Usuário criado")
	fmt.Println(user)
	return nil
}

func NewUserRepository() UserRepositoryInterface { //construtor da struct, será chamado na main
	return &userRepository{}
}