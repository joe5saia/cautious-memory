package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SubField struct
type SubField struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // primary key
	Name      string    `json:"name"`
	ProfileId uuid.UUID `json:"profileId" gorm:"type:uuid;foreignKey:Id;references:ProfileId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
