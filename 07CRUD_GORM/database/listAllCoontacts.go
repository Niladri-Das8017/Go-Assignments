package database

import (
	"CRUD_GORM/model"
	"errors"

	"gorm.io/gorm"
)

func ListAllContacts(db *gorm.DB) ([]model.Contact, error) {
	var contacts []model.Contact
	db.Find(&contacts)

	if len(contacts) == 0 {
		return nil, errors.New("Empty PhoneBook")
	}

	return contacts, nil
}
