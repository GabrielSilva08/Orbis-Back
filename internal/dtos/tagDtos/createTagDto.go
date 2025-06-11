package tagdtos

import(
	"github.com/google/uuid"
)

type CreateTagDto struct {
	Name   string    `json:"name" validate:"required"`
	Color  string    `json:"color" validate:"required,hexcolor"`
	UserID uuid.UUID `json:"userId" validate:"required"`
}