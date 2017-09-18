package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Beer represents a beer record.
type Bar struct {
	gorm.Model
	Name string `json:"name"`
}

// Validate validates the Beer fields.
func (m Bar) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}
