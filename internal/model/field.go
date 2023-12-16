package model

import (
	"gorm.io/gorm"
)

type Field struct {
	gorm.Model `json:"-"`
	Name       string    `gorm:"unique;not null" json:"name"`
	Profiles   []Profile `gorm:"many2many:profile_fields;" json:"-"`
}
