package database

import (
	"CRUD_GORM/model"
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)
var validate = validator.New()

func CreateContact(newContact model.Contact, newPh model.Ph) (model.Ph, error) {

	
	err := validate.Struct(newContact)
	if err != nil {
		return model.Ph{}, err
	}

	DB.Create(&newContact)

	newPh.Contact = newContact
	err = validate.Struct(newPh)
	if err != nil {
		return model.Ph{}, err
	}

	DB.Create(&newPh)

	return newPh, nil

}

func ListAllContacts() ([]model.Ph, error) {
	var listOfPhNo []model.Ph

	//Query using Join
	DB.Joins("Contact").Find(&listOfPhNo)

	if len(listOfPhNo) == 0 {
		return nil, errors.New("Empty PhoneBook")
	}

	return listOfPhNo, nil
}

func SearchContacts(name string) ([]model.Contact, error) {

	var searchedContacts []model.Contact

	DB.Where(&model.Contact{Name: name}).Find(&searchedContacts)
	if len(searchedContacts) == 0 {
		return nil, errors.New("No such Contact exist!")
	}

	return searchedContacts, nil
}

func PrintContacts( Contacts []model.Contact) error {

	if len(Contacts) == 0 {
		return errors.New("No contact to show")
	}

	var ph model.Ph

	for index, c := range Contacts {

		//Query using where
		DB.Where(&model.Ph{Contact: c}).Find(&ph)

		fmt.Printf("%d	Name: %s	Add: %s		Number: %s\n", index+1, c.Name, c.Add, ph.Number)

	}
	return nil
}

func UpdateContact( contact model.Contact, updateReqContact model.Contact, updateReqPh model.Ph) (model.Contact, model.Ph) {

	DB.Model(&contact).Updates(updateReqContact)

	var ph model.Ph
	DB.Where(&model.Ph{Contact: contact}).Find(&ph)
	DB.Model(&ph).Updates(updateReqPh)

	return contact, ph
}

func DeleteContact(contact model.Contact) gorm.DeletedAt {

	DB.Delete(&contact)

	return contact.DeletedAt
}
