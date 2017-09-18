package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
	"github.com/lucasb-eyer/go-colorful"
)

type BeerStyle struct {
	gorm.Model
	Srm   int            `json:"srm"`
	Name  string         `json:"name"`
	color colorful.Color `json:"color"`
	Ebc   int            `json:"ebc"`
}

// Validate validates the Beer fields.
func (m BeerStyle) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 100)),
	)
}
