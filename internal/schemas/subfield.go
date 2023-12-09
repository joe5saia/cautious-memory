package schemas

import (
	"github.com/google/uuid"
)

type SubField struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
