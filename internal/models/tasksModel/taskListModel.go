package tasksModel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskList struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string    `gorm:"varchar(255);not null" json:"title"`
	Tasks     []Task    `gorm:"foreignKey:TaskListID" json:"tasks"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}

func (u *TaskList) BeforeCreate(tx *gorm.DB) (err error) { //método para gerar automaticamente o uuid antes de inserir no banco de dados
	u.ID = uuid.New()
	return
}
