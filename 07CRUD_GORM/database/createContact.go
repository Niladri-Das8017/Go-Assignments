package database

import (
	"CRUD_GORM/model"

	"gorm.io/gorm"
)

func CreateContact(db *gorm.DB, newContact model.Contact) model.Contact {

	db.Create(&newContact)
	return newContact

}
