package userModel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// o model define como as nossas entidades serão, o GORM faz a conversão automática disso pra uma tabela no banco de dados
type User struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `json:"name"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) { //método para gerar automaticamente o uuid antes de inserir no banco de dados
	u.ID = uuid.New()
	return
}
