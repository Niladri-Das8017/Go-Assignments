package model

import "gorm.io/gorm"


//Structure of Contact
type Contact struct {
	gorm.Model
	Name string
	Number string
}
