package main

import "fmt"

type Man struct {
	Name string
}

// penggunaan asterisk untuk pointer
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	hermawan := Man{"Hermawan"}
	hermawan.Married()
	fmt.Println(hermawan.Name)
}