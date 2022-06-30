package helper

import (
	"CRUD_GORM/database"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DeleteContacts() error {

	//Input
	fmt.Print("Enter Name to delete: ")
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

	fmt.Println("Enter the Sr. No. of the contact you want to delete ")
	input, _ := reader.ReadString('\n')
	sNo, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		return errors.New("Failed to convert string into int")

	}

	sNo = sNo - 1
	if sNo < 0 || sNo >= int64(len(searchedContacts)) {

		return errors.New("Sr. no. Exited")

	}

	deletedAt := database.DeleteContact(searchedContacts[sNo])

	fmt.Println("Contact Deletd\n Deleted at : ", deletedAt)

	return nil
}
