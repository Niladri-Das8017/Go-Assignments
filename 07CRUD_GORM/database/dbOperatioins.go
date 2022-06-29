package database

import (
	"CRUD_GORM/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateContact(db *gorm.DB, newContact model.Contact, newPh model.Ph) model.Ph {

	db.Create(&newContact)

	newPh.Contact = newContact
	db.Create(&newPh)

	return newPh

}

func ListAllContacts(db *gorm.DB) ([]model.Ph, error) {
	var listOfPhNo []model.Ph

	//Query using Join
	db.Joins("Contact").Find(&listOfPhNo)

	if len(listOfPhNo) == 0 {
		return nil, errors.New("Empty PhoneBook")
	}

	return listOfPhNo, nil
}

func SearchContacts(db *gorm.DB, name string) ([]model.Contact, error) {

	var searchedContacts []model.Contact

	db.Where(&model.Contact{Name: name}).Find(&searchedContacts)
	if len(searchedContacts) == 0 {
		return nil, errors.New("No such Contact exist!")
	}

	return searchedContacts, nil
}

func PrintContacts(db *gorm.DB, Contacts []model.Contact) error {

	if len(Contacts) == 0 {
		return errors.New("No contact to show")
	}

	var ph model.Ph

	for index, c := range Contacts {

		//Query using where
		db.Where(&model.Ph{Contact: c}).Find(&ph)

		fmt.Printf("%d	Name: %s	Add: %s		Number: %s\n", index+1, c.Name, c.Add, ph.Number)

	}
	return nil
}

func UpdateContact(db *gorm.DB, contact model.Contact, updateReqContact model.Contact, updateReqPh model.Ph) (model.Contact, model.Ph) {

	db.Model(&contact).Updates(updateReqContact)

	var ph model.Ph
	db.Where(&model.Ph{Contact: contact}).Find(&ph)
	db.Model(&ph).Updates(updateReqPh)

	return contact, ph
}

func DeleteContact(db *gorm.DB, contact model.Contact) gorm.DeletedAt {

	db.Delete(&contact)

	return contact.DeletedAt
}
