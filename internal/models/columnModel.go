package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Column struct {
	ColumnID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `gorm:"not null" json:"name"`
	Color    string    `gorm:"varchar(7);not null" json:"color"`
	UserID   uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
	Tasks    []Task    `gorm:"foreignKey:ColumnId;constraint:OnDelete:SET NULL" json:"tasks"`
}

func (c *Column) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ColumnID == uuid.Nil {
		c.ColumnID = uuid.New()
	}
	return
}
