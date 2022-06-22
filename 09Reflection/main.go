package main

import (
	"fmt"
	"reflect"
)

type st struct {
	A int
	B obj
	C []obj
}

type obj struct {
	X int
	Y string
	Z map[string]st
}

//Main Function
func main() {

	a := 9

	b := obj{
		X: 10,
		Y: "ten",
	}

	c := []obj{
		{X: 1, Y: "one"},
		{X: 2, Y: "two"},
		{Z: map[string]st{
			"Check": {
				A: 99,
				B: obj{X: 1, Y: "one"},
				C: []obj{
					{X: 54},
					{Z: map[string]st{
						"Hello": {A: 4},
					},
					},
				},
			},
			"test": {A: 52},
		},
		}}

	data := st{
		A: a,
		B: b,
		C: c,
	}

	//fmt.Println(data)

	parse(data)

}

//Parse function that accepts an intrface, that will allow us to pass anything we want
//It will treat every data-structure with keyvalue pairs or fields as structure, check every fields and nested fields and print it
//if it gets other than that, like array int etc. "parse" will print it and return.
func parse(ip interface{}) {

	s := reflect.ValueOf(ip)

	typeOfST := s.Type()

	//Check if struct
	if s.Kind() != reflect.Struct {
		fmt.Println("Value : ", s)
		return
	}

	for i := 0; i < s.NumField(); i++ {

		field := s.Field(i)

		//If nested, call parse recursively
		if field.Kind() == reflect.Struct || field.Kind() == reflect.Slice || field.Kind() == reflect.Map {

			//If structure
			if field.Kind() == reflect.Struct {

				//nested structure
				ns := field.Interface()

				//Cheecking which type of structure it is
				switch ns.(type) {
				case obj:
					empSt := obj{}
					//proceed if all fields are not empty
					if reflect.DeepEqual(ns, empSt) == false {

						fmt.Printf("Field = %s Type = %s -> \n", typeOfST.Field(i).Name, field.Type())

						parse(field.Interface())

					}
				case st:
					empSt := st{}
					//proceed if all fields are not empty
					if reflect.DeepEqual(ns, empSt) == false {

						fmt.Printf("Field = %s Type = %s -> \n", typeOfST.Field(i).Name, field.Type())

						parse(field.Interface())

					}
				}

			}

			//If Slice
			if field.Kind() == reflect.Slice {

				content := reflect.ValueOf(field.Interface())

				//Cheeck empty
				if content.Len() != 0 {

					fmt.Printf("Field = %s Type = %s -> \n", typeOfST.Field(i).Name, field.Type())

					for i := 0; i < content.Len(); i++ {

						fmt.Printf("[] Obj%d:\n", i+1)
						parse(content.Index(i).Interface())

					}

				}

			}

			//If map
			if field.Kind() == reflect.Map {

				content := reflect.ValueOf(field.Interface())
				keys := content.MapKeys()

				//MapKeys stores keys in reverse form,
				//Hence we are reversing it to get the original flow
				swap := reflect.Swapper(keys)
				for i := 0; i < len(keys)/2; i++ {
					swap(i, len(keys)-1-i)
				}

				//Check eempty
				if len(keys) != 0 {

					fmt.Printf("Field = %s Type = %s -> \n", typeOfST.Field(i).Name, field.Type())

					for _, key := range keys {
						fmt.Println("key: ", key, " ->")
						parse(content.MapIndex(key).Interface())
					}
				}
			}

		} else {
			//print Result

			//Check zero value to ensure field not empty
			if field.Interface() != reflect.Zero(reflect.TypeOf(field.Interface())).Interface() {
				fmt.Printf("Field = %s Value = %v Type = %s \n", typeOfST.Field(i).Name, field.Interface(), field.Type())
			}

		}
	}

}
