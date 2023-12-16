package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	School      string
	Field       string
	Subfields   []Subfield `gorm:"many2many:profile_subfields;"`
}
