package helper

import (
	"CRUD_GORM/database"
	"CRUD_GORM/model"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator"
)

func CreateContact() error {

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

	//Create New Contact
	newContact := model.Contact{
		Name: name,
		Add:  add,
	}

	//Validation
	validate := validator.New()
	err = validate.Struct(newContact)
	if err != nil {
		return err
	}

	newPh := model.Ph{
		Number: number,
	}

	//Validation
	err = validate.StructPartial(newPh, "Number")
	if err != nil {
		return errors.New(fmt.Sprint("Failed to create contact", err.Error()))
	}

	ph, err := database.CreateContact(newContact, newPh)
	if err != nil {
		return err
	}

	fmt.Println("Contact Created: ", ph.Contact.Name)

	return nil
}
