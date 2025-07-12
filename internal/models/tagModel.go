package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	TagID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"varchar(255);not null" json:"name"`
	Color     string    `gorm:"varchar(7);not null" json:"color"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
	Tasks     []Task    `gorm:"foreignKey:TagID;constraint:OnDelete:SET NULL" json:"tasks"` // caso queira manter tasks se tag for deletada
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}


func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) { //método para gerar automaticamente o uuid antes de inserir no banco de dados
	t.TagID = uuid.New()
	return
}
