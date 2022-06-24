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
	fmt.Println("")
	parse("Niladri")

}

//Parse function that accepts an intrface, that will allow us to pass anything we want
func parse(ip interface{}) {

	input := reflect.ValueOf(ip)

	typeOfIP := input.Type()

	//If structure
	if input.Kind() == reflect.Struct {

		s := input.Interface()

		switch s.(type) {
		case obj:
			empSt := obj{}
			if reflect.DeepEqual(s, empSt) == false {
				for i := 0; i < input.NumField(); i++ {

					field := input.Field(i)

					if reflect.DeepEqual(field.Interface(), reflect.Zero(reflect.TypeOf(field.Interface())).Interface()) == false {

						fmt.Printf("\nNo = %d Field = %s Type = %s  ", i, typeOfIP.Field(i).Name, field.Type())
						parse(field.Interface())
					}
				}
			}
		case st:
			empSt := st{}

			//proceed if all fields are not empty
			if reflect.DeepEqual(s, empSt) == false {
				for i := 0; i < input.NumField(); i++ {

					field := input.Field(i)
				
					if reflect.DeepEqual(field.Interface(), reflect.Zero(reflect.TypeOf(field.Interface())).Interface()) == false {


						fmt.Printf("\nNo = %d Field = %s Type = %s  ", i, typeOfIP.Field(i).Name, field.Type())
						parse(field.Interface())
					}

				}

			}
		}

	}
	//If slice
	if input.Kind() == reflect.Slice {

		//Cheeck empty
		if input.Len() != 0 {

			for i := 0; i < input.Len(); i++ {

				fmt.Printf("\n[] Obj%d:", i+1)
				parse(input.Index(i).Interface())

			}

		}
	}

	//If map
	if input.Kind() == reflect.Map {

		keys := input.MapKeys()

		//MapKeys stores keys in reverse form,
		//Hence we are reversing it to get the original flow
		swap := reflect.Swapper(keys)
		for i := 0; i < len(keys)/2; i++ {
			swap(i, len(keys)-1-i)
		}

		//Check eempty
		if len(keys) != 0 {

			for _, key := range keys {
				fmt.Print("\nkey: ", key, " ->")
				parse(input.MapIndex(key).Interface())
			}
		}
	}

	if input.Kind() != reflect.Struct && input.Kind() != reflect.Slice && input.Kind() != reflect.Map {
		//print Result
		if input != reflect.Zero(reflect.TypeOf(input)) {

			fmt.Printf("Value = %v", input)

		}
	}

}
