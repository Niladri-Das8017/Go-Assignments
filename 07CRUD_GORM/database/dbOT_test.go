package database

import (
	"CRUD_GORM/model"
	"fmt"
	"reflect"
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
	assert.EqualValues(t, ph.Number, number)

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
	updatedContact, updatedPh, err := UpdateContact(ph.Contact, updateContact, updatePh)
	assert.Nil(t, err)
	assert.IsType(t, model.Contact{}, updatedContact)
	assert.IsType(t, model.Ph{}, updatedPh)
	assert.EqualValues(t, updateContact.Name, updateContact.Name)
	assert.EqualValues(t, updatedContact.Add, updateContact.Add)
	assert.EqualValues(t, updatedPh.Number, updatePh.Number)

	equalValuesExcept(t, updatePh, updatedPh, "gorm.Model")

	//test DeleteContact
	deletedContact, err := DeleteContact(updatedContact)
	assert.Nil(t, err)
	assert.IsType(t, gorm.DeletedAt{}, deletedContact.DeletedAt)

}

func equalValuesExcept(t *testing.T, expected interface{}, actual interface{}, field string) bool {

	exp := reflect.ValueOf(expected)
	ac := reflect.ValueOf(actual)

	for i := 0; i < exp.NumField(); i++ {

		if exp.Type().Field(i).Name == field {

			continue

		}

		if !assert.EqualValues(t, exp.Field(i), ac.Field(i)) {

			fmt.Println("field: ", exp.Type().Field(i).Name)
			return false

		}

	}

	return true
}
