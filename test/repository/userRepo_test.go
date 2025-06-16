package repository_test

import (
	"testing"

	"github.com/GabrielSilva08/Orbis/internal/models/userModel"
	"github.com/GabrielSilva08/Orbis/internal/repositories"
	"github.com/GabrielSilva08/Orbis/internal/repositories/userRepo"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T)  {
	envErr := godotenv.Load("../../.env")
	assert.NoError(t, envErr, "Erro ao carregar o arquivo env")

	db.Connect()
	repo := userRepo.NewUserRepository()

	mockUser := userModel.User{ //o ID é criado apenas quando o usuário entra no banco de dados, portanto não deve ser comparado
		Name: "Teste",
	}

	createdUser, err := repo.Create(mockUser)

	assert.NoError(t, err, "Erro ao criar o user mockado")
	assert.Equal(t, mockUser.Name, createdUser.Name, "Usuário mandado e retornado não batem")
}