package models

import "github.com/go-ozzo/ozzo-validation"

// Beer represents a beer record.
type Bar struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"title"`
}

// Validate validates the Beer fields.
func (m Bar) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}

