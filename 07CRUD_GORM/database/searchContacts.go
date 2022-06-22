package database

import (
	"CRUD_GORM/model"
	"errors"

	"gorm.io/gorm"
)

func SearchContacts(db *gorm.DB, name string) ([]model.Contact, error) {

	var searchedContacts []model.Contact

	db.Where(&model.Contact{Name: name}).Find(&searchedContacts)
	if len(searchedContacts) == 0 {
		return nil, errors.New("No such Contact exist!")
	}

	return searchedContacts, nil
}
