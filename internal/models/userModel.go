package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// o model define como as nossas entidades serão, o GORM faz a conversão automática disso pra uma tabela no banco de dados
type User struct {
	UserID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `json:"name"`
	Tag  []Tag     `json:"tag"` //relação 1 to many de tag com user
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) { //método para gerar automaticamente o uuid antes de inserir no banco de dados
	u.UserID = uuid.New()
	return
}
