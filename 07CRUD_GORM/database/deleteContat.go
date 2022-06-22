package database

import (
	"CRUD_GORM/model"

	"gorm.io/gorm"
)

func DeleteContact(db *gorm.DB, contact model.Contact)  gorm.DeletedAt{

	db.Delete(&contact)
	
	return contact.DeletedAt
}
