package tasksModel

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Priority string

const (
	PriorityLow    Priority = "Low"
	PriorityMedium Priority = "Medium"
	PriorityHigh   Priority = "High"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string    `gorm:"varchar(255);not null" json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadLine"`
	Priority    Priority  `gorm:"type:varchar(10);check:priority IN ('Low','Medium','High')" json:"priority"`
	Progress    bool      `gorm:"default:false;not null" json:"progress"`
	TagID       uuid.UUID `gorm:"type:uuid;index" json:"tagId"`
	Tag         Tag       `gorm:"foreignKey:TagID" json:"tag"`
	// UserID NOT NULL
	// ColumnID NULLABLE
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}

func (p Priority) IsValid() bool {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return true
	default:
		return false
	}
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) { //método para gerar automaticamente o uuid antes de inserir no banco de dados
	t.ID = uuid.New()
	if !t.Priority.IsValid() {
		return errors.New("invalid priority value")
	}
	return
}

func (t *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	if !t.Priority.IsValid() {
		return errors.New("invalid priority value")
	}
	return
}
