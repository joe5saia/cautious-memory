package schemas

import "github.com/google/uuid"

/*
profile struct for the api representing the full profile json object
*/

// Profile struct
type Profile struct {
	Id          uuid.UUID  `json:"id"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phoneNumber"`
	School      string     `json:"school"`
	Field       string     `json:"field"`
	SubFields   []SubField `json:"subFields"`
}
