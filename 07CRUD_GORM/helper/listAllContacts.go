package helper

import (
	"CRUD_GORM/database"
	"fmt"
)

func ListAllContacts() error {

	listOfPhNo, err := database.ListAllContacts()
	if err != nil {

		return err

	}

	//Print Result
	for index, ph := range listOfPhNo {
		fmt.Printf("%d	Name: %s	Add: %s		Number: %s\n", index+1, ph.Contact.Name, ph.Contact.Add, ph.Number)
	}

	return nil
}
