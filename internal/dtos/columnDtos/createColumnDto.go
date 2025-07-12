package columndtos

import (
	"github.com/google/uuid"
)

type CreateColumnDto struct {
	Name   string    `json:"name" binding:"required"`
	Color  string    `json:"color" binding:"required,hexcolor"`
	UserID uuid.UUID `json:"userId" binding:"required"`
}