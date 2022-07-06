package database

import (
	"CRUD_GORM/model"
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func CreateContact(newContact model.Contact, newPh model.Ph) (model.Ph, error) {

	err := validate.Struct(newContact)
	if err != nil {
		return model.Ph{}, err
	}

	err = DB.Create(&newContact).Error
	if err != nil {
		return model.Ph{}, err
	}

	newPh.Contact = newContact
	err = validate.Struct(newPh)
	if err != nil {
		return model.Ph{}, err
	}

	err = DB.Create(&newPh).Error
	if err != nil {
		return model.Ph{}, err
	}

	return newPh, nil

}

func ListAllContacts() ([]model.Ph, error) {
	var listOfPhNo []model.Ph

	//Query using Join
	err := DB.Joins("Contact").Find(&listOfPhNo).Error
	if err != nil {
		return nil, err
	}

	if len(listOfPhNo) == 0 {
		return nil, errors.New("Empty PhoneBook")
	}

	return listOfPhNo, nil
}

func SearchContacts(name string) ([]model.Contact, error) {

	var searchedContacts []model.Contact

	err := DB.Where(&model.Contact{Name: name}).Find(&searchedContacts).Error
	if err != nil {
		return nil, err
	}

	if len(searchedContacts) == 0 {
		return nil, errors.New("No such Contact exist!")
	}

	return searchedContacts, nil
}

func PrintContacts(Contacts []model.Contact) error {

	if len(Contacts) == 0 {
		return errors.New("No contact to show")
	}

	var ph model.Ph

	for index, c := range Contacts {

		//Query using where
		err := DB.Where(&model.Ph{Contact: c}).Find(&ph).Error
		if err != nil {
			return err
		}

		fmt.Printf("%d	Name: %s	Add: %s		Number: %s\n", index+1, c.Name, c.Add, ph.Number)

	}
	return nil
}

func UpdateContact(contact model.Contact, updateReqContact model.Contact, updateReqPh model.Ph) (model.Contact, model.Ph, error) {

	err := DB.Model(&contact).Updates(updateReqContact).Error
	if err != nil {
		return model.Contact{}, model.Ph{}, err
	}

	var ph model.Ph
	err = DB.Where(&model.Ph{Contact: contact}).Find(&ph).Error
	if err != nil {
		return model.Contact{}, model.Ph{}, err
	}

	err = DB.Model(&ph).Updates(updateReqPh).Error
	if err != nil {
		return model.Contact{}, model.Ph{}, err
	}

	return contact, ph, nil
}

func DeleteContact(contact model.Contact) (*model.Contact, error) {

	err := DB.Delete(&contact).Error
	if err != nil {
		return nil, err
	}

	return &contact, nil
}
