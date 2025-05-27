package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// implementasi interface
func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}


func main(){
	person := Person{Name: "Hermawan"}
	SayHello(person)

	animal := Animal{Name: "Semut"}
	SayHello(animal)
}