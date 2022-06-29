package model

import "gorm.io/gorm"

//Structure of Contact
type Contact struct {
	gorm.Model
	Name string
	Add  string
}

//Structure of Phone No.
type Ph struct {
	gorm.Model
	ContactID uint
	Number string
	Contact   Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
