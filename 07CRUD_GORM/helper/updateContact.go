package helper

import (
	"CRUD_GORM/database"
	"CRUD_GORM/model"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
)

func UpdateContact() error {

	//Input
	fmt.Print("Enter Name to search: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Wrong input: Name")
	}
	name = strings.TrimSpace(name)

	//Searching Contact to update
	searchedContacts, err := database.SearchContacts(name)
	if err != nil {

		return err
	}

	fmt.Println("Contacts found: ")
	for index, c := range searchedContacts {
		fmt.Printf("%d	Name: %s	Add: %s\n", index+1, c.Name, c.Add)
	}

	fmt.Println("Enter the Sr. No. of the contact you want to update ")
	input, _ := reader.ReadString('\n')
	sNo, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		return errors.New("Failed to convert string into int")

	}

	sNo = sNo - 1
	if sNo < 0 || sNo >= int64(len(searchedContacts)) {

		return errors.New("Sr. no. Exited")

	}

	fmt.Println("Enter Updates. \nLeave blank for no changes")
	fmt.Print("Name: ")
	name, err = reader.ReadString('\n')
	if err != nil {

		return errors.New("Wrong input: Name")

	}
	name = strings.TrimSpace(name)

	fmt.Print("Address: ")
	add, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Wrong input: Address")
	}
	add = strings.TrimSpace(add)

	fmt.Print("Number: ")
	number, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Wrong input: Number")
	}
	number = strings.TrimSpace(number)

	if number != "" && len(number) != 10 {
		return errors.New("Please Input a 10 digit valid Number")
	}

	updateReqContact := model.Contact{}
	updateReqPh := model.Ph{}

	validate := validator.New()

	if name != "" {

		updateReqContact.Name = name
		err = validate.StructPartial(updateReqContact, "Name")
		if err != nil {
			return err
		}

	}
	if add != "" {

		updateReqContact.Add = add
		err = validate.StructPartial(updateReqContact, "Add")
		if err != nil {
			return err
		}

	}

	if number != "" {

		updateReqPh.Number = number
		err = validate.StructPartial(updateReqPh, "Number")
		if err != nil {
			return err
		}
	}

	updatedContact, updatedPh := database.UpdateContact(searchedContacts[sNo], updateReqContact, updateReqPh)
	fmt.Printf(`Update Successful: 
			Name: %s
			Add: %s
			Updated Number: %s`, updatedContact.Name, updatedContact.Add, updatedPh.Number)
	fmt.Println("")
	return nil
}
