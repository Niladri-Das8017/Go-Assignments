package database

import (
	"CRUD_GORM/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestDBInteractors(t *testing.T) {

	//test IniiitDB
	db, err := InitDB()
	assert.Nil(t, err)
	assert.IsType(t, &gorm.DB{}, db)

	name := "testName"
	number := "1234567890"

	//test CreateContact
	contact := CreateContact(db, name, number)
	assert.IsType(t, model.Contact{}, contact)

	//test ListAllContacts
	allContacts, err := ListAllContacts(db)
	assert.Nil(t, err)
	assert.IsType(t, []model.Contact{}, allContacts)

	//test SearchContacts
	searchedContacts, err := SearchContacts(db, name)
	assert.Nil(t, err)
	assert.IsType(t, []model.Contact{}, searchedContacts)

	//test UpdateContact
	updatedContact := UpdateContact(db, contact, "updatedName", "9874563210")
	assert.IsType(t, model.Contact{}, updatedContact)

	//test DeleteContact
	deletedAt := DeleteContact(db, updatedContact)
	assert.IsType(t, gorm.DeletedAt{}, deletedAt)

}
