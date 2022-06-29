package main

import (
	"CRUD_GORM/database"
	"CRUD_GORM/helper"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Initialize Database
	dbPath := "database/PhoneBook.db"
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	//Input Option
	for {
		fmt.Println(`Phonebook...
 
	**MENU**
	1. Create Contact
	2. List All Contacts
	3. Search a Contact
	4. Update Contact
	5. Delete Contact
	Choose other number to exit!`)

		fmt.Printf("Choose Option: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		option, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
		if err != nil {
			log.Fatal("Failed to convert string into int")
		}

		switch option {

		//Create Contact
		case 1:

			err := helper.CreateContact(db)
			if err != nil {
				fmt.Println(err)
			}

			continue

		//List All Contacts
		case 2:

			err := helper.ListAllContacts(db)
			if err != nil {
				fmt.Println(err)
			}
			continue

		//Search Contacts
		case 3:

			err := helper.SearchContact(db)
			if err != nil {
				fmt.Println(err)
			}

			continue

		//Update Contacts
		case 4:

			err := helper.UpdateContact(db)
			if err != nil {
				fmt.Println(err)
			}

			continue

		//Delete Contact
		case 5:

			err := helper.DeleteContacts(db)
			if err != nil {
				fmt.Println(err)
			}

			continue

		//Exit
		default:
			os.Exit(0)
		}
	}
}
