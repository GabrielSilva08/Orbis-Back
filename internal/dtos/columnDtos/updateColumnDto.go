package columndtos

import (
	"github.com/google/uuid"
)

type UpdateColumnDto struct {
	ID     uuid.UUID `json:"id" binding:"required"`
	Name   *string   `json:"name,omitempty"`
	Color  *string   `json:"color,omitempty"`
}