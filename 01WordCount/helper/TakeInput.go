package helper

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func TakeInput() (string, error) {
	//Taking user input
	fmt.Print("Input String: ")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	//Error Handling
	if err != nil {
		return "", errors.New(fmt.Sprintf(`Could not take input!
		error : %v`, err))
	}

	//Removing white spaces
	str = strings.TrimSpace(str)

	//Checking empty string
	if str == "" {
		return "", errors.New("TakeInput : Empty String Input!")
	}

	return str, nil
}
