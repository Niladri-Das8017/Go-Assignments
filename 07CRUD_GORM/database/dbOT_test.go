package database

import (
	"CRUD_GORM/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestDBOT(t *testing.T) {

	//test IniiitDB
	db, err := InitDB("PhoneBook.db")
	assert.Nil(t, err)
	assert.IsType(t, &gorm.DB{}, db)

	name := "testName"
	add := "test Address"
	number := "1234567890"

	//test CreateContact
	ph := CreateContact(db, model.Contact{Name: name, Add: add}, model.Ph{Number: number})
	assert.IsType(t, model.Ph{}, ph)

	//test ListAllContacts
	allContacts, err := ListAllContacts(db)
	assert.Nil(t, err)
	assert.IsType(t, []model.Ph{}, allContacts)

	//test SearchContacts
	searchedContacts, err := SearchContacts(db, name)
	assert.Nil(t, err)
	assert.IsType(t, []model.Contact{}, searchedContacts)

	//test PrintContacts
	err = PrintContacts(db, searchedContacts)
	assert.Nil(t, err)

	//test UpdateContact
	updateContact := model.Contact{Name:  "updatedName", Add: "updated Address"}
	updatePh := model.Ph{Number: "3216549870"}
	updatedContact, updatedPh := UpdateContact(db, ph.Contact, updateContact, updatePh )
	assert.IsType(t, model.Contact{}, updatedContact)
	assert.IsType(t, model.Ph{}, updatedPh)

	//test DeleteContact
	deletedAt := DeleteContact(db, updatedContact)
	assert.IsType(t, gorm.DeletedAt{}, deletedAt)

}
