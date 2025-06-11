package tasksModel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"varchar(255);not null" json:"name"`
	Color     string    `gorm:"varchar(7)" json:"color"`
	Tasks     []Task    `gorm:"foreignKey:TagID" json:"tasks"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}

func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) { //método para gerar automaticamente o uuid antes de inserir no banco de dados
	t.ID = uuid.New()
	return
}
