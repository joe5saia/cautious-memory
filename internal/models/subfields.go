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
	ProfileId uuid.UUID `gorm:"type:uuid"`                                      // Foreign key referencing Profile
	Profile   Profile   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Adding a reference to Profile
}
