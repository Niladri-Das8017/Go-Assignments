package database

import (
	"CRUD_GORM/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestDBOT(t *testing.T) {

	name := "testName"
	add := "test Address"
	number := "1234567890"

	//test CreateContact
	ph, err := CreateContact(model.Contact{Name: name, Add: add}, model.Ph{Number: number})
	assert.Nil(t, err)
	assert.IsType(t, model.Ph{}, ph)

	//test ListAllContacts
	allContacts, err := ListAllContacts()
	assert.Nil(t, err)
	assert.IsType(t, []model.Ph{}, allContacts)

	//test SearchContacts
	searchedContacts, err := SearchContacts(name)
	assert.Nil(t, err)
	assert.IsType(t, []model.Contact{}, searchedContacts)

	//test PrintContacts
	err = PrintContacts(searchedContacts)
	assert.Nil(t, err)

	//test UpdateContact
	updateContact := model.Contact{Name: "updatedName", Add: "updated Address"}
	updatePh := model.Ph{Number: "3216549870"}
	updatedContact, updatedPh := UpdateContact(ph.Contact, updateContact, updatePh)
	assert.IsType(t, model.Contact{}, updatedContact)
	assert.IsType(t, model.Ph{}, updatedPh)

	//test DeleteContact
	deletedAt := DeleteContact(updatedContact)
	assert.IsType(t, gorm.DeletedAt{}, deletedAt)

}
