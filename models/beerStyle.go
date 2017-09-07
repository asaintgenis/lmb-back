package models

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/go-ozzo/ozzo-validation"
)

type BeerStyle struct {
	Id   int    `json:"id" db:"id"`
	Srm int `json:"srm" db:"srm"`
	Name string `json:"name" db:"name"`
	color colorful.Color `json:"color" db:"color"`
	Ebc int `json:"ebc" db:"ebc"`
}

// Validate validates the Beer fields.
func (m BeerStyle) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 100)),
	)
}