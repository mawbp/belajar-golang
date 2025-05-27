package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string `required:"true" max:"10"`
	Adress string `required:"true" max:"10"`
	Email string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)

	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
		fmt.Println(valueField.Tag.Get("required"))
		fmt.Println(valueField.Tag.Get("max"))
	}
}

func isValid(data any) (result bool) {
	result = true
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(data).Field(i).Interface() 
			result = data != ""
			if result == false {
				return result
			}
		}
	}

	return result
}

func main() {
	readField(Sample{"Hermawan"})
	readField(Person{"Hermawan", "", ""})

	person := Person{"Hermawan", "Surabaya", "email"}
	fmt.Println(isValid(person))
}
