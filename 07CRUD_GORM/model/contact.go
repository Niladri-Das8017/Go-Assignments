package model

import (
	"gorm.io/gorm"
)

//Structure of Contact
type Contact struct {
	gorm.Model
	Name string `validate:"required"`
	Add  string `validate:"required"`
}

//Structure of Phone No.
type Ph struct {
	gorm.Model
	ContactID uint
	Number string `validate:"required,gte=10"`
	Contact   Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"required"`
}
