/*WordCount:	This code takes user input as string & output word occourrance after trimming special characters and white spaces.

The "main" function takes the input from terminal using "buffio", make a slice of it, then trim it and send a copy of the slice
to the "wordCount" functioin, where a dictionary is mapped to hold each word as the key & no. of occourance as the value.
Now, result gets printed using the dictionary named "dict". */

package main

import (
	"fmt"
	"log"

	"github.com/niladridas/WordCount/helper"
)

func main() {
	fmt.Println("\n******Welcome to WordCount*****")

	//Taking user input
	str, err := helper.TakeInput()
	if err != nil {
		log.Fatal(err)
	}
	//Function wordCount
	wcMap := helper.WordCount(str)
	if wcMap == nil {
		log.Fatal("Empty Map Returned!")
	}
	//Print result
	err = helper.PrintTopTen(wcMap)
	if err != nil {
		print(err)
	}

}
