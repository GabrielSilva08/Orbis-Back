package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// o model define como as nossas entidades serão, o GORM faz a conversão automática disso pra uma tabela no banco de dados
type User struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Tags      []Tag     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"tags"`
	Tasks     []Task    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"tasks"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) { //método para gerar automaticamente o uuid antes de inserir no banco de dados
	u.UserID = uuid.New()
	return
}
