package database

import (
	"CRUD_GORM/model"

	"gorm.io/gorm"
)

func CreateContact(db *gorm.DB, name string, number string) model.Contact{

	newContact := model.Contact{
		Name:   name,
		Number: number,
	}

	db.Create(&newContact)
	return newContact
}
