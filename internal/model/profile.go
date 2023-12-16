package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model  `json:"-"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	PhoneNumber string
	School      string
	Fields      []Field    `gorm:"many2many:profile_fields;"`
	Subfields   []Subfield `gorm:"many2many:profile_subfields;"`
}
