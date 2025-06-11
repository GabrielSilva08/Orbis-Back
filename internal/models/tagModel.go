package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	TagID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"varchar(255);not null" json:"name"`
	Color     string    `gorm:"varchar(7);not null" json:"color"`

	Task      []Task    `json:"task"` // Relação: Uma tag para várias tarefas
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"userId"` // chave estrangeira
	User   User      	`json:"user"`    // relação com User

	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}

func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) { //método para gerar automaticamente o uuid antes de inserir no banco de dados
	t.TagID = uuid.New()
	return
}
