package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Profile struct
type Profile struct {
	gorm.Model
	Id          uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // primary key
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phoneNumber"`
	School      string     `json:"school"`
	Field       string     `json:"field"`
	SubFields   []SubField `gorm:"foreignKey:ProfileId"` // Adding a slice of SubField to relate multiple SubFields to a Profile

}
