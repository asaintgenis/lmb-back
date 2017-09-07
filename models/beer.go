package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Beer represents a beer record.
type Beer struct {
	gorm.Model
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Content string `json:"content" db:"content"`
	Pic string `json:"pic" db:"pic"`
	beerStyle BeerStyle `json:"beerStyle_id" db:"beerstyle_id"`
}

// Validate validates the Beer fields.
func (m Beer) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
		validation.Field(&m.Content,validation.Required),
	)
}