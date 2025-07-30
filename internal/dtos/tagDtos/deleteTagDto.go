package tagdtos

import(
	"github.com/google/uuid"
)

type DeleteTagDto struct {
	Id	   uuid.UUID `json:"id" validate:"required"`
}