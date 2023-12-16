package model

import (
	"gorm.io/gorm"
)

type Subfield struct {
	gorm.Model
	Name     string    `gorm:"unique;not null"`
	Profiles []Profile `gorm:"many2many:profile_subfields;"`
}
