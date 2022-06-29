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

			fmt.Print("Address: ")
			add, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Wrong input: Name")
				continue
			}
			add = strings.TrimSpace(add)

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

			//Create New Contact
			newContact := model.Contact{
				Name: name,
				Add: add,
			}

			newPh := model.Ph{
				Number: number,
			} 

			ph:= database.CreateContact(db, newContact, newPh)


			fmt.Println("Contact Created: ", ph.Contact.Name)

			continue

		case 2:

			listOfPhNo, err := database.ListAllContacts(db)
			if err != nil {

				fmt.Println(err)
				continue
			}

			//Print Result
			for index, ph := range listOfPhNo {
				fmt.Printf("%d	Name: %s	Add: %s		Number: %s\n", index+1, ph.Contact.Name,ph.Contact.Add, ph.Number)
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

			database.PrintContacts(db, searchedContacts)

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
				fmt.Printf("%d	Name: %s	Add: %s\n", index+1, c.Name, c.Add)
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

			fmt.Print("Address: ")
			add, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Wrong input: Name")
				continue
			}
			add = strings.TrimSpace(add)

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

			updateContact := model.Contact{
				Name: name,
				Add: add,
			}
			updatePh := model.Ph{
				Number: number,
			}

			updatedContact, updatedPh := database.UpdateContact(db, searchedContacts[sNo], updateContact, updatePh)
			fmt.Printf(`Update Successful: 
			Name: %s
			Add: %s
			updated Number: %s`, updatedContact.Name, updateContact.Add , updatedPh.Number )

			continue

		case 5:

			//Input
			fmt.Print("Enter Name to delete: ")
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
				fmt.Printf("%d	Name: %s	Add: %s\n", index+1, c.Name, c.Add)
			}

			fmt.Println("Enter the Sr. No. of the contact you want to delete ")
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
