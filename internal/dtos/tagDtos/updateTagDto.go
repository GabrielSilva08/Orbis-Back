package tagdtos

import (
	"github.com/google/uuid"
)

type UpdateTagDto struct {
	Id	   uuid.UUID `json:"id" validate:"required"`
	Name   *string    `json:"name" validate:"required"`
	Color  *string    `json:"color" validate:"required,hexcolor"`
}