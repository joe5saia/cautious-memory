package model

import (
	"gorm.io/gorm"
)

type Subfield struct {
	gorm.Model
	Name     string
	Profiles []Profile `gorm:"many2many:profile_subfields;"`
}
