package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Beer represents a beer record.
type Beer struct {
	gorm.Model
	Name    string `json:"name"`
	Content string `json:"content"`
	Pic     string `json:"pic"`
}

// Validate validates the Beer fields.
func (m Beer) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
		validation.Field(&m.Content, validation.Required),
	)
}
