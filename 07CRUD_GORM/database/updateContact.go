package database

import (
	"CRUD_GORM/model"

	"gorm.io/gorm"
)

func UpdateContact(db *gorm.DB, contact model.Contact, name string, number string) model.Contact {

	db.Model(&contact).Updates(model.Contact{Name: name, Number: number})
	return contact
}
