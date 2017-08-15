package models

import "github.com/go-ozzo/ozzo-validation"

// Beer represents a beer record.
type Beer struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Validate validates the Beer fields.
func (m Beer) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}
