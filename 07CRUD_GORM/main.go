package main

import (
	"CRUD_GORM/database"
	"CRUD_GORM/model"
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
		case 1:

			//taking Input
			fmt.Print("Name: ")
			name, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Wrong input: Name")
				continue
			}
			name = strings.TrimSpace(name)

			fmt.Print("Number: ")
			number, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Wrong input: Number")
				continue
			}
			number = strings.TrimSpace(number)

			//Phone no must bee of 10 digits
			if len(number) != 10 {
				fmt.Println("Please Input a 10 digit valid Number")
				continue
			}

			newContact := model.Contact{
				Name: name,
				Number: number,
			} 
			//calling function
			createdContact:= database.CreateContact(db, newContact)

			fmt.Println("Contact Created: ", createdContact.Name)

			continue

		case 2:

			contacts, err := database.ListAllContacts(db)
			if err != nil {

				fmt.Println(err)
				continue
			}

			//Print Result
			for index, c := range contacts {
				fmt.Printf("%d	Name: %s	Number: %s\n", index+1, c.Name, c.Number)
			}

			continue

		case 3:

			//Input
			fmt.Print("Enter Name to search: ")
			reader := bufio.NewReader(os.Stdin)
			name, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Wrong input: Name")
				continue
			}
			name = strings.TrimSpace(name)

			searchedContacts, err := database.SearchContacts(db, name)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Contacts found: ")
			for index, c := range searchedContacts {
				fmt.Printf("%d	Name: %s	Number: %s\n", index+1, c.Name, c.Number)
			}

			continue

		case 4:

			//Input
			fmt.Print("Enter Name to search: ")
			reader := bufio.NewReader(os.Stdin)
			name, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Wrong input: Name")
				continue
			}
			name = strings.TrimSpace(name)

			//Searching Contact to update
			searchedContacts, err := database.SearchContacts(db, name)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Contacts found: ")
			for index, c := range searchedContacts {
				fmt.Printf("%d	Name: %s	Number: %s\n", index+1, c.Name, c.Number)
			}

			fmt.Println("Enter the Sr. No. of the contact you want to update ")
			input, _ := reader.ReadString('\n')
			sNo, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
			if err != nil {
				fmt.Println("Failed to convert string into int")
				continue
			}

			sNo = sNo - 1
			if sNo < 0 || sNo >= int64(len(searchedContacts)) {

				fmt.Println("Sr. no. Exited")
				continue

			}

			fmt.Println("Enter Updates")
			fmt.Print("Name: ")
			name, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Wrong input: Name")
				continue
			}
			name = strings.TrimSpace(name)

			fmt.Print("Number: ")
			number, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Wrong input: Number")
				continue
			}
			number = strings.TrimSpace(number)

			if len(number) != 10 {
				fmt.Println("Please Input a 10 digit valid Number")
				continue
			}

			updateData := model.Contact{
				Name: name,
				Number: number,
			}
			updatedContact := database.UpdateContact(db, searchedContacts[sNo], updateData)
			fmt.Printf(`Update Successful: 
			Name: %s
			Number: %s`, updatedContact.Name, updatedContact.Number )

			continue

		case 5:

			//Input
			fmt.Print("Enter Name to dlete: ")
			reader := bufio.NewReader(os.Stdin)
			name, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Wrong input: Name")
				continue
			}
			name = strings.TrimSpace(name)

			//Searching Contact to update
			searchedContacts, err := database.SearchContacts(db, name)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Contacts found: ")
			for index, c := range searchedContacts {
				fmt.Printf("%d	Name: %s	Number: %s\n", index+1, c.Name, c.Number)
			}

			fmt.Println("Enter the Sr. No. of the contact you want to update ")
			input, _ := reader.ReadString('\n')
			sNo, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
			if err != nil {
				fmt.Println("Failed to convert string into int")
				continue
			}

			sNo = sNo - 1
			if sNo < 0 || sNo >= int64(len(searchedContacts)) {

				fmt.Println("Sr. no. Exited")
				continue

			}

			deletedAt := database.DeleteContact(db, searchedContacts[sNo])

			fmt.Println("Contact Deletd\n Deleted at : ", deletedAt)

			continue
		default:
			os.Exit(0)
		}
	}
}
