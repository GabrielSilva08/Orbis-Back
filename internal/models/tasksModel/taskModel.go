package tasksModel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string    `gorm:"varchar(255);not null" json:"title"`
	SubTitle    string    `gorm:"varchar(255)" json:"subTitle"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadLine"`
	Done        bool      `gorm:"default:false;not null" json:"done"`
	TaskListID  uuid.UUID `gorm:"type:uuid;index" json:"taskListId"`
	CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt"`
}

func (u *Task) BeforeCreate(tx *gorm.DB) (err error) { //método para gerar automaticamente o uuid antes de inserir no banco de dados
	u.ID = uuid.New()
	return
}
