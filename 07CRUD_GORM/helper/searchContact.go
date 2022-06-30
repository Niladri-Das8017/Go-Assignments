package helper

import (
	"CRUD_GORM/database"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func SearchContact() error {

	//Input
	fmt.Print("Enter Name to search: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Wrong input: Name")
	}
	name = strings.TrimSpace(name)

	searchedContacts, err := database.SearchContacts(name)
	if err != nil {
		return err
	}

	fmt.Println("Contacts found: ")
	err = database.PrintContacts(searchedContacts)
	if err != nil {
		return err
	}

	return nil
}
