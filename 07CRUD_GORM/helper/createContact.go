package helper

import (
	"CRUD_GORM/database"
	"CRUD_GORM/model"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)

func CreateContact(db *gorm.DB) error {

	//taking Input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Wrong input: Name")
	}
	name = strings.TrimSpace(name)

	fmt.Print("Address: ")
	add, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Wrong input: Name")
	}
	add = strings.TrimSpace(add)

	fmt.Print("Number: ")
	number, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Wrong input: Number")

	}
	number = strings.TrimSpace(number)

	//Phone no must bee of 10 digits
	if len(number) != 10 {
		return errors.New("Please Input a 10 digit valid Number")

	}

	//Create New Contact
	newContact := model.Contact{
		Name: name,
		Add:  add,
	}

	newPh := model.Ph{
		Number: number,
	}

	ph := database.CreateContact(db, newContact, newPh)

	fmt.Println("Contact Created: ", ph.Contact.Name)

	return nil
}
