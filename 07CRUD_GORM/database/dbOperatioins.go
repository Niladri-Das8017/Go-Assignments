package database

import (
	"CRUD_GORM/model"
	"errors"

	"gorm.io/gorm"
)

func CreateContact(db *gorm.DB, newContact model.Contact) model.Contact {

	db.Create(&newContact)
	return newContact

}

func ListAllContacts(db *gorm.DB) ([]model.Contact, error) {
	var contacts []model.Contact
	db.Find(&contacts)

	if len(contacts) == 0 {
		return nil, errors.New("Empty PhoneBook")
	}

	return contacts, nil
}

func SearchContacts(db *gorm.DB, name string) ([]model.Contact, error) {

	var searchedContacts []model.Contact

	db.Where(&model.Contact{Name: name}).Find(&searchedContacts)
	if len(searchedContacts) == 0 {
		return nil, errors.New("No such Contact exist!")
	}

	return searchedContacts, nil
}

func UpdateContact(db *gorm.DB, contact model.Contact, updateData model.Contact) model.Contact {

	db.Model(&contact).Updates(updateData)
	return contact
}

func DeleteContact(db *gorm.DB, contact model.Contact) gorm.DeletedAt {

	db.Delete(&contact)

	return contact.DeletedAt
}
